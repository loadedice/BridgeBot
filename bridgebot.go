package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/organ/golibtox"
    "github.com/thoj/go-ircevent"
)

type ToxServer struct {
	Address   string
	Port      uint16
	PublicKey string
}
type IrcServer struct {
    Address string
    Channel string
}

func main() {

	Tserver := &ToxServer{"37.187.46.132", 33445, "A9D98212B3F972BD11DA52BEB0658C326FCCC1BFD49F347F9C2D3D8B61E1B927"}
    Isever  := &IrcServer{"irc.freenode.net:6667", "#some-test44554"}

	bridgebot, err := golibtox.New()
	if err != nil {
		panic(err)
	}
	bridgebot.SetName("BridgeBot")

	bridgebotAddr, _ := bridgebot.GetAddress()
	fmt.Println("ID bridgebot: ", hex.EncodeToString(bridgebotAddr))

	// We can set the same callback function for both *Tox instances
	bridgebot.CallbackFriendRequest(onFriendRequest)
	bridgebot.CallbackFriendMessage(onFriendMessage)
    bridgebot.CallbackGroupInvite(onGroupInvite)
    bridgebot.CallbackGroupMessage(onGroupMessage)

	err = bridgebot.BootstrapFromAddress(Tserver.Address, Tserver.Port, Tserver.PublicKey)
	if err != nil {
		panic(err)
	}

	isRunning := true

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ticker := time.NewTicker(25 * time.Millisecond)

	for isRunning {
		select {
		case <-c:
				fmt.Println(" Killing")
				isRunning = false
				bridgebot.Kill()
			
		case <-ticker.C:
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
    name,_ := t.GetSelfName()
    friend,_ := t.GetName(friendnumber)
    fmt.Printf("[%s] Group invite from %s\n", name, friend)
    t.JoinGroupchat(friendnumber,groupPublicKey)
}

func onGroupMessage(t *golibtox.Tox, groupnumber int, friendgroupnumber int, message []byte, length uint16){
    //name, _ := t.GetName(friendgroupnumber)
    fmt.Printf("[%s]:%s", "message", string(message))
}
