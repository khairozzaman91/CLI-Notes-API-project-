package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"cli-notes-api/database"
	"cli-notes-api/models"
)

func main() {
	database.ConnectDB()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== CLI Notes API ===")
		fmt.Println("1. Create Note")
		fmt.Println("2. Update Note")
		fmt.Println("3. Show All Notes")
		fmt.Println("4. Delete Note")
		fmt.Println("5. Exit")
		fmt.Print("Enter choice: ")

		input, _ := reader.ReadString('\n')
		switch strings.TrimSpace(input) {
		case "1":
			fmt.Print("Enter title: ")
			title, _ := reader.ReadString('\n')
			fmt.Print("Enter content: ")
			content, _ := reader.ReadString('\n')
			note := models.Note{
				Title:   strings.TrimSpace(title),
				Content: strings.TrimSpace(content),
			}
			database.DB.Create(&note)
			fmt.Println("Note created.")

		case "2":
			fmt.Print("Note ID to update: ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))
			var note models.Note
			if err := database.DB.First(&note, id).Error; err != nil {
				fmt.Println("Not found"); continue
			}
			fmt.Print("New title: ")
			t, _ := reader.ReadString('\n')
			fmt.Print("New content: ")
			c, _ := reader.ReadString('\n')
			note.Title = strings.TrimSpace(t)
			note.Content = strings.TrimSpace(c)
			database.DB.Save(&note)
			fmt.Println("Note updated.")

		case "3":
			var notes []models.Note
			database.DB.Find(&notes)
			fmt.Println("\n--- Notes ---")
			for _, n := range notes {
				fmt.Printf("ID:%d Title:%s\n%s\n\n", n.ID, n.Title, n.Content)
			}

		case "4":
			fmt.Print("Note ID to delete: ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))
			database.DB.Delete(&models.Note{}, id)
			fmt.Println("Deleted.")

		case "5":
			fmt.Println("Bye!")
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}
