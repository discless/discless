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
	"time"
)

func Apply(w http.ResponseWriter, r *http.Request)  {
	err := r.ParseMultipartForm(32 << 20) // maxMemory 32MB
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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

	dispatcher.Manager[r.FormValue("Bot")].Commands[command.Name] = command
}
