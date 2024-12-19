package TerminalRPG

import (
	"encoding/json"
	"fmt"
)

func Parser(data string, parsed *Project) {
	if err := json.Unmarshal([]byte(data), &parsed); err != nil {
		fmt.Println("Error:", err)
		return
	}
}
