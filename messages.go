package main

import (
	"flag"
	"fmt"
	"strings"
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
	// --- IMPLEMENTATION OF MODIFICATION HERE ---

	return m.Content
}
