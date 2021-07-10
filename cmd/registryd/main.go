package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/hleb-albau/registry/app"
	"github.com/hleb-albau/registry/cmd/registryd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
