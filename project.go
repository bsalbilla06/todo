package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func SwitchProject(projectName string, configBytes []byte) {
	projectFilePath :=fmt.Sprintf("%v%v.csv", ConfigDir(), projectName)

	// switch project in config
	config, err := os.OpenFile(ConfigDir() + ConfigFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer config.Close()

	configData := GetConfig(configBytes)
	configData.CurrentProject = projectFilePath

	newBytes, err := json.Marshal(configData)
	if err != nil {
		panic(err)
	}

	_, err = config.Write(newBytes)
	if err != nil {
		panic(err)
	}

	// ensure project esists, if not create a new one
	_, err = os.ReadFile(projectFilePath)
	if err != nil {
		if err.Error() == fmt.Sprintf("open %v: no such file or directory", projectFilePath) {
			CreateProject(projectName)
		} else {
			panic(err)
		}
	}
}

func CreateProject(projectName string) {
	filepath := fmt.Sprintf("%v%v.csv", ConfigDir(), projectName)
	project, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
	project.Close()
}