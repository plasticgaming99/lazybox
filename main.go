package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		//quit with error
		os.Exit(1)
	}
	if os.Args[1] == "cd" {
		chdir(os.Args[2])
	}
	if os.Args[1] == "chdir" {
		chdir(os.Args[2])
	}
	if os.Args[1] == "echo" {
		echo(os.Args[2])
	}
	if os.Args[1] == "yes" {
		yes()
	}
	if os.Args[1] == "pwd" {
		pwd()
	}
	if os.Args[1] == "ls" {
		if len(os.Args) >= 3 {
			dir := os.Args[2]
			ls(dir)
		} else {
			dir, err := os.Getwd()
			if err != nil {
				panic(err)
			}
			ls(dir)
		}

	}
}

func chdir(path string) {
	os.Chdir(path)
}

func echo(text string) {
	fmt.Println(text)
}

func yes() {
	for {
		fmt.Println("yes")
	}
}

func ls(path string) {
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, file := range files {
		fmt.Print("  " + file.Name())
	}
}

func pwd() {
	fmt.Println(os.Getwd())
}
