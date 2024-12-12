package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func connectToServer() net.Conn {
	server := os.Args[1]
	port := os.Args[2]

	addr := fmt.Sprintf("%s:%s", server, port)
	// Connect to the server
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	return conn
}

func sendMessage(conn net.Conn, message string, username string) {
	data := fmt.Sprintf("%s: %s", username, message)
	_, err := conn.Write([]byte(data))
	if err != nil {
		log.Println("Error sending data to server:", err)
		os.Exit(1)
	}
}

func main() {
	// Create a new application
	app := tview.NewApplication()

	// Connect to the server
	conn := connectToServer()

	// Create a channel to receive messages from the server
	messageChannel := make(chan string)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter your username: ")
	scanner.Scan()
	username := scanner.Text()

	sendMessage(conn, "has joined the chat", username)

	// Start a goroutine to read messages from the server
	go func() {
		for {
			buf := make([]byte, 4096)
			n, err := conn.Read(buf)
			if err != nil {
				log.Println("Error reading data from server:", err)
				return
			}
			if n != 0 {
				messageChannel <- string(buf[:n])
			}
		}
	}()

	// Create a text view for displaying messages
	textView := tview.NewTextView().SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() {
			app.Draw()
		})

	textView.SetTitle(" CLi-Chat ").SetBorder(true).SetTitleAlign(tview.AlignCenter)

	// Create an input field for user input
	inputField := tview.NewInputField().
		SetLabel("Enter your message: ").SetAcceptanceFunc(nil)

	inputField.SetFieldBackgroundColor(tcell.ColorBlack)
	inputField.SetFieldTextColor(tcell.ColorWhite)
	inputField.SetLabelColor(tcell.Color171)

	inputField.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			// Get the user's input
			userInput := inputField.GetText()
			// Append the user's input to the output view
			fmt.Fprintf(textView, "[green]you: %s", userInput+"\n")
			// Clear the input field
			inputField.SetText("")
			data := fmt.Sprintf("%s: %s", username, userInput)
			_, writeErr := conn.Write([]byte(data))
			if writeErr != nil {
				log.Println("Error writing to server:", writeErr)
				os.Exit(1)
			}
		}
	})

	copyright := tview.NewTextView()

	copyright.Write([]byte("Developer: github.com/ThawThuHan\n"))

	copyright.SetTextAlign(tview.AlignCenter)

	// Layout combining the output view and input field
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(textView, 0, 1, false).
		AddItem(inputField, 3, 1, true).
		AddItem(copyright, 1, 0, false)

	// Set the layout as the root and run the application
	go func() {
		for {
			message := <-messageChannel
			fmt.Fprintf(textView, "[yellow]%s", message+"\n")
			app.Draw()
		}
	}()

	if err := app.SetRoot(layout, true).Run(); err != nil {
		log.Fatalf("Error running application: %v", err)
	}
}
