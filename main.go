package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/lfernandezlo/hackrn/model"
)

func main() {
	rootFolder := model.Folder{Name: "root", Path: "/root"}

	d := model.Directory{
		Folders:          []*model.Folder{&rootFolder},
		CurrentFolder:    &rootFolder,
		CurrentDirectory: "/root",
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()

		command := strings.Split(text, " ")

		if len(command) == 0 {
			fmt.Println("Should specify command, types can be mkdir, touch, ls, cd, pwd or quit. i.e. mkdir test")
		}

		commandType := command[0]

		var argument string

		if len(command) > 1 && command[1] != "" {
			argument = command[1]
		}

		var res interface{}
		var err error

		switch commandType {
		case "mkdir":
			if argument == "" {
				fmt.Println("Should specify folder name i.e. mkdir test")
				break
			}

			_, err = d.Mkdir(argument)
		case "touch":
			if argument == "" {
				fmt.Println("Should specify file name i.e. touch file")
				break
			}

			_, err = d.Touch(argument)
		case "ls":
			res = d.Ls()
		case "cd":
			if argument == "" {
				fmt.Println("Should specify folder name i.e. cd folder")
				break
			}

			_, err = d.Cd(argument)
		case "pwd":
			res = d.Pwd()
		case "quit":
			fmt.Println("exit")
			d.Quit(0)
		default:
			fmt.Println("Unfound command, types can be mkdir, touch, ls, cd, pwd or quit. i.e. mkdir test")
		}

		if err != nil {
			fmt.Println(err.Error())
		}

		if res != nil && res != "" {
			fmt.Println(res)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(fmt.Sprintf("Error on scanner: %v", err.Error()))
		d.Quit(1)
	}
}
