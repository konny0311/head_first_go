package main

import (
	"math/rand"
	"strconv"
	"strings"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var targetNum int64 = int64(rand.Intn(100) + 1)
	targetNum = 10
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 10; i++ {
		fmt.Println("You have ", 10 - i, "chances left.\n Input your guess")
		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if err != nil {
			log.Fatal(err)
		}
		inputInt, err := strconv.ParseInt(input, 0, 0)
		if err != nil {
			log.Fatal(err)
		}
		if targetNum == inputInt {
			fmt.Println("Congrats! You win")
			return 
		} else {
			if inputInt >= targetNum {
				fmt.Println("Guess smaller")
			} else {
				fmt.Println("Guess bigger")
			}
		}
	}
	fmt.Println("Sorry, you lose..")
}