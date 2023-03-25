package main

import (
	"fmt"
)

//Struct that keeps the details every todo item must have
type todoItem struct {
	id   int
	name string
}


var items []todoItem //empty todoItem array
var idCounter = 0

//function to add an item 
func addItem(name string) {
	idCounter++
	item := todoItem{id: idCounter, name: name}
	items = append(items, item)
}

// function to delete an item
func deleteItem(id int) {
	for i, item := range items {
		if item.id == id {
			items = append(items[:i], items[i+1:]...)
			return
		}
	}
}

// function to display an item
func displayItems() {
	for _, item := range items {
		fmt.Printf("ID: %d, Name: %s\n", item.id, item.name)
	}
}

func main() {
	fmt.Println("Welcome to your todo list\n")

	for {
		fmt.Println("Please select an option")
		fmt.Println("1. Add item")
		fmt.Println("2. Remove item")
		fmt.Println("3. Display items")
		fmt.Println("4. Quit")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var name string
			fmt.Println("Enter the name of the item to be added:")
			fmt.Scanln(&name)
			addItem(name)
			fmt.Println("Item has been added.")
		case 2:
			var id int
			fmt.Println("Enter the ID of the item to be removed:")
			fmt.Scanln(&id)
			deleteItem(id)
			fmt.Println("Item has been removed.")
		case 3:
			displayItems()
		case 4:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

		fmt.Println()
	}
}
