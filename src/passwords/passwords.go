package passwords

import (
	"encoding/json"
	"fmt"
	"os"
)

type Password struct {
	Website  string `json:"website"`
	Username string `json:"username"`
	Password string `json:"password"`
}

const filename = "passwords.json"

func savePasswords(passwords []Password) error {
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

func loadPasswords() ([]Password, error) {
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

func AddPassword(website, username, password string) error {
	passwords, err := loadPasswords()
	if err != nil {
		return err
	}

	newPassword := Password{Website: website, Username: username, Password: password}
	passwords = append(passwords, newPassword)

	return savePasswords(passwords)
}

func EditPassword(website, username, password string) error {
	passwords, err := loadPasswords()
	if err != nil {
		return err
	}

	for i, pw := range passwords {
		if pw.Website == website {
			if username != "" {
				passwords[i].Username = username
			}
			if password != "" {
				passwords[i].Password = password
			}
			return savePasswords(passwords)
		}
	}

	return fmt.Errorf("password not found for website %s", website)
}

func DeletePassword(website string) error {
	passwords, err := loadPasswords()
	if err != nil {
		return err
	}

	for i, pw := range passwords {
		if pw.Website == website {
			passwords = append(passwords[:i], passwords[i+1:]...)
			return savePasswords(passwords)
		}
	}

	return fmt.Errorf("password not found for website %s", website)
}

func ListPasswords() ([]Password, error) {
	return loadPasswords()
}
