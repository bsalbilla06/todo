package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	configFile = "config.json"
	projectFile = "project1.csv"
)

type ConfigData struct {
	CurrentProject string `json:"currentProject"`
}

func CreateConfig(userConfigDir string) {
	// 0755 breakdown:
	// - Leading Zero (0): Indicates that the number is in octal (base 8).
	// - Owner Permissions (7): The owner of the file (the user who created the directory) has
	//   read (4), write (2), and execute (1) permissions. Adding these up (4 + 2 + 1 = 7) gives the owner full permissions.
	// - Group Permissions (5): Members of the file's group have read (4) and execute (1) permissions.
	//   Adding these up (4 + 1 = 5) gives the group read and execute permissions.
	// - Others Permissions (5): All other users have read (4) and execute (1) permissions.
	//   Adding these up (4 + 1 = 5) gives others read and execute permissions.
	err := os.Mkdir(userConfigDir, 0755)
	if err != nil {
		panic(err)
	}

	// Creating config file
	file, err := os.Create(userConfigDir + configFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Creating default project
	f, err := os.Create(userConfigDir + projectFile)
	if err != nil {
		panic(err)
	}
	f.Close()

	configData := ConfigData{CurrentProject: fmt.Sprintf("%v/%v", userConfigDir, projectFile)}
	bytes, err := json.Marshal(configData)
	if err != nil {
		panic(err)
	}
	_, err = file.Write(bytes)
	if err != nil {
		panic(err)
	}
}

func ConfigDir() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%v/todo/", configDir)
}

func CheckForConfig(userConfigDir string) error {
	_, err := os.ReadFile(userConfigDir + configFile)
	if err != nil {
		if err.Error() == fmt.Sprintf("open %v%v: no such file or directory", userConfigDir, configFile) {
			return fmt.Errorf("config not found")
		} else {
			return err
		}
	} else {
		return nil
	}
}

// func GetConfig() ConfigData
