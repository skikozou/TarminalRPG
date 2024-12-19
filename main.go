package main

import (
	TerminalRPG "TerminalRPG/src"
	"TerminalRPG/src/syosetu"
	"fmt"
)

func main() {
	novel, err := syosetu.NoveltoProject("n6316bn")
	if err != nil {
		fmt.Println(err)
		return
	}

	TerminalRPG.Play(novel)

	//TerminalRPG.SaveData(novel, "Data/tensura.json")
}
