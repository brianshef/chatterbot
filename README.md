# chatterbot

A [Discord bot](https://discord.com/developers/applications/809107559033208936/information) that just chats away, occasionally. Use it to keep the conversation moving in your Discord server.
Every few messages, it will interject with its own. Supply your own custom list of phrases for chatterbot to use, easy-peasy!

## Usage

Place a text file named `phrases.txt` in the bot's root directory.
Add one phrase per line in the file; these will be randomly chosen by the bot at runtime to be sent into the Discord server.

### Running the bot

```
./chatterbot -t <token>
```

## Inviting the bot to your Discord server

https://discord.com/api/oauth2/authorize?client_id=809107559033208936&permissions=68608&scope=bot

## Sample `phrases.txt`

```
Hello, world!
Beautiful day, isn't it?
Are you working hard, or hardly working?
Heard any interesting news lately?
It sure would be nice to be in Hawaii right about now.
You know what they say, all roads lead to Rome.
We couldn't ask for a nicer day, could we?
How about this weather?
What's cookin', good lookin'?
```