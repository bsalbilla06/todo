package main

import (
	"bytes"
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

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

func CompleteTask(projectName, arg string) {
	taskID, err := strconv.Atoi(arg)
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile(projectName, os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w := csv.NewWriter(file)

	tasks := ExtractTasks(projectName)
	tasks[taskID-1].Completed = true
	
	strs := make([][]string, len(tasks))
	for i := range tasks {
		strs[i] = tasks[i].String()
	}
	w.WriteAll(strs)
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

func PrioritizeTask(projectName, arg string) {
	taskID, err := strconv.Atoi(arg)
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile(projectName, os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w := csv.NewWriter(file)

	tasks := ExtractTasks(projectName)
	tasks[taskID-1].Prioritized = !tasks[taskID-1].Prioritized
	
	strs := make([][]string, len(tasks))
	for i := range tasks {
		strs[i] = tasks[i].String()
	}
	w.WriteAll(strs)
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

func DeleteTask(projectName, arg string) {
	taskID, err := strconv.Atoi(arg)
	if err != nil {
		panic(err)
	}
	taskID--

	tasks := ExtractTasks(projectName)
	tasks = append(tasks[:taskID], tasks[taskID+1:]...)

	file, err := os.OpenFile(projectName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

    for _, task := range tasks {
        if err := w.Write(task.String()); err != nil {
			panic(err)
        }
    }
}

/*
todo is a tool to help manage tasks

Usage:
	todo <command> [arguments]

The commands are:
	new			create a new task
	complete	complete a new task
	prioritize	prioritize a task
	switch		switch current project
	delete		delete a task
*/
func HelpUser() string {
	return "todo is a tool to help manage tasks\n\nUsage:\n\ttodo <command> [arguments]\n\nThe commands are:\n\tnew\t\tcreate a new task\n\tcomplete\tcomplete a new task\n\tprioritize\tprioritize a task\n\tswitch\t\tswitch current project\n\tdelete\t\tdelete a task\n"
}