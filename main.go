package main

import (
	"bytes"
	"fmt"
	"flag"
	"net/http"
	"os"
	"io"
	"path/filepath"
	"mime/multipart"
)

func main() {
	// get --file and --url flags
	fileflag := flag.String("file", "", "file to upload")
	url := flag.String("url", "", "url to upload to")
	flag.Parse()

	file, err := os.Open(*fileflag)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	r, _ := http.NewRequest("POST", *url, body)
	r.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	client.Do(r)

	fmt.Println("Response: ", r.Response)
}
