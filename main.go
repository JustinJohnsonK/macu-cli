package main

import (
	"fmt"

	"github.com/JustinJohnsonK/macu-cli/cmd/root"
	"github.com/JustinJohnsonK/macu-cli/cmd/systemd"
)

func main() {
	root_cmd := root.RootCmd()

	root_cmd.AddCommand(systemd.NewFile())

	if err := root_cmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
