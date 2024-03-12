# Password File Manager

The Password File Manager is a lightweight utility for managing password data conveniently in a JSON file format. It offers simple commands for adding, editing, deleting, and viewing password entries, making it ideal for individual users and small teams. Securely store your sensitive password data and access it whenever needed with ease.

## Installation
To install the Password File Manager, simply clone the repository and build the binary using Go:
```bash
 git clone https://github.com/your-username/password-manager.git
 cd password-manager
 go mod tidy
```

## Usage

**Build the project:**
```bash
 go build -o pwd-manager src/main.go
```

**Add Password Entry**: Add a new password entry with website, username, and password.
```bash
pwd-manager add -w website -u username -p password
```

**Edit Password Entry:** Update an existing password entry with new website, username, or password.
```bash
pwd-manager edit -w website -u new_username -p new_password
```

**Delete Password Entry:** Remove a password entry from the database.
```bash
pwd-manager delete -w website
```

**List Passwords:** View all stored password entries.
```bash
pwd-manager list
```

**Help:** Display usage information and available commands.
```bash
pwd-manager help
```

