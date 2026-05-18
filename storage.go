package main

import (
	"encoding/json"
	"os"
)

type storage[T any] struct {
	FileName string
}

func NewStorage[T any](fileName string) *storage[T] {
	return &storage[T]{FileName: fileName}
}

func (s *storage[T]) Save(data T) error {
	file, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err	
	}
	return os.WriteFile(s.FileName, file, 0644)
}

func (s *storage[T]) Load(data *T)  error {
	file, err := os.ReadFile(s.FileName)
	if err != nil {
		return  err
	}
	return json.Unmarshal(file, data)
	
}