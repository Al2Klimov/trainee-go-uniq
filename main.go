package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var checkEquality []byte

	for {
		content, readErr := buf.ReadBytes('\n')
		if readErr != nil && readErr != io.EOF {
			fmt.Fprintln(os.Stderr, readErr)
			os.Exit(1)
		}

		if len(content) != 0 && !bytes.HasSuffix(content, []byte{'\n'}) {
			content = append(content, '\n')
		}

		if bytes.Compare(checkEquality, content) != 0 {
			_, _ = os.Stdout.Write(content)
		}

		if readErr == io.EOF {
			break
		}

		checkEquality = content
	}
}
