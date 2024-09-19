package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/BurntSushi/toml"
)

type Account struct {
	Number string `toml:"number"`
	Role   string `toml:"role"`
	Name   string `toml:"name"`
}

type Config struct {
	Accounts []Account `toml:"account"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Provide account name or number")
		return
	}

	// Hardcoded path to the TOML file
	const tomlFile = "accounts.toml"

	var config Config
	if _, err := toml.DecodeFile(tomlFile, &config); err != nil {
		fmt.Println("Error decoding TOML file:", err)
		return
	}

	arg := os.Args[1]

	for _, acc := range config.Accounts {
		if acc.Number == arg || acc.Name == arg {
			awsUrl := fmt.Sprintf("https://signin.aws.amazon.com/switchrole?account=%s&roleName=%s&displayName=%s", acc.Number, acc.Role, acc.Name)
			fmt.Println("Opening URL:", awsUrl)
			err := exec.Command("open", awsUrl).Start()
			if err != nil {
				fmt.Println("Failed to open URL:", err)
			}
			return
		}
	}

	fmt.Println("Account not found")
}
