package main

import (
	TarminalRPG "TarminalRPG/src"
	"os"
)

func main() {
	data, _ := os.ReadFile("Data/sample.json")

	var MyProject TarminalRPG.Project
	TarminalRPG.Parser(string(data), &MyProject)
	TarminalRPG.Lexer(&MyProject)
	TarminalRPG.Runner(&MyProject)
}
