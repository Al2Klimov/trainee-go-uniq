package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	uniqueParameter := flag.Bool("u", false, "allows you to print only unique lines.")
	flag.Parse()

	buf := bufio.NewReader(os.Stdin)
	checkEquality := []byte("")
	skip := false

	for {
		content, readErr := buf.ReadBytes('\n')
		if readErr != nil && readErr != io.EOF {
			fmt.Fprintln(os.Stderr, readErr)
			os.Exit(1)
		}

		if len(content) != 0 && !bytes.HasSuffix(content, []byte{'\n'}) {
			content = append(content, '\n')
		}

		if *uniqueParameter {
			if bytes.Compare(checkEquality, []byte("")) == 0 {
				skip = true
				checkEquality = content
				if readErr == io.EOF {
					break
				}
				continue
			} else if bytes.Compare(content, checkEquality) == 0 {
				checkEquality = content
				skip = false
				if readErr == io.EOF {
					break
				}
				continue
			} else if bytes.Compare(content, checkEquality) != 0 {
				if skip {
					skip = false
					_, _ = os.Stdout.Write(checkEquality)
				}
				checkEquality = content
				skip = true
				if readErr == io.EOF {
					break
				}
				continue
			}
		} else {
			if bytes.Compare(checkEquality, content) != 0 {
				_, _ = os.Stdout.Write(content)
			}
			if readErr == io.EOF {
				break
			}
		}

		checkEquality = content
	}
}
