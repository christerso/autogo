package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"

	"github.com/christerso/autojump/config"
	"github.com/rs/zerolog/log"
)

func AddPath(config *config.Config, path string) {
	if _, exists := config.Paths[path]; exists {
		config.Paths[path]++
	} else {
		config.Paths[path] = 1
	}
	if err := config.SaveConfig(config); err != nil {
		log.Fatal().Err(err).Msg("Failed to save config")
	}
}

func Jump(config *config.Config, query string) {
	candidates := []string{}
	for path := range config.Paths {
		if strings.Contains(path, query) {
			candidates = append(candidates, path)
		}
	}

	if len(candidates) == 0 {
		fmt.Println("No matching directories found.")
		return
	}

	sort.Slice(candidates, func(i, j int) bool {
		return config.Paths[candidates[i]] > config.Paths[candidates[j]]
	})

	selectedPath := candidates[0]
	fmt.Println(selectedPath)
	changeDirectory(selectedPath)
}

func changeDirectory(path string) {
	cmd := exec.Command("cmd", "/C", "cd /d", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("Error changing directory")
	}
}

func PrintUsage() {
	fmt.Println("Usage:")
	fmt.Println("  autojump add <path>  - Add the specified directory to the database")
	fmt.Println("  autojump jump <query> - Jump to the most used directory matching the query")
}
