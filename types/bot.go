package types

import (
	"github.com/bwmarrin/discordgo"
	"reflect"
)

type Self struct {
	Name		string				`json:"name"`
	Bot 		*discordgo.Session
	Prefix 		string				`json:"prefix"`
	Commands	map[string]*Command
	Id			string				`json:"id"`
}

type Command struct {
	Exec 		FCommand
	Name 		string
	Args 		map[string]reflect.Type
	Category 	string
}
