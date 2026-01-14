package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type UwuBot struct {
	Session *discordgo.Session
}

func NewUwuBot(token string) (*UwuBot, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	bot := &UwuBot{
		Session: dg,
	}

	dg.AddHandler(bot.onMessageCreate)
	return bot, nil
}

func (b *UwuBot) Run() error {
	err := b.Session.Open()
	if err != nil {
		return err
	}
	defer b.Session.Close()

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	select {}
}

func (b *UwuBot) onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	msg, err := NewMessageFromStr(m.Content)
	if err != nil {
		// If it was just the prefix but no text, remind the user
		if strings.HasPrefix(m.Content, "UwU ") && err.Error() == "empty content" {
			s.ChannelMessageSend(m.ChannelID, "p-pwease give me some text to uwu-ify! ^w^")
		}
		// Otherwise ignore (might not be a command at all)
		return
	}

	finalMessage := msg.Uwuify()
	s.ChannelMessageSend(m.ChannelID, finalMessage)
}
