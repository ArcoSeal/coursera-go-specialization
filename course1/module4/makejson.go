package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	entry := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Enter name: ")
	scanner.Scan()
	entry["name"] = scanner.Text()

	fmt.Printf("Enter address: ")
	scanner.Scan()
	entry["address"] = scanner.Text()

	entryJSON, _ := json.Marshal(entry)
	fmt.Println(string(entryJSON))
}
