package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/codegangsta/cli"
)

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
				fmt.Println("MAXSIXE IS ", maxSize)
				for _, s := range streak {
					metaFormat := fmt.Sprintf("%%-%ds", maxSize)
					fmt.Println(metaFormat)
					fmt.Printf(metaFormat, s[0])
					fmt.Print(token)
					fmt.Print("\n")
				}
				streak = nil
				fmt.Println(scanner.Text())
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error scanning from <STDIN>")
		}
	}

	app.Run(os.Args)
}
