package main

import (
	"TarminalRPG/src/syosetu"
	"fmt"
)

func main() {
	/*/
	data, _ := os.ReadFile("Data/sample.json")

	var MyProject TarminalRPG.Project
	TarminalRPG.Parser(string(data), &MyProject)
	TarminalRPG.Lexer(&MyProject)
	TarminalRPG.Runner(&MyProject)
	/*/

	list, err := syosetu.GetSyosetuList("n6316bn")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(list)
}
