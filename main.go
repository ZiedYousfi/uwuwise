package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type UwuFlags struct {
	Stutter bool
	Faces   bool
	NoNya   bool
}

type MessageToUwuify struct {
	Content string
	Flags   UwuFlags
}

func NewMessageFromStr(input string) (*MessageToUwuify, error) {
	// Check if message starts with "UwU "
	if !strings.HasPrefix(input, "UwU ") {
		return nil, fmt.Errorf("not an uwu command")
	}

	// Get the arguments after "UwU "
	argStr := strings.TrimPrefix(input, "UwU ")
	args := strings.Fields(argStr)

	// Define flags for the CLI-style interaction
	fs := flag.NewFlagSet("uwu", flag.ContinueOnError)
	stutter := fs.Bool("stuttew", false, "m-makes da stawts of wowds b-bouncy")
	faces := fs.Bool("faces", false, "adds a wittle facey-wacey")
	noNya := fs.Bool("no-nya", false, "stops da nyanyanya")

	// Parse flags from the message content
	err := fs.Parse(args)
	if err != nil {
		return nil, err
	}

	content := strings.Join(fs.Args(), " ")
	if content == "" {
		return nil, fmt.Errorf("empty content")
	}

	return &MessageToUwuify{
		Content: content,
		Flags: UwuFlags{
			Stutter: *stutter,
			Faces:   *faces,
			NoNya:   *noNya,
		},
	}, nil
}

func (m *MessageToUwuify) Uwuify() string {
	// --- IMPLEMENTATION OF MODIFICATION STARTS HERE ---
	// User (Inaya) will implement the magic here!
	// Use: m.Content, m.Flags.Stutter, m.Flags.Faces, m.Flags.NoNya

	return m.Content
}

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

func main() {
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		fmt.Println("No DISCORD_TOKEN found. Please set DISCORD_TOKEN environment variable.")
		return
	}

	bot, err := NewUwuBot(token)
	if err != nil {
		fmt.Printf("Error creating Discord session: %v\n", err)
		return
	}

	if err := bot.Run(); err != nil {
		fmt.Printf("Error running bot: %v\n", err)
	}
}
