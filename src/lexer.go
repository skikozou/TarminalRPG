package TerminalRPG

import (
	"encoding/json"
	"fmt"
)

func Lexer(parsed *Project) {
	var lexerd Project
	lexerd.Author = parsed.Author
	lexerd.Title = parsed.Title
	lexerd.Body = make([]Module, len(parsed.Body))

	lexerd.Body = ModuleLexer(lexerd.Body, parsed.Body)

	parsed = &lexerd
}

func ModuleLexer(base []Module, Modules []Module) []Module {
	for i, module := range Modules {
		switch module.ModuleName {
		case "say":
			var args SayArgs
			JsonToObject(module.Args, &args)
			if HasBrank(&args) {
				//set break flag
				//HasBrank setting
			}
			base[i].Args = args
			base[i].ModuleName = module.ModuleName
		case "if":
			var args IfArgs
			JsonToObject(module.Args, &args)
			var YesModule []Module = make([]Module, len(args.Yes))
			var NoModule []Module = make([]Module, len(args.No))
			args.Yes = ModuleLexer(YesModule, args.Yes)
			args.No = ModuleLexer(NoModule, args.No)
			base[i].Args = args
			base[i].ModuleName = module.ModuleName
		case "func":
			var args FuncArgs
			JsonToObject(module.Args, &args)
			var FuncBody []Module = make([]Module, len(args.Body))
			args.Body = ModuleLexer(FuncBody, args.Body)
			base[i].Args = args
			base[i].ModuleName = module.ModuleName
		case "call":
			var args CallArgs
			JsonToObject(module.Args, &args)
			base[i].Args = args
			base[i].ModuleName = module.ModuleName
		default:
			fmt.Println("Unknown module:", module.ModuleName)
		}
	}

	return base
}

func JsonToObject(Json interface{}, ObjectType any) {
	jsonBytes, _ := json.Marshal(Json)
	json.Unmarshal(jsonBytes, &ObjectType)
}

func HasBrank(Json interface{}) bool {
	switch Json.(type) {
	case *SayArgs:
		for _, t := range Json.(SayArgs).Text {
			if t == "" {
				return true
			}
		}
	case *IfArgs:

	case *FuncArgs:

	case *CallArgs:

	}
}
