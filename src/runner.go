package TarminalRPG

import (
	"fmt"

	"github.com/mattn/go-tty"
)

func Runner(project *Project) {
	fmt.Println(project.Title, "\n", project.Author, "\n")

	ModuleRunner(project.Body)
}

func ModuleRunner(project []Module) {
	for _, module := range project {
		switch module.ModuleName {
		case "say":
			var args SayArgs
			JsonToObject(module.Args, &args)
			SayFunc(args)
		case "if":
			var args IfArgs
			JsonToObject(module.Args, &args)
			IfFunc(args)
		case "call":
			var args CallArgs
			JsonToObject(module.Args, &args)
			CallFunc(args)
		case "func":
			var args FuncArgs
			JsonToObject(module.Args, &args)
			FuncFunc(args)
		default:
			//none
		}
	}
}

func SayFunc(args SayArgs) {
	tty, err := tty.Open()
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	for _, text := range args.Text {
		fmt.Print(text, "\n\n")

	waitRoop:
		for {
			r, err := tty.ReadRune()
			if err != nil {
				fmt.Println("Error reading input:", err)
				return
			}

			switch r {
			case '\r':
				break waitRoop
			case ' ':
				break waitRoop
			}
		}
	}

	tty.Close()
}

func IfFunc(args IfArgs) {
	options := []string{"Yes", "No"}
	selectedIndex := 0
	fmt.Print("\033[?25l")
	tty, err := tty.Open()
	if err != nil {
		fmt.Println("Error initializing tty:", err)
		return
	}

	renderMenu(options, selectedIndex)

readRoop:
	for {
		r, err := tty.ReadRune()
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		fmt.Printf("\033[%dA", len(options))
		switch r {
		case 'A':
			selectedIndex = 1 - selectedIndex
		case 'B':
			selectedIndex = 1 - selectedIndex
		case '\r':
			fmt.Printf("\033[%dB", len(options))
			fmt.Print("\033[?25h\n")
			tty.Close()
			break readRoop
		}
		renderMenu(options, selectedIndex)
	}

	if selectedIndex == 0 {
		ModuleRunner(args.Yes)
	} else {
		ModuleRunner(args.No)
	}
}

func CallFunc(args CallArgs) {
	if val, ok := FuncList[args.Name]; ok {
		ModuleRunner(val.Body)
	} else {
		//not found
	}
}

func FuncFunc(args FuncArgs) {
	FuncList[args.Name] = args
}

func renderMenu(options []string, selectedIndex int) {
	for i, option := range options {
		if i == selectedIndex {
			fmt.Printf("> %s\n", option)
		} else {
			fmt.Printf("  %s\n", option)
		}
	}
}
