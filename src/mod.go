package TerminalRPG

import (
	"os"
)

func Test() {
	data, _ := os.ReadFile("sample.json")

	var MyProject Project
	Parser(string(data), &MyProject)
	Lexer(&MyProject)
	Runner(&MyProject)
}

var FuncList map[string]FuncArgs = map[string]FuncArgs{}

type Project struct {
	Title  string   `json:"title"`
	Author string   `json:"author"`
	Body   []Module `json:"body"`
}

type Module struct {
	ModuleName string      `json:"moduleName"`
	Args       interface{} `json:"args"`
}

type SayArgs struct {
	Text []string `json:"text"`
}

type IfArgs struct {
	Yes []Module `json:"yes"`
	No  []Module `json:"no"`
}

type FuncArgs struct {
	Name string   `json:"name"`
	Body []Module `json:"body"`
}

type CallArgs struct {
	Name string
}
