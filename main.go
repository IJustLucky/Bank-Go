package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var a account

	reader := bufio.NewReader(os.Stdin)
	for {
		//Prompt for user input
		fmt.Print("--> ")

		//Read user input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}
		input = strings.TrimSpace(input)

		//Process user input
		inputs := strings.Fields(input)
		if len(inputs) < 1 {
			continue
		}

		//Take action!
		switch inputs[0] {
		case "withdraw":
			if len(inputs) < 2 {
				fmt.Println("you need to add a value to your withdraw")
				break
			}
			str := inputs[1]

			result, err := strconv.ParseFloat(str, 32)
			if err != nil {
				fmt.Println("can only withdraw numerical values")
				break
			}

			if err = a.withdraw(result); err != nil {
				fmt.Println(err)
			}
		case "deposit":
			if len(inputs) < 2 {
				fmt.Println("you need to add a value to your deposit")
				break
			}
			str := inputs[1]

			result, err := strconv.ParseFloat(str, 32)
			if err != nil {
				fmt.Println("can only deposit numerical values")
				break
			}

			err = a.deposit(result)
			if err != nil {
				fmt.Println(err)
			}
		case "balance":
			fmt.Println(a.balance())
		}
	}
}
