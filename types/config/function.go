package config

import "github.com/discless/discless/types/kinds"

type Config struct {
	Kind		kinds.Kind			`yaml:"kind"`
	Functions 	map[string]Function `yaml:"functions"`
}

type Function struct {
	File		string `yaml:"file",json:"file"`
	Function	string `yaml:"function",json:"function"`
	Category	string `yamle:"category",json:"category"`
}