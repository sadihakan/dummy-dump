package util

import "os"

func checkCommands(commands []string) (command string) {
	for _, c := range commands {
		if _, err := os.Stat(c); err == nil {
			command = c
			break
		}
	}

	return command
}
