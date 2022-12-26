package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"io"
	"path/filepath"
	"mime/multipart"
)

func main() {
	fmt.Println("file: ", os.Getenv("FILE"))
	fmt.Println("url: ", os.Getenv("URL"))
	fileflag := os.Getenv("FILE")
	url := os.Getenv("URL")

	file, err := os.Open(fileflag)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		fmt.Printf("Error creating form file: %v", err)
	}
	io.Copy(part, file)
	writer.Close()

	r, err := http.NewRequest("POST", url, body)
	if err != nil {
		fmt.Printf("Error creating request: %v", err)
	}
	
	r.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	client.Do(r)

	fmt.Println("Response: ", r.Response)
}
