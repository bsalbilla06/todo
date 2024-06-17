package main

import (
	"bytes"
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

// Since taskID is just the tasks line number, there is no reason to store it
type Task struct {
	Instruction string
	Prioritized	bool
	Completed   bool
}

func (t Task) String() []string {
	instruction := t.Instruction
	priority := strconv.FormatBool(t.Prioritized)
	completed := strconv.FormatBool(t.Completed)

	return []string{instruction, priority, completed}
}

func ExtractTasks(projectName string) []Task {
	data, err := os.ReadFile(projectName)
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(bytes.NewReader(data))

	tasks := make([]Task, 0, 10)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		instruction := record[0]
		priority, err := strconv.ParseBool(record[1])
		if err != nil {
			panic(err)
		}
		completed, err := strconv.ParseBool(record[2])
		if err != nil {
			panic(err)
		}

		tasks = append(tasks, Task{Instruction: instruction, Prioritized: priority, Completed: completed})
	}
	return tasks
}

func WriteNewTask(task Task, projectName string) {
	file, err := os.OpenFile(projectName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w := csv.NewWriter(file)

	instruction := task.Instruction
	priority := strconv.FormatBool(task.Prioritized)
	completed := strconv.FormatBool(task.Completed)

	w.Write([]string{instruction, priority, completed})
	w.Flush()
}

func CompleteTask(taskId int, projectName string) {
	file, err := os.OpenFile(projectName, os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w := csv.NewWriter(file)

	tasks := ExtractTasks(projectName)
	tasks[taskId-1].Completed = true
	
	strs := make([][]string, len(tasks))
	for i := range tasks {
		strs[i] = tasks[i].String()
	}
	w.WriteAll(strs)
}

func PrioritizeTask(taskId int, projectName string) {
	file, err := os.OpenFile(projectName, os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w := csv.NewWriter(file)

	tasks := ExtractTasks(projectName)
	tasks[taskId-1].Prioritized = !tasks[taskId-1].Prioritized
	
	strs := make([][]string, len(tasks))
	for i := range tasks {
		strs[i] = tasks[i].String()
	}
	w.WriteAll(strs)
}

// goal --help output, very similar to "go help"
/*
todo is a tool to help manage tasks

Usage:
	todo <command> [arguments]

The commands are:
	new			create a new task
	complete	complete a new task
	prioritize	prioritize a task

Commands coming soon:
	switch		switch current project
	delete		delete a task
*/

func HelpUser() string {
	return "todo is a tool to help manage tasks\n\nUsage:\n\ttodo <command> [arguments]\n\nThe commands are:\n\tnew\t\tcreate a new task\n\tcomplete\tcomplete a new task\n\tprioritize\tprioritize a task\n\nCommand(s) coming soon:\n\tswitch\t\tswitch current project\n\tdelete\t\tdelete a task\n"
}