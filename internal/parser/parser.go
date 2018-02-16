package parser

import "errors"
import "github.com/raindevteam/rain/internal/handler"
import "strings"

// ParseCommand takes an input string and parses it into a CommandData struct.
// This struct can be fed to the handler for reading. If the input string is not
// a valid command, an error is returned.
func ParseCommand(input string) (*handler.CommandData, error) {
	// Ensure string starts with '!'
	if input[0] != '!' {
		return nil, errors.New("input string not a valid command")
	}

	cd := &handler.CommandData{}

	// Check if its internal
	if input[1] == '!' {
		cd.Owner = handler.Internal
		cd.Command = input[2:]
		// Check if droplet specified
	} else if strings.Contains(input, ":") {
		result := strings.Split(input[1:], ":")
		cd.Owner = result[0]
		cd.Command = result[1]
	} else {
		cd.Owner = ""
		cd.Command = input[1:]
	}

	return cd, nil
}
