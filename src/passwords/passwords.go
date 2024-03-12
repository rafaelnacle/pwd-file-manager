package passwords

import (
	"encoding/json"
	"os"
)

type Password struct {
	Website  string `json:"website"`
	Username string `json:"username"`
	Password string `json:"password"`
}

const filename = "passwords.json"

func savePasswords(passwords []Password, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")

	if err := encoder.Encode(passwords); err != nil {
		return err
	}

	return nil
}

func loadPassword(filename string) ([]Password, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var passwords []Password
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&passwords); err != nil {
		return nil, err
	}

	return passwords, nil
}
