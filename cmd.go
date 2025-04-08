package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func CammandsInit() {
	allTasks := &[]TaskTracker{}
	Init(allTasks)

	if len(os.Args) < 2 {
		fmt.Println("Expected a command")
		return
	}

	command := os.Args[1]
	suffixCammand := ""
	suffixCammand2 := ""
	if len(os.Args) >= 3 {
		suffixCammand = os.Args[2]
	}
	if len(os.Args) == 4 {
		suffixCammand2 = os.Args[3]
	}
	switch command {
	case list:
		excuteListCammand(*allTasks, suffixCammand)
	case add:
		excuteAddCammand(*allTasks, suffixCammand)
	case update:
		excuteUpdateCammand(*allTasks, suffixCammand2, suffixCammand)
	case markInProgress:
		excuteStatusUpdateCammand(*allTasks, command, suffixCammand)
	case markDone:
		excuteStatusUpdateCammand(*allTasks, command, suffixCammand)
	case delete:
		excuteDeleteCammand(*allTasks, suffixCammand)
	}
}

func excuteDeleteCammand(taskTracker []TaskTracker, suffixCammand string) {
	id, err := strconv.Atoi(suffixCammand)
	if err != nil {
		fmt.Println(
			"❌ Invalid command!\n\nUsage:\n  task-cli delete <task_id> \n\nExample:\n  task-cli delete 1",
		)
		
		return
	}
	Delete(&taskTracker, id)
}

func excuteStatusUpdateCammand(taskTracker []TaskTracker, command string, suffixCammand string) {
	status := -1
	if command == markInProgress {
		status = StatusInProgress
	}
	if command == markDone {
		status = StatusDone
	}

	task := TaskTracker{}
	if status != -1 {
		task.Status = status
	}

	id, err := strconv.Atoi(suffixCammand)
	if err != nil || suffixCammand == "" {
		fmt.Println(
			"❌ Invalid command!\n\nUsage:\n  task-cli <command> <task_id>",
		)
		return
	}
	task.Update(&taskTracker, id)
}

func excuteUpdateCammand(taskTracker []TaskTracker, suffixCammand string, suffixCammand2 string) {
	task := TaskTracker{}
	task.Description = suffixCammand
	id, err := strconv.Atoi(suffixCammand2)
	if err != nil || suffixCammand == "" {
		fmt.Println(
			"❌ Invalid command!\n\nUsage:\n  task-cli update <task_id> \"<new description>\"\n\nExample:\n  task-cli update 1 \"Buy groceries and cook dinner\"",
		)

		return
	}
	task.Update(&taskTracker, id)
}
func excuteAddCammand(taskTracker []TaskTracker, suffixCammand string) {
	task := TaskTracker{}
	task.Description = suffixCammand
	task.Save(&taskTracker)
}

func excuteListCammand(allTasks []TaskTracker, suffixCammand string) {
	listType := -1
	switch suffixCammand {
	case "done":
		listType = StatusDone

	case "todo":
		listType = StatusTodo

	case "in-progress":
		listType = StatusInProgress

	}
	s, err := json.MarshalIndent(GetList(listType, allTasks), "", "  ")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(string(s))

}
