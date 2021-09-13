package api

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/discless/discless/dispatcher"
	"github.com/discless/discless/types"
	"github.com/discless/discless/types/config"
	"net/http"
	"strings"
)

func genBearer(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

func NewBot(w http.ResponseWriter, r *http.Request)  {
	botc := &config.Bot{}

	err := json.NewDecoder(r.Body).Decode(botc)
	if err != nil {
		panic(err)
	}
	r.Body.Close()
	fmt.Println("New bot:",botc.Name)

	token := botc.Token
	ses, err := discordgo.New("Bot "+token)

	if err != nil {
		panic(err)
	}
	bot := &types.Self{
		//Id:			ses.State.User.ID,
		Name:     	botc.Name,
		Bot:      	ses,
		Prefix:   	botc.Prefix,
		Commands: 	make(map[string]*types.Command),
	}


	dispatcher.Manager[bot.Name] = bot


	if err != nil {
		panic(err)
	}

	dispatcher.Manager[bot.Name].Bot.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		//fmt.Println(dispatcher.Manager[s.State.User.ID])
		if !m.Author.Bot && strings.Split(m.Content,"")[0] == dispatcher.Manager[s.State.User.ID].Prefix{
			//fmt.Println("Checks out",strings.Split(strings.Replace(m.Content, dispatcher.Manager[s.State.User.ID].Prefix,"", -1)," ")[0])
			if command, ok := dispatcher.Manager[s.State.User.ID].Commands[strings.Split(strings.Replace(m.Content, dispatcher.Manager[s.State.User.ID].Prefix,"", -1)," ")[0]]; ok {
				command.Exec(dispatcher.Manager[s.State.User.ID], s, m, strings.Split(strings.Replace(m.Content, dispatcher.Manager[s.State.User.ID].Prefix,"", -1)," ")[1:])
			}
		}
	})
	dispatcher.Manager[bot.Name].Bot.Open()

	dispatcher.Manager[ses.State.User.ID] = bot
	dispatcher.Manager[ses.State.User.ID].Id = ses.State.User.ID

	bearer := genBearer(12)
	dispatcher.Sessions[bearer] = bot.Name
	fmt.Fprint(w,bearer)
}
