package commands

import (
	"fmt"
	"os"
)

func CommandExit(string) error {
	fmt.Println("Exiting Rally Maker...")
	os.Exit(0)
	return nil
}
