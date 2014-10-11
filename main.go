package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/codegangsta/cli"
)

func streamFormattedLines(reader io.Reader, token string) {
	scanner := bufio.NewScanner(os.Stdin)
	r := regexp.MustCompile(token)
	streak := [][]string{}
	streak = nil
	maxSize := 0
	for scanner.Scan() {
		splitLineOnToken := r.Split(scanner.Text(), 2)
		//spew.Dump(splitLineOnToken)
		if len(splitLineOnToken) > 1 {
			if streak == nil {
				streak = [][]string{}
			} else {
				streak = append(streak, splitLineOnToken)
			}
			if len(splitLineOnToken[0]) > maxSize {
				maxSize = len(splitLineOnToken[0])
			}
		} else {
			for _, s := range streak {
				metaFormat := fmt.Sprintf("%%-%ds%%s%%s\n", maxSize)
				fmt.Printf(metaFormat, s[0], token, s[1])
			}
			streak = nil
			fmt.Println(scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error scanning from ", reader)
	}

}

func main() {
	app := cli.NewApp()
	app.Name = "patign"
	app.Usage = "Align streams of data based on provided token"
	app.Action = func(c *cli.Context) {
		if len(c.Args()) != 1 {
			fmt.Fprintln(os.Stderr, "Wrong number of arguments, usage: patign [token]")
			return
		}
		token := c.Args()[0]
		reader := os.Stdin
		streamFormattedLines(reader, token)
	}

	app.Run(os.Args)
}
