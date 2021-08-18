package types

import (
	"github.com/bwmarrin/discordgo"
)

type FCommand func(self *Self, s *discordgo.Session, m *discordgo.MessageCreate, args []string) error
