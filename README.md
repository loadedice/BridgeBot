note: this is broken and I don't have any plans to fix it any time soon.

BridgeBot
=========
####What is it?
BridgeBot is an IRC bot/Tox bot that is designed to bridge the gap (hence bridgebot) between IRC bots and tox group chats. This follows a very similar concept to SyncBot, but with a twist more on relaying only what could be necessary to IRC, although you could change the regular expression in the configuration file to let it send all messages, like SyncBot does. It can also fully function like syncbot (relaying the names as well) with the boolean SyncBotMode set to true in the configuration file.
The bot will connect happily be invited to any group (on tox), meanwhile the IRC component of it joins the channel on the specified irc channel in the configuration file.

You will manually need to stick any IRC bot into the channel, it is recommended you take action to prevent others from entering your channel if you want to only relay the response of one bot. The BrigeBot in tox will then filter out any messages that don't match the regular expression, and ones that are left will be sent over to the IRC channel, from there the bot responds and the response is sent over to the Tox group chat, seamlessly. This is done to prevent spam, other measures may be put in place later to prevent the IRC component from being kicked when being spammed with messages.

####How to build
Resolve dependancies
```
go get github.com/organ/golibtox
go get github.com/thoj/go-ircevent
go get github.com/BurntSushi/toml
```
To run
```
go run bridgebot.go
```
