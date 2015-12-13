package rainbot

import "strings"

type Parser struct {
	Prefix string
}

func NewParser(prefix string) *Parser {
	parser := &Parser{prefix}
	return parser
}

func (p *Parser) IsCommand(text string) bool {
	return string(text[0]) == p.Prefix
}

func (p *Parser) ParseCommand(raw string) (string, []string) {
	splitMessage := strings.Split(string(raw[1:]), " ")
	return splitMessage[0], splitMessage[1:]
}
