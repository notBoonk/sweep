package main

import (
	"github.com/c-bata/go-prompt"
	"os"
	"strings"
	"sweep/modules"
)

var targets []string

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{}
	args := strings.Split(d.Text, " ")
	if len(args) < 2 {
		s = []prompt.Suggest{
			{Text: "scan", Description: "Begin scan on target list"},
			{Text: "add", Description: "Add subnet(s) to target list"},
			{Text: "remove", Description: "Remove subnet(s) to target list"},
			{Text: "targets", Description: "Show target list"},
			{Text: "clear", Description: "Clear terminal"},
			{Text: "exit", Description: "Exit program"},
		}
	} else if args[0] == "remove" {
		s = append(s, prompt.Suggest{Text: "all"})
		for x := range targets {
			s = append(s, prompt.Suggest{Text: targets[x]})
		}
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func main() {
	modules.ClearTerminal()
	for {
		t := prompt.Input("sweep:~$ ", completer)
		command := strings.Split(t, " ")
		switch command[0] {
		case "add":
			modules.AddCommand(t)

		case "remove":
			modules.RemoveCommand(t)

		case "targets":
			modules.TargetsCommand()

		case "scan":
			modules.ScanCommand()

		case "clear":
			modules.ClearTerminal()

		case "exit":
			modules.ClearTerminal()
			os.Exit(0)
		}
	}
}
