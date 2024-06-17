package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func main() {
	userConfigDir := ConfigDir()

	err := CheckForConfig(userConfigDir)
	if err != nil {
		if err.Error() == "config not found" {
			CreateConfig(userConfigDir)
		} else {
			panic(err)
		}
	}

	configBytes, err := os.ReadFile(userConfigDir + "config.json")
	if err != nil {
		panic(err)
	}

	configData := new(ConfigData)
	json.Unmarshal(configBytes, configData)
	project := configData.CurrentProject

	command := GetArg(1)
	switch command {
	case "new":
		WriteNewTask(Task{Instruction: GetArg(2), Prioritized: false, Completed: false}, project)
	case "complete":
		taskID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			panic(err)
		}
		CompleteTask(taskID, project)
	case "prioritize":
		taskID, err := strconv.Atoi(GetArg(2))
		if err != nil {
			panic(err)
		}
		PrioritizeTask(taskID, project)
	case "help":
		fmt.Println(HelpUser())
	// case "switch":
	// 	break
	// case "delete":
	// 	break
	default:
		tasks := ExtractTasks(project)
		if len(tasks) < 1 {
			fmt.Println("-----ALL TASKS COMPLETED-----")
		} else {
			fmt.Println("-----TODO-----")
			for i, v := range tasks {
				if !v.Completed && v.Prioritized {
					fmt.Println("*", i+1, v.Instruction)
				}
			}
			for i, v := range tasks {
				if !v.Completed && !v.Prioritized {
					fmt.Println(i+1, v.Instruction)
				}
			}
		}
	}
}