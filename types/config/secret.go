package config

import "github.com/discless/discless/types/kinds"

type Secret struct {
	Kind	kinds.Kind			`yaml:"kind"`
	Secrets	map[string]string	`yaml:"secrets"`
}