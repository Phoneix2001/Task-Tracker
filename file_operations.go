package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"os"
)

func Init(allTasks *[]TaskTracker) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("File Doesn't Exist")
		os.WriteFile(filePath, make([]byte, 0), 0644)
	}

	errs := json.Unmarshal(data, allTasks)
	if errs != nil {
		log.Print(err)
	}

}

func (s *TaskTracker) Save(tasks *[]TaskTracker) {

	// Updating Values
	s.ID = len(*tasks) + 1
	s.UpdatedAt = time.Now().Format(time.DateTime)
	s.CreatedAt = time.Now().Format(time.DateTime)
	s.Status = StatusTodo
	// Step 1: Append the new tracker
	*tasks = append(*tasks, *s)

	// Step 2: Marshal updated slice
	updatedData, err := json.Marshal(tasks)
	if err != nil {
		log.Fatalf("Failed to marshal data: %v", err)
	}

	// Step 3: Write it back to the file
	err = os.WriteFile(filePath, updatedData, 0644)
	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}
	fmt.Printf("✅ Task saved successfully! [ID: %d, Description: %s]\n", s.ID, s.Description)
}

func (s *TaskTracker) Update(tasks *[]TaskTracker, id int) {
	for i, task := range *tasks {
		if task.ID == id {
			if s.Description != "" {
				(*tasks)[i].Description = s.Description
			}
			if s.Status != 0 {
				(*tasks)[i].Status = s.Status
			}
			(*tasks)[i].UpdatedAt = time.Now().Format(time.DateTime)
			break
		}
	}

	// Step 1: Marshal updated slice
	updatedData, err := json.Marshal(tasks)
	if err != nil {
		log.Fatalf("Failed to marshal data: %v", err)
	}

	// Step 2: Write it back to the file
	err = os.WriteFile(filePath, updatedData, 0644)
	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}
	fmt.Printf("✅ Task updated successfully! [ID: %d]\n", id)
}

func GetList(listType int, allTasks []TaskTracker) []TaskTracker {
	switch listType {
	case StatusDone:
		doneTodo := []TaskTracker{}
		for _, todo := range allTasks {
			if todo.Status == StatusDone {
				doneTodo = append(doneTodo, todo)
			}
		}
		return doneTodo
	case StatusInProgress:
		InProgressTodo := []TaskTracker{}
		for _, todo := range allTasks {
			if todo.Status == StatusInProgress {
				InProgressTodo = append(InProgressTodo, todo)
			}
		}
		return InProgressTodo
	case StatusTodo:
		intialTodo := []TaskTracker{}
		for _, todo := range allTasks {
			if todo.Status == StatusTodo {
				intialTodo = append(intialTodo, todo)
			}
		}
		return intialTodo
	default:
		return allTasks
	}
}

func Delete(taskTracker *[]TaskTracker, id int) {
	for index, task := range *taskTracker {
		if task.ID == id {
			// Remove task from slice
			*taskTracker = append((*taskTracker)[:index], (*taskTracker)[index+1:]...)
			fmt.Printf("✅ Task with ID %d deleted successfully.\n", id)
			break
		}
	}

	// Marshal updated slice
	updatedData, err := json.Marshal(*taskTracker)
	if err != nil {
		log.Fatalf("❌ Failed to marshal data: %v", err)
	}

	// Write updated data to file
	err = os.WriteFile(filePath, updatedData, 0644)
	if err != nil {
		log.Fatalf("❌ Failed to write to file: %v", err)
	}
}
