package types

import (
	"database/sql"
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
)

type Self struct {
	Name		string				`json:"name"`
	Bot 		*discordgo.Session
	Prefix 		string				`json:"prefix"`
	Commands	map[string]*Command
	Id			string				`json:"id"`
	Dispatcher	map[string]interface{}
}

type Command struct {
	Exec 		FCommand
	Name 		string
	Args 		map[string]reflect.Type
	Category 	string
}

type DB struct {
	Mongo 	*mongo.Database
	SQL		*sql.DB
}