package config

import "github.com/discless/discless/types/kinds"

type Bot struct {
	Kind		kinds.Kind							`yaml:"kind",json:"kind"`
	Name		string								`yaml:"name",json:"name"`
	Token		string								`yaml:"token",json:"token"`
	Prefix		string								`yaml:"prefix",json:"prefix"`
	Services	struct{
		Database	struct{
			MongoDB	struct{
				User 			string 				`yaml:"user",json:"user"`
				Password 		string				`yaml:"password",json:"password"`
				Adress			string				`yaml:"adress",json:"adress"`
				Database		string				`yaml:"database",json:"database"`
			}										`yaml:"mongodb",json:"mongo_db"`

			SQL		struct{
				User 			string 				`yaml:"user",json:"user"`
				Password 		string				`yaml:"password",json:"password"`
				Adress			string				`yaml:"adress",json:"adress"`
				Database		string				`yaml:"database",json:"database"`
				Params			map[string]string	`yaml:"params",json:"params"`
			}										`yaml:"sql",json:"sql"`
		}											`yaml:"database",json:"database"`
	}												`yaml:"services,omitempty",json:"services"`
}
