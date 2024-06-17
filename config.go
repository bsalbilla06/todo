package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	ConfigFile = "config.json"
)

type ConfigData struct {
	CurrentProject string `json:"currentProject"`
}

func ConfigDir() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%v/todo/", configDir)
}

func CreateConfig() []byte {
	// don't need to check error because it doesn't matter if the directory already exists

	// 0755 breakdown:
	// - Leading Zero (0): Indicates that the number is in octal (base 8).
	// - Owner Permissions (7): The owner of the file (the user who created the directory) has
	//   read (4), write (2), and execute (1) permissions. Adding these up (4 + 2 + 1 = 7) gives the owner full permissions.
	// - Group Permissions (5): Members of the file's group have read (4) and execute (1) permissions.
	//   Adding these up (4 + 1 = 5) gives the group read and execute permissions.
	// - Others Permissions (5): All other users have read (4) and execute (1) permissions.
	//   Adding these up (4 + 1 = 5) gives others read and execute permissions.
	os.Mkdir(ConfigDir(), 0755)

	// Creating config file
	config, err := os.Create(ConfigDir() + ConfigFile)
	if err != nil {
		panic(err)
	}
	defer config.Close()

	CreateProject("project1")

	configData := ConfigData{CurrentProject: fmt.Sprintf("%vproject1.csv", ConfigDir())}
	bytes, err := json.Marshal(configData)
	if err != nil {
		panic(err)
	}
	_, err = config.Write(bytes)
	if err != nil {
		panic(err)
	}

	return bytes
}

func GetConfig(configBytes []byte) *ConfigData {
	ret := new(ConfigData)
	err := json.Unmarshal(configBytes, ret)
	if err != nil {
		panic(err)
	}
	return ret
}