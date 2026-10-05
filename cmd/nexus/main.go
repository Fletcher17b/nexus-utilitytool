package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

type Config map[string]string

func loadConfig() (Config, error) {
	exe, err := os.Executable()
	if err != nil {
		return nil, err
	}

	cfgPath := filepath.Join(filepath.Dir(exe), "aliases.json")

	data, err := os.ReadFile(cfgPath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = json.Unmarshal(data, &cfg)
	return cfg, err
}

func main() {
	cfg, err := loadConfig()
	if err != nil {
		fmt.Println("Failed to load aliases:", err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  nexus <alias>")
		fmt.Println("  nexus list")
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {

	case "list":
		keys := make([]string, 0, len(cfg))
		for k := range cfg {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			fmt.Printf("%-15s %s\n", k, cfg[k])
		}

	default:
		path, ok := cfg[cmd]
		if !ok {
			fmt.Println("Unknown alias:", cmd)
			os.Exit(1)
		}

		fmt.Print(path)
	}
}
