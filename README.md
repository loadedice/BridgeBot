BridgeBot
=========
####What is it?
BridgeBot is an IRC bot/Tox bot that is designed to bridge the gap (hence bridge) between IRC bots and tox bots.
It works in a very simple way, somewhat simular to how I imagine the synbot works.
The bot will connect happily be invited to any group channel, meanwhile the IRC componet of it joins the channel on the specified irc channel. You will manually need to stick any IRC bot into the channel and set it to +m or something else, unless you want other people sending messages via the bot. The BrigeBot in tox will then filter out any messages that start with the specified prefix for the irc commands, and ones that do match will be sent over to the IRC channel, from there the bot responds and the response is sent over to the Tox group chat, seemlessly.

More info later.
