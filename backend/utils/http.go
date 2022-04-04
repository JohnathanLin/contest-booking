package utils

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

func Get(url string, headerMap map[string]string) string {
	client := &http.Client{Timeout: 5 * time.Second}
	request, err := http.NewRequest("GET", url, nil)
	for header := range headerMap {
		request.Header.Add(header, headerMap[header])
	}
	if err != nil {
		panic(err)
	}
	resp, _ := client.Do(request)

	defer resp.Body.Close()
	var buffer [1024]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	return result.String()
}
