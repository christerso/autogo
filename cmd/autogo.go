package main

import (
	"fmt"
	"github.com/christerso/autogo/cmd"
	cfg "github.com/christerso/autogo/config"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		cmd.printUsage()
		os.Exit(1)
	}

	config, err := cfg.loadConfig()
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		os.Exit(1)
	}

	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Path is required for add command.")
			os.Exit(1)
		}
		path := os.Args[2]
		cmd.addPath(config, path)
	case "jump":
		if len(os.Args) < 3 {
			fmt.Println("Query is required for jump command.")
			os.Exit(1)
		}
		query := os.Args[2]
		cmd.jump(config, query)
	default:
		cmd.printUsage()
	}
}
