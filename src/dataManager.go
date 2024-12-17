package TarminalRPG

import (
	"encoding/json"
	"os"
)

func SaveData(project *Project, path string) error {
	data, err := json.Marshal(project)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, data, 0)
	if err != nil {
		return err
	}

	return nil
}

func LoadData(path string) (*Project, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var project Project
	Parser(string(data), &project)
	Lexer(&project)

	return &project, nil
}
