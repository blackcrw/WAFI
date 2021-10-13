package printer

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

var (
	stdin    = *os.Stdin
	stdout   = *os.Stdout
	stderr   = *os.Stderr
)

func PrintString(t ...string) {
	var hour = time.Now().Format("15:04:05")
	var raw = strings.Join(t, " ")

	io.WriteString(&stdout, "["+hour+"] "+raw+"\n")
}

func ScanQ(t ...string) string {
	var raw = strings.Join(t, " ")

	PrintString(raw)

	scanner := bufio.NewReader(os.Stdin)
	response, err := scanner.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	response = strings.ToLower(response)

	if response == "\n" {
		return response
	}

	response = strings.ReplaceAll(response, "\n", "")

	return response
}
