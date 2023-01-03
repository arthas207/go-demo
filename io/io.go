package io

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func Example() {
	var jsonBody = []byte(`{"type":"page","title":"WEMIX","space":{"key":"BitDeskCore"},"ancestors":[{"id":360591}],"body":{"storage":{"value":"test","representation":"storage"}}}`)
	var bodyReader = bytes.NewReader(jsonBody)
	fmt.Println("body:", string(jsonBody))
	var req, err = http.NewRequest(http.MethodPost, "http://47.92.98.6:8090/confluence/rest/api/content/", bodyReader)
	if err == nil {
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		req.Header.Set("Authorization", "Basic emhhbmd4aWFvaGVuZzpaaGFuZ3hpYW9oZW5nQEJpdGRlc2s=")
		var client = http.Client{
			Timeout: 30 * time.Second,
		}
		var resp, error = client.Do(req)
		if error != nil {
			fmt.Println(error)
		}
		defer resp.Body.Close()
		file, err := os.Create("/Users/zxh/test")
		if err != nil {
			log.Fatalf("failed creating file: %s", err)
		}
		defer file.Close()

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		file.WriteString(string(b))
	}
}
