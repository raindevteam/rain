package rain

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func doInstall(install string, name string) error {
	// Install Core Dependencies
	err := GetDepends(CoreDepends)
	if err != nil {
		return err
	}

	// Install Specific Dependencies
	switch install {
	case "neverfree":
		err = GetDepends(NeverfreeDepends)
	}

	if err != nil {
		return err
	}

	path := os.Getenv("GOPATH") + "/bin/" + name

	if runtime.GOOS == "windows" {
		path = path + ".exe"
	}

	output, err := exec.Command("go", "build", "-o", path,
		"github.com/raindevteam/rain/install/"+install).CombinedOutput()
	if err != nil {
		fmt.Printf(" Internal command error\n ----\n%s\n -----\n\n", string(output[:]))
		return err
	}

	return nil
}
