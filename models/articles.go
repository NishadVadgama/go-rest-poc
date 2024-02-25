package models

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var Articles []Article

// Load dummy articles from json
func init() {
    // Open the JSON file
    file, err := os.Open("data/sample.json")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close() // it will close the file at the end of the function

    // Read the file content
    bytes, err := io.ReadAll(file)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Unmarshal the JSON data into the Person struct
    err = json.Unmarshal(bytes, &Articles)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
}