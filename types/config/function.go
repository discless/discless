package config

import "github.com/discless/discless/types/kinds"

type Config struct {
	Kind		kinds.Kind			`yaml:"kind"`
	Functions 	map[string]Function `yaml:"functions"`
}

type Function struct {
	Type		string `yaml:"type,omitempty",json:"type,omitempty"`
	File		string `yaml:"file",json:"file"`
	Function	string `yaml:"function",json:"function"`
	Category	string `yamle:"category,omitempty",json:"category,omitempty"`
}