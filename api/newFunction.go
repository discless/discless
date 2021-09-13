package api

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/discless/discless/dispatcher"
	"github.com/discless/discless/types"
	"io/ioutil"
	"net/http"
	"os/exec"
	"plugin"
	"strconv"
	"strings"
	"time"
)

func Apply(w http.ResponseWriter, r *http.Request)  {
	bot := r.FormValue("Bot")
	auth := strings.SplitN(r.Header.Get("Authorization")," ",2)

	if len(auth) != 2 || auth[0] != "Basic" {
		http.Error(w, "authorization failed, no authentication provided", http.StatusUnauthorized)
		return
	} else if val, ok := dispatcher.Sessions[auth[1]]; !ok || val != bot {
		http.Error(w, "authorization failed, token doesn't match bot or does not exist", http.StatusUnauthorized)
		return
	}
	
	err := r.ParseMultipartForm(32 << 20) // maxMemory 32MB
	if err != nil {
		http.Error(w,"file is too big",http.StatusBadRequest)
		return
	}
	fmt.Println("New function:",r.FormValue("Name"))
	_, h, _ := r.FormFile("Function")
	f, err := h.Open()
	if err != nil {
		panic(err)
	}
	reader, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	tempfile := strconv.FormatInt(time.Now().Unix(),16)
	ioutil.WriteFile(tempfile + ".go", reader, 0644)
	cmd := exec.Command("go", "build","--buildmode=plugin",tempfile + ".go")
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
	p, err := plugin.Open(tempfile + ".so")
	if err != nil {
		panic(err)
	}

	function, err := p.Lookup(r.FormValue("Function"))
	if err != nil {
		panic(err)
	}
	command := &types.Command{
		Exec:     function.(func(self *types.Self, s *discordgo.Session, m *discordgo.MessageCreate, args []string) error ),
		Name:     r.FormValue("Name"),
		Args:     nil,
		Category: "",
	}

	dispatcher.Manager[bot].Commands[command.Name] = command
}
