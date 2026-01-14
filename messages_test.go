package main

import (
	"testing"
)

func TestNewMessageFromStr(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		stutter bool
		faces   bool
		noNya   bool
		wantErr bool
	}{
		{
			name:    "Valid command - basic",
			input:   "UwU Hello World",
			want:    "Hello World",
			stutter: false,
			faces:   false,
			noNya:   false,
			wantErr: false,
		},
		{
			name:    "Valid command - with stutter",
			input:   "UwU -stuttew Hello World",
			want:    "Hello World",
			stutter: true,
			faces:   false,
			noNya:   false,
			wantErr: false,
		},
		{
			name:    "Valid command - with faces",
			input:   "UwU -faces Hello World",
			want:    "Hello World",
			stutter: false,
			faces:   true,
			noNya:   false,
			wantErr: false,
		},
		{
			name:    "Valid command - with no-nya",
			input:   "UwU -no-nya Hello World",
			want:    "Hello World",
			stutter: false,
			faces:   false,
			noNya:   true,
			wantErr: false,
		},
		{
			name:    "Invalid command - no prefix",
			input:   "Hello World",
			wantErr: true,
		},
		{
			name:    "Invalid command - empty content",
			input:   "UwU ",
			wantErr: true,
		},
		{
			name:    "Invalid command - just prefix",
			input:   "UwU",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg, err := NewMessageFromStr(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMessageFromStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if msg.Content != tt.want {
				t.Errorf("NewMessageFromStr() Content = %v, want %v", msg.Content, tt.want)
			}
			if msg.Flags.Stutter != tt.stutter {
				t.Errorf("NewMessageFromStr() Stutter = %v, want %v", msg.Flags.Stutter, tt.stutter)
			}
			if msg.Flags.Faces != tt.faces {
				t.Errorf("NewMessageFromStr() Faces = %v, want %v", msg.Flags.Faces, tt.faces)
			}
			if msg.Flags.NoNya != tt.noNya {
				t.Errorf("NewMessageFromStr() NoNya = %v, want %v", msg.Flags.NoNya, tt.noNya)
			}
		})
	}
}

func TestUwuify(t *testing.T) {
	tests := []struct {
		name    string
		content string
		flags   UwuFlags
		want    string
	}{
		{
			name:    "Basic transformation",
			content: "Road Roller",
			flags:   UwuFlags{},
			want:    "Woad Wowwew",
		},
		{
			name:    "Nya transformation",
			content: "Night",
			flags:   UwuFlags{},
			want:    "nyaaght",
		},
		{
			name:    "Stutter transformation",
			content: "Hello",
			flags:   UwuFlags{Stutter: true},
			want:    "H-Hewwo",
		},
		{
			name:    "Repeated vowels",
			content: "Noooo wayyyy",
			flags:   UwuFlags{NoNya: true},
			want:    "No way",
		},
		{
			name:    "Protection - URL",
			content: "Check this: https://google.com",
			flags:   UwuFlags{},
			want:    "Check this: https://google.com",
		},
		{
			name:    "Protection - Mention",
			content: "Hello <@123456789>",
			flags:   UwuFlags{},
			want:    "Hewwo <@123456789>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MessageToUwuify{
				Content: tt.content,
				Flags:   tt.flags,
			}
			got := m.Uwuify()
			if tt.want != "" && got != tt.want {
				t.Errorf("Uwuify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUwuifyFaces(t *testing.T) {
	m := &MessageToUwuify{
		Content: "Hello",
		Flags:   UwuFlags{Faces: true},
	}
	got := m.Uwuify()

	// Basic check that it contains a face at the end
	faces := []string{"uwu", ">w<", "^w^", ":3", "owo", "x3", "rawr"}
	found := false
	for _, f := range faces {
		// Output is "Hewwo " + face
		if got == "Hewwo "+f {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Uwuify() with faces = %v, didn't match any expected face suffix", got)
	}
}
