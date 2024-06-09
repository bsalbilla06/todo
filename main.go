package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	const project = "project2.csv"
	// check if the user passed any args
	if len(os.Args) > 1 {
		// check which command the user entered
		newArgIndex := CheckArg("new", 1)
		if newArgIndex > 0 {
			WriteNewTask(Task{Instruction: os.Args[newArgIndex+1], Prioritized: false, Completed: false}, project)
		}

		completeArgIndex := CheckArg("complete", 1)
		if completeArgIndex > 0 {
			taskID, err := strconv.Atoi(os.Args[2])
			if err != nil {
				panic(err)
			}
			CompleteTask(taskID, project)
		}

		prioritizeArgIndex := CheckArg("prioritize", 1)
		if prioritizeArgIndex > 0 {
			taskID, err := strconv.Atoi(os.Args[2])
			if err != nil {
				panic(err)
			}
			PrioritizeTask(taskID, project)
		}

		helpArgIndex := CheckArg("help", 1)
		if helpArgIndex > 0 {
			HelpUser()
		}

	} else {
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
