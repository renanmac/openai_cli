package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/renanmac/openai_cli/internal/client"
	"github.com/renanmac/openai_cli/internal/infra"
)

func main() {
	var response string

	chatCmd := flag.NewFlagSet("chat", flag.ExitOnError)
	promptPtr := chatCmd.String("prompt", "", "Prompt to be used in the chat function")

	stat, err := os.Stdin.Stat()

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		var buffer []byte
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			buffer = append(buffer, scanner.Bytes()...)
		}

		response, err = client.Chat(infra.Requests{}, string(buffer))
	} else {
		switch os.Args[1] {
		case "chat":
			chatCmd.Parse(os.Args[2:])
			response, err = client.Chat(infra.Requests{}, *promptPtr)
		default:
			fmt.Println("Command not found")
		}
	}

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(response)
}
