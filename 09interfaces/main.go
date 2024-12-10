package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/zin-min-thu/note/note"
	"github.com/zin-min-thu/note/todo"
)

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

type outputtable interface {
	saver // embedded interface
	Display()
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

func main() {

	printSomething(1)
	printSomething(1.2)
	printSomething("Test string")

	title, content := getNoteDate()

	todoText := getTodoData()

	todo, err := todo.New(todoText)

	printSomething(todo)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		return
	}

	err = outputData(userNote)

	if err != nil {
		return
	}

}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed")
		return err
	}

	fmt.Println("Saving the note success")

	return nil
}

func getTodoData() string {
	return getUserInput("Todo text:")
}

func getNoteDate() (string, string) {
	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	// Scan can't accept longer input text
	// var value string
	// fmt.Scan(&value)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r") // remove special character

	return text
}

// interface type can access any kind of type value(int,float,string)
func printSomething(value interface{}) {

	intVal, ok := value.(int)
	if ok {
		fmt.Println("Integer: ", intVal)
	}

	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("Float: ", floatVal)
	}

	stringVal, ok := value.(string)
	if ok {
		fmt.Println("String: ", stringVal)
	}

	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer: ", value)
	// case float64:
	// 	fmt.Println("Float: ", value)
	// case string:
	// 	fmt.Println("String: ", value)
	// }

	// fmt.Println(value)
}
