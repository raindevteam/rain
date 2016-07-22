package parser

import (
	"strings"
	"unicode"
)

func trimAll(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

// The Parser is used occasionaly throughout different parts of the bot. It mostly handles parsing
// irc messages for commands. The bot then takes what is parsed and is usually passed to the
// handler (in a form that it can read).
type Parser struct {
	Prefix string
}

// NewParser returns a new parser given a command prefix.
func NewParser(prefix string) *Parser {
	parser := &Parser{prefix}
	return parser
}

func (p *Parser) ParseModuleValue(value string) (string, []string) {
	cut := strings.SplitN(value, ":", 2)
	if len(cut) == 1 {
		return cut[0], []string{}
	}
	route := cut[0]
	options := strings.Split(cut[1], " ")
	return route, options
}

// IsCommand checks if an IRC message is a command.
func (p *Parser) IsCommand(text string) bool {
	return string(text[0]) == p.Prefix
}

// ParseCommand will return the name of the command and its argument list.
func (p *Parser) ParseCommand(raw string) (string, []string) {
	splitMessage := strings.Split(strings.Trim(string(raw[1:]), " "), " ")
	return splitMessage[0], splitMessage[1:]
}

// ParsePrefix takes an IRC prefix and returns the user and host.
func (p *Parser) ParsePrefix(raw string) (string, string) {
	if !strings.Contains(raw, "!") {
		return raw, ""
	}
	splitPrefix := strings.Split(raw, "!")
	return splitPrefix[0], splitPrefix[1]
}
