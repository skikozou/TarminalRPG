package main

import (
	TarminalRPG "TarminalRPG/src"
	"TarminalRPG/src/syosetu"
	"fmt"
)

func main() {
	novel, err := syosetu.NoveltoProject("n6316bn")
	if err != nil {
		fmt.Println(err)
		return
	}

	TarminalRPG.Runner(novel)

	//TarminalRPG.SaveData(novel, "Data/tensura.json")
}
