package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"regexp"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/organ/golibtox"
	"github.com/thoj/go-ircevent"
)

type Config struct {
	IRC      IrcServer
	Tox      ToxServer
	Settings Settings
}

type ToxServer struct {
	Address   string
	Port      uint16
	PublicKey string
}
type IrcServer struct {
	Address string
	Channel string
}
type Settings struct {
	Regex string
}

var ircMessage string
var toxMessage string
var toxGroupNum int

//Once everything else works, I plan on changing the above to slices. As now messages may get overridden.

var cfg Config

func main() {
	if _, err := toml.DecodeFile("config", &cfg); err != nil {
		panic(err)
	}
	//tox connecting
	bridgebot, err := golibtox.New()
	if err != nil {
		panic(err)
	}
	err = loadData(bridgebot)
	if err != nil {
		fmt.Println("Could not load save data!")
	}

	bridgebot.SetStatusMessage([]byte("Invite me to a groupchat!"))
	bridgebot.SetName("BridgeBot")
	// irc connecting
	con := irc.IRC("BridgeBot", "BridgeBot")
	err = con.Connect(cfg.IRC.Address)
	if err != nil {
		panic(err)
	}
	bridgebotAddr, _ := bridgebot.GetAddress()
	fmt.Println("ID bridgebot: ", hex.EncodeToString(bridgebotAddr))

	bridgebot.CallbackFriendRequest(onFriendRequest)
	bridgebot.CallbackFriendMessage(onFriendMessage)
	bridgebot.CallbackGroupInvite(onGroupInvite)
	bridgebot.CallbackGroupMessage(onGroupMessage)

	con.AddCallback("001", func(e *irc.Event) {
		con.Join(cfg.IRC.Channel)
	})
	con.AddCallback("PRIVMSG", onIrcMessage)
	err = bridgebot.BootstrapFromAddress(cfg.Tox.Address, cfg.Tox.Port, cfg.Tox.PublicKey)
	if err != nil {
		panic(err)
	}

	isRunning := true

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ticker := time.NewTicker(25 * time.Millisecond)

	go con.Loop()
	for isRunning {
		select {
		case <-c:
			fmt.Println("Saving...")
			if err := saveData(bridgebot); err != nil {
				fmt.Println(err)
			}
			fmt.Println(" Killing")
			con.Quit()
			bridgebot.Kill()
			isRunning = false
		case <-ticker.C:
			if len(ircMessage) > 0 {
				bridgebot.GroupMessageSend(toxGroupNum, []byte(ircMessage))
				ircMessage = ""
			}
			if len(toxMessage) > 0 {
				con.Privmsg(cfg.IRC.Channel, toxMessage)
				toxMessage = ""
			}
			bridgebot.Do()
			break
		}
	}
}

func onFriendRequest(t *golibtox.Tox, publicKey []byte, data []byte, length uint16) {
	name, _ := t.GetSelfName()
	fmt.Printf("[%s] New friend request from %s\n", name, hex.EncodeToString(publicKey))
	// Auto-accept friend request
	t.AddFriendNorequest(publicKey)
}

func onFriendMessage(t *golibtox.Tox, friendnumber int32, message []byte, length uint16) {
	name, _ := t.GetSelfName()
	friend, _ := t.GetName(friendnumber)
	fmt.Printf("[%s] New message from %s : %s\n", name, friend, string(message))
}

func onGroupInvite(t *golibtox.Tox, friendnumber int32, groupPublicKey []byte) {
	name, _ := t.GetSelfName()
	friend, _ := t.GetName(friendnumber)
	fmt.Printf("[%s] Group invite from %s\n", name, friend)
	t.JoinGroupchat(friendnumber, groupPublicKey)
	fmt.Printf("%s", cfg.Settings.Regex)
}

func onGroupMessage(t *golibtox.Tox, groupnumber int, friendgroupnumber int, message []byte, length uint16) {
	//TODO: Check if the message is from the bot! This causes the response to be echoed in the IRC
	fmt.Printf("[Groupchat #%d]:%s\n", groupnumber, string(message))
	validMessage, err := regexp.Match(cfg.Settings.Regex, message)
	if err != nil {
		panic(err)
	}
	if validMessage {
		toxMessage = string(message)
		toxGroupNum = groupnumber
	} else {
		toxMessage = ""
	}
}

func loadData(t *golibtox.Tox) error {
	data, err := ioutil.ReadFile("bridge_data")
	if err != nil {
		return err
	}
	err = t.Load(data)
	return err
}

func saveData(t *golibtox.Tox) error {
	data, err := t.Save()
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("bridge_data", data, 0644)
	return err
}

//irc functions
func onIrcMessage(e *irc.Event) {
	if e.Nick == "BridgeBot" { //I'll unhardcode this once I get around to making the config files
		ircMessage = ""
		return
	}
	ircMessage = e.Message
}
