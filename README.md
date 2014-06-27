BridgeBot
=========
####What is it?
BridgeBot is an IRC bot/Tox bot that is designed to bridge the gap (hence bridge) between IRC bots and tox group chats. This follows a very simular concept to SyncBot, but with a twist more on relaying only what could be nessissary to IRC.
The bot will connect happily be invited to any group (on tox), meanwhile the IRC componet of it joins the channel on the specified irc channel in the configuration file.

You will manually need to stick any IRC bot into the channel, it is recomended you take action to prevent others from entering your channel. The BrigeBot in tox will then filter out any messages that don't match the regular expression, and ones that are left will be sent over to the IRC channel, from there the bot responds and the response is sent over to the Tox group chat, seemlessly. This is done to prevent spam, other measures may be put in place later to prevent the IRC component from being kicked when being spammed with messages.

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
