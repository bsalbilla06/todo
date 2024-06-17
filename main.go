package main

import (
	"fmt"
	"os"
)

func main() {
	configBytes, err := os.ReadFile(ConfigDir() + "config.json")
	if err != nil {
		if err.Error() == fmt.Sprintf("open %v%v: no such file or directory", ConfigDir(), ConfigFile) {
			configBytes = CreateConfig()
		} else {
			panic(err)
		}
	}

	configData := GetConfig(configBytes)
	project := configData.CurrentProject

	command, _ := GetArg(1)
	switch command {
	case "new":
		instruction, err := GetArg(2)
		if err != nil {
			panic(err)
		}
		WriteNewTask(Task{Instruction: instruction, Prioritized: false, Completed: false}, project)
	case "complete":
		taskID, err := GetArg(2)
		if err != nil {
			panic(err)
		}
		CompleteTask(project, taskID)
	case "prioritize":
		taskID, err := GetArg(2)
		if err != nil {
			panic(err)
		}
		PrioritizeTask(project, taskID)
	case "help":
		fmt.Println(HelpUser())
	case "switch":
		projectName, err := GetArg(2)
		if err != nil {
			panic(err)
		}
		SwitchProject(projectName, configBytes)
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