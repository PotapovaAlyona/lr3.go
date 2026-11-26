package main

import (
	"encoding/gob"
	"encoding/json"
	"os"
)

// SaveJSON сохраняет любую структуру данных в текстовом формате JSON
func SaveJSON(data interface{}, filename string) error {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, bytes, 0644)
}

// LoadJSON загружает любую структуру данных из текстового формата JSON
func LoadJSON(target interface{}, filename string) error {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, target)
}

// SaveBinary сохраняет любую структуру данных в бинарном формате (encoding/gob)
func SaveBinary(data interface{}, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	return encoder.Encode(data)
}

// LoadBinary загружает любую структуру данных из бинарного формата (encoding/gob)
func LoadBinary(target interface{}, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	return decoder.Decode(target)
}
