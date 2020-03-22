# Aisaka Taiga (逢坂•大河)
[![Discord](https://discordapp.com/api/guilds/397479560876261377/embed.png)](https://discord.gg/mDkMbEh)

A [Discord](https://discordapp.com) Bot written in [Golang](https://golang.org) with [discordgo](https://github.com/bwmarrin/discordgo) and the [Sapphire Framework](https://github.com/sapphire-cord/sapphire)

This could also serve as a real-life example of using the Sapphire framework, it is constantly up to date with the latest features of the framework.

- Aisaka Taiga Lounge: [discord.gg/mDkMbEh](https://discord.gg/mDkMbEh) (Bot updates, help etc.)
- Sapphire Server: [discord.gg/ArwQrH4](https://discord.gg/ArwQrH4) (Using the framework for your own bots etc.)

## Invite
If you want to use the bot, invite it to your server [here](https://discordapp.com/oauth2/authorize?client_id=397796982120382464&permissions=2016537702&scope=bot)

## Running it yourself
To run this bot you need:
- Golang 1.13+
- PostgreSQL 9+

Fetch the code with
```sh
$ go get github.com/pollen5/taiga
```
That will fetch the project along with all required dependencies (Don't worry there aren't too much, I myself like to use as little dependencies as possible)

then `cd` to `$GOPATH/src/github.com/pollen5/taiga` (Use `go env GOPATH` to find your `$GOPATH`)

Copy `config.json.example` to `config.json` and replace the fields with actual values.

> **NOTE:** For Bot token make sure you prefix it with `Bot ` yourself or API requests will fail.

Setup a PostgreSQL database (you are on your own for that) and put the connection URI in `config.json`

> **NOTE:** If you are not using SSL append `?sslmode=disable` to the PostgreSQL connection URI, the SQL driver by default enables it.

You may want to edit `constants/constants.go` to match your settings.

Build the code:
```sh
$ go build main.go
```
For **first time only** run the executable with `-initdb`
```sh
$ ./main -initdb
```
This will connect to the given PostgreSQL database and create the required tables, once this is done you can start the bot.
```sh
$ ./main
```
> **NOTE:** On Windows replace the following `./main` calls to `.\main.exe`. Windows is stupid i know.

And success! Enjoy the results.

> **NOTE:** Please do no contact me for help, what is written is what is written and that's all you need to know to run it so if you didn't understand something that's your fault. I'm not willing to help total noobs run a copy of the bot, instead [invite the bot to your server.](https://discordapp.com/oauth2/authorize?client_id=397796982120382464&permissions=2016537702&scope=bot) However I'm willing to help in bot related issues, things that aren't related to "How do i setup postgresql" and similar.

### Schema Updates
Whenever i update the bot and changes happen in the database schema, just read the changes in `db/schema.go` and apply `ALTER TABLE ADD COLUMN (whatever is in the schema)` to add the column. In the future there will be an easier way update the schema but this is it for now.

### Development flow
If you want to customize the bot then make sure you install [spgen](https://github.com/sapphire-cord/spgen) and run it when a new command is added or it's existing **information** (not code) changes.

See the [spgen](https://github.com/sapphire-cord/spgen) page for more information on how it works.

## License
All the code here is released under the [MIT LICENSE](LICENSE)!

Do feel free to take any code from here as long as you credit me for it. (Except for very small snippets, I don't mind)
