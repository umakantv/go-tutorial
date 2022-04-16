package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// Custom Writer for writting logs to the terminal in a trimmed fashion
type TrimLogWriter struct{}

func (TrimLogWriter) Write(p []byte) (int, error) {
	n, err := fmt.Println(time.Now().Format(time.UnixDate), string(p[:120]))
	return n, err
}

func main() {

	Bot()
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	// bs := make([]byte, 120)
	// readBytes, _ := resp.Body.Read(bs)
	// fmt.Println(readBytes, string(bs))

	// io.Copy(os.Stdout, resp.Body)
	io.Copy(TrimLogWriter{}, resp.Body)

}
