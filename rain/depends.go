package rain

import (
	"fmt"
	"os/exec"
)

var AllDepends = []string{
	"github.com/RyanPrintup/Nimbus",
	"github.com/chyzer/readline",
	"gopkg.in/yaml.v2",
}

var CoreDepends = []string{
	"github.com/RyanPrintup/Nimbus",
	"gopkg.in/yaml.v2",
}

var NeverfreeDepends = []string{
	"github.com/chyzer/readline",
}

func GetDepends(depends []string) error {
	for _, dep := range depends {
		output, err := exec.Command("go", "get", "-u", dep).CombinedOutput()
		if err != nil {
			fmt.Printf(" Internal command error\n ----\n%s\n ----\n\n", string(output[:]))
			return err
		}
	}
	return nil
}
