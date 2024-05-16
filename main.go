package main

import (
	"github.com/mlange-42/beecs-cli/cli"
	"github.com/mlange-42/beecs-cli/registry"
	"github.com/mlange-42/beecs-template/params"
	"github.com/mlange-42/beecs-template/sys"
)

func main() {
	// Register your custom parameters, resources and systems
	registry.RegisterResource[params.TestParams]()
	registry.RegisterSystem[sys.PrintSystem]()

	// Run the command line app
	cli.Run()
}
