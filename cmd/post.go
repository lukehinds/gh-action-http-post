/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"io"
	"path/filepath"
	"mime/multipart"

	"github.com/spf13/cobra"
)

// postCmd represents the post command
var postCmd = &cobra.Command{
	Use:   "post",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fileflag := cmd.Flag("file").Value.String()
		url := cmd.Flag("url").Value.String()
		// open file and upload it to the server

		file, err := os.Open(fileflag)
		if err != nil {
			fmt.Printf("Error opening file: %v", err)
		}
		defer file.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
		io.Copy(part, file)
		writer.Close()

		r, _ := http.NewRequest("POST", url, body)
		r.Header.Add("Content-Type", writer.FormDataContentType())
		client := &http.Client{}
		client.Do(r)

		fmt.Println("Response: ", r.Response)
	},
}

func init() {
	rootCmd.AddCommand(postCmd)
	postCmd.Flags().StringP("file", "f", "", "file to post")
	postCmd.Flags().StringP("url", "u", "", "url to post")
}
