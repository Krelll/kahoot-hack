package main

import (
	"fmt"
	"github.com/unixpickle/kahoot-hack/kahoot"
	"os"
	"strconv"
)

func main() {
	var pin string
	var nickname string
	fmt.Print("868987: ")
	fmt.Scanln(&pin)
	fmt.Print("Sodapoppin: ")
	fmt.Scanln(&nickname)
	fmt.Println("Connecting...")
	conn, err := kahoot.NewConnection(pin)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	if err := conn.Register(nickname); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Awaiting questions...")
	for {
		if conn.WaitQuestion() != nil {
			fmt.Println("Done question loop:", err)
			os.Exit(1)
		}
		fmt.Print("Got question. Answer: ")
		var number string
		fmt.Scanln(&number)
		parsed, _ := strconv.Atoi(number)
		if err := conn.SendAnswer(parsed); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
