BridgeBot
=========
Currently not complete, it doesn't work yet, but it's getting there.
####What is it?
BridgeBot is an IRC bot/Tox bot that is designed to bridge the gap (hence bridge) between IRC bots and tox bots. This follows a very simular concept to SyncBot.
The bot will connect happily be invited to any group (on tox), meanwhile the IRC componet of it joins the channel on the specified irc channel.

You will manually need to stick any IRC bot into the channel and set it to +m or something else (you'll see why in a sec). The BrigeBot in tox will then filter out any messages that don't start with the specified prefix for the irc commands, and ones that are left will be sent over to the IRC channel, from there the bot responds and the response is sent over to the Tox group chat, seemlessly. This is done to prevent spam, other measers may be put in place later to prevent being kicked.

A heck of a lot of code is borrowed from the examples provided in golibtox, which can be found at https://github.com/organ/golibtox/tree/master/ in addition this is the actual library used.

More info later.
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
