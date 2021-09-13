# Discless [![Go](https://github.com/discless/discless/actions/workflows/go.yml/badge.svg)](https://github.com/discless/discless/actions/workflows/go.yml)
Discless is a framework that lets you create Discord bots the FaaS way, meaning that you only need to write the commands and Discless handles the rest for you.

##🚀 Getting Started
### 1. Install Discless
Before creating your first function, install the Discless CLI by running
```shell
$ sudo env GOBIN=/bin go install github.com/discless/cli/cmd/discless@v0.0.2.1
```
_Make sure you have go installed_
### 2. Create a function and bot
First, run the `new bot` command and enter your bots token
```shell
$ discless new bot <bot name> <prefix>
Created bot in bot.yaml
```

To create a new function, run
```shell
$ discless new function <function name>
Created the function <function name>
Edit its configuration in function.yaml or edit the function in <function name>.go
```
This creates a configuration file for the function (`function.yaml`) and a golang file that looks like
```go
package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/discless/discless/types"
)

func Handler(self *types.Self, s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
	s.ChannelMessageSend(m.ChannelID, "Pong!")
	return nil
}
```
You can freely edit this file.

### 3. Configure your bot's token
First, create a new secret for your token
```shell
$ discless new secret token NDMyMTkx...
Created secret in secret.yaml
```

To use the token in your bots configuration, open your bots configuration and change the following
```
- token: 
+ token: secret.token
```

### 4. Run the bot and deploy your function
Now you can run your bot
```shell
$ discless up bot.yaml
<bot-name> is running
```

To get your command up and running on the bot, run
```shell
$ discless deploy <bot name> <function configuration>.yaml
Succesfully uploaded the <function name> command
```

## 📝 Features

🔐 All API calls are secured by TLS </br>
🧠 Simple and intuitive Command Line Interface </br>
💨 Fast and easy to manage Discord applications

## 🗺️ Roadmap

- [x] SSL encryption
- [x] Secret configuration
- [x] Deployable commands and bots 
- [ ] Support all Discord events
- [ ] Recursive command deployment
- [ ] Integrate databases </br>
  </br>
  _More coming soon..._
  
## 📕 License

**[MIT License](https://github.com/discless/discless/blob/main/LICENSE)**