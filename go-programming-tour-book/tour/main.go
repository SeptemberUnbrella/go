package main

import (
	"flag"
	"fmt"
	"log"
)

var name string

func main() {
	flag.Parse()
	args := flag.Args()
	fmt.Printf("%s", args)
	if len(args) <= 0 {
		return
	}

	switch args[0] {
	case "go":
		goCmd := flag.NewFlagSet("go", flag.ExitOnError)
		goCmd.StringVar(&name, "name", "Go语言", "帮助信息")
		_ = goCmd.Parse(args[1:])

	case "php":
		phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
		phpCmd.StringVar(&name, "n", "PHP语言", "帮助信息")
		_ = phpCmd.Parse(args[1:])
	}
	log.Printf("name: %s", name)

}
