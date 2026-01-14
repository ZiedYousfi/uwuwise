package main

import (
	"flag"
	"fmt"
	"regexp"
	"strings"
)

const LIMIT_REPEATED_VOWELS = 2

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
	result := m.Content

	// Protection phase: Identify parts to ignore (mentions, emojis, URLs)
	ignoreRegex := regexp.MustCompile(`(https?://\S+|<a?:\w+:[0-9]+>|<@!?[0-9]+>|<#[0-9]+>|<@&[0-9]+>)`)
	var protected []string
	result = ignoreRegex.ReplaceAllStringFunc(result, func(match string) string {
		protected = append(protected, match)
		return fmt.Sprintf("{{P%d}}", len(protected)-1)
	})

	// Transformation phase

	// Replace 'r' and 'l' with 'w'
	result = regexp.MustCompile(`[rl]`).ReplaceAllString(result, "w")
	result = regexp.MustCompile(`[RL]`).ReplaceAllString(result, "W")

	// N + vowel replacement
	if !m.Flags.NoNya {
		nyaCount := 0
		result = regexp.MustCompile(`[nN]([aeiouyAEIOUY])`).ReplaceAllStringFunc(
			result,
			func(match string) string {
				replacement := "ny"
				if nyaCount%2 == 0 {
					replacement += "aa"
				} else {
					replacement += "ee"
				}
				nyaCount++
				return replacement
			},
		)
	}

	// Add stuttering
	if m.Flags.Stutter {
		result = regexp.MustCompile(`\b([a-zA-Z])`).ReplaceAllStringFunc(
			result,
			func(match string) string {
				return fmt.Sprintf("%s-%s", match, match)
			},
		)
	}

	// Repeated vowels
	vowels := "aeiouyAEIOUY"
	for _, v := range vowels {
		re := regexp.MustCompile(fmt.Sprintf("[%c]{%d,}", v, LIMIT_REPEATED_VOWELS+1))
		result = re.ReplaceAllString(result, string(v))
	}

	// Restoration phase: Replace placeholders back with original content
	for i, original := range protected {
		placeholder := fmt.Sprintf("{{P%d}}", i)
		result = strings.ReplaceAll(result, placeholder, original)
	}

	// Add faces if requested
	if m.Flags.Faces {
		faces := []string{"uwu", ">w<", "^w^", ":3", "owo", "x3", "rawr"}
		// Use a simple selection based on message length as a pseudo-random seed
		face := faces[len(result)%len(faces)]
		result = result + " " + face
	}

	return result
}
