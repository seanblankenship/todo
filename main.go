package main

import (
	"fmt"
	"net/http"
)

var taskItems = []string{
	"- add to list",
	"- delete from list",
	"- serve formatted html and css",
	"- store / preserve list over time",
	"- track multiple lists",
}

func main() {
	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)
	http.ListenAndServe(":8080", nil)
}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	for _, task := range taskItems {
		fmt.Fprintf(writer, "%s\n", task)
	}
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	greeting := "Hello, User. Welcome."
	fmt.Fprintf(writer, greeting)
}

func printTasks(taskItems []string) {
	fmt.Println("List One")
	for index, task := range taskItems {
		fmt.Printf("%d: %s\n", index+1, task)
	}
}

func addTask(taskItems []string, newTask string) []string {
	updatedTaskItems := append(taskItems, newTask)
	return updatedTaskItems
}
