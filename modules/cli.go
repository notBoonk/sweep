package modules

import (
	"fmt"
	"strings"
)

var targets []string

func ScanCommand() {
	if len(targets) == 0 {
		fmt.Println("\nNo targets provided\n")
		return
	}
	Scan(targets)
}

func AddCommand(t string) {
	input := t
	input = strings.Replace(input, " ", "", -1)
	input = strings.Replace(input, "add", "", -1)
	hosts := strings.Split(input, ",")

	if hosts[0] != "" {
		for x := range hosts {
			targets = append(targets, hosts[x])
		}
	}
}

func RemoveCommand(t string) {
	input := t
	input = strings.Replace(input, " ", "", -1)
	input = strings.Replace(input, "remove", "", -1)
	hosts := strings.Split(input, ",")

	var newTargets []string

	if hosts[0] == "all" {
		targets = newTargets
		return
	}

	if hosts[0] == "" {
		return
	}

	for x := range targets {
		contains := false
		for y := range hosts {
			if targets[x] == hosts[y] {
				contains = true
				break
			}
		}
		if contains == false {
			newTargets = append(newTargets, targets[x])
		}
	}
	targets = newTargets
}

func TargetsCommand() {
	if len(targets) > 0 {
		fmt.Println("")
		for x := range targets {
			fmt.Println(targets[x])
		}
		fmt.Println("")
	} else {
		fmt.Println("\nNo targets provided\n")
	}
}
