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
	repeatParameter := flag.Bool("d", false, "only print duplicate lines, one for each group.")
	caseParameter := flag.Bool("i", false, "ignore differences in case when comparing.")
	countParameter := flag.Bool("c", false, "prefix lines by the number of occurrences.")
	flag.Parse()

	buf := bufio.NewReader(os.Stdin)
	checkEquality := []byte("")
	skip := false
	var count int
	if *repeatParameter {
		count = 0
	} else {
		count = 1
	}

	for {
		content, readErr := buf.ReadBytes('\n')
		if readErr != nil && readErr != io.EOF {
			fmt.Fprintln(os.Stderr, readErr)
			os.Exit(1)
		}

		if len(content) != 0 && !bytes.HasSuffix(content, []byte{'\n'}) {
			content = append(content, '\n')
		}

		if *caseParameter {
			content = bytes.ToUpper(content)
		}

		if *repeatParameter {
			if bytes.Compare(checkEquality, []byte("")) == 0 {
				checkEquality = content
				if readErr == io.EOF {
					break
				}
				count++
				continue
			} else if bytes.Compare(content,checkEquality) != 0 {
				if skip {
					skip = false
					if *countParameter {
						fmt.Printf("      %d ", count)
					}
					_, _ = os.Stdout.Write(checkEquality)
					count = 1
				}
				checkEquality = content
				if readErr == io.EOF {
					break
				}
				continue
			} else if bytes.Compare(content, checkEquality) == 0 {
				skip = true
				checkEquality = content
				count++
				if readErr == io.EOF {
					if *countParameter {
						fmt.Printf("      %d ", count)
					}
					_, _ = os.Stdout.Write(checkEquality)
					break
				}
				continue
			}
		} else if *uniqueParameter {
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
					if *countParameter {
						fmt.Printf("      %d ", count)
					}
					_, _ = os.Stdout.Write(checkEquality)
				}
				checkEquality = content
				skip = true
				if readErr == io.EOF {
					_, _ = os.Stdout.Write(content)
					break
				}
				continue
			}
		} else {
			if *countParameter {
				if bytes.Compare(checkEquality, []byte("")) == 0 {
					checkEquality = content
					if readErr == io.EOF {
						break
					}
					continue
				}
				if bytes.Compare(checkEquality, content) != 0 {
					fmt.Printf("      %d ", count)
					_, _ = os.Stdout.Write(checkEquality)
					count = 0
				}
				count++
			} else {
				if bytes.Compare(checkEquality, content) != 0 {
					_, _ = os.Stdout.Write(content)
				}
				if readErr == io.EOF {
					break
				}
			}
		}

		checkEquality = content
	}
}
