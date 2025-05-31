package main

import "fmt"

type Task struct {
	Name string
	DueDate string
	Priority string
	IsDone bool
}

func AddTask() {
	var newTask string
	var priority string
	var duedate string
	var tasks []Task
	var more string
	for {
		fmt.Print("\n Enter Task's name: ")
		fmt.Scanln(&newTask)

		fmt.Print("\n Enter Task's priority:  ")
		fmt.Scanln(&priority)

		fmt.Print("\n Enter Task's due date: ")
		fmt.Scanln(&duedate)

		tasks = append(tasks, Task{
			Name: newTask,
			Priority: priority,
			DueDate: duedate,
		})

		fmt.Print("\n Do you want to add more tasks: ")
		fmt.Scanln(&more)

		if more != "yes" {
			break
		}
	}
}