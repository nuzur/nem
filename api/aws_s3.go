package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/nuzur/nem/aws"
)

func serverHandlerUploadFunc() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("origin")
		origin = strings.ReplaceAll(origin, "api.", "admin.")
		method := r.Method
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Accept", "application/json, multipart/mixed")
		w.Header().Set("Content-Type", "application/json")
		if method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		awsClient, err := aws.New(nil)
		if err != nil {
			fmt.Printf("error creating aws client\n")
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}

		file, header, err := r.FormFile("file")
		if err != nil {
			fmt.Printf("error reading file\n")
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, file); err != nil {
			fmt.Printf("error reading file to bytes\n")
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}

		filename := "nem/" + header.Filename
		res, err := awsClient.UploadToS3(filename, buf.Bytes())
		if err != nil {
			fmt.Printf("error uploading file: %v, %v\n", err, res)
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}

		filepath := "https://" + awsClient.Config().Bucket + "." + "s3-" + awsClient.Config().Region + ".amazonaws.com/" + filename
		resJson, _ := json.Marshal(struct {
			FilePath string `json:"file_path"`
		}{
			FilePath: filepath,
		})
		w.Write(resJson)

	})
}
