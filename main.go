package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type Flags struct {
	c bool
	d bool
	u bool
	f int
	s int
	i bool
}

var flags Flags

func BoolToInt(el bool) int {
	if el {
		return 1
	} else {
		return 0
	}
}

func init() {
	flag.BoolVar(&flags.c, "c", false, "counting the frequency of each line in the text")
	flag.BoolVar(&flags.d, "d", false, "show only REPEATED lines")
	flag.BoolVar(&flags.u, "u", false, "show only UNIQUE lines")
	flag.IntVar(&flags.f, "f", 0, "to not read first N words")
	flag.IntVar(&flags.s, "s", 0, "to not read first N symbols")
	flag.BoolVar(&flags.i, "i", false, "to not focus on registr")
	flag.Parse()
}

func main() {
	// An artificial input source.

	if BoolToInt(flags.c)+BoolToInt(flags.d)+BoolToInt(flags.u) > 1 {
		fmt.Println("Incorrect combination of flags (you can't choose more than 1 from --c, --d, --u)")
		return
	}
	if len(flag.Args()) == 0 {
		scanner := bufio.NewScanner(os.Stdin)

		var lines []string
		lines_count := make(map[string]int)
		for {
			// read line from stdin using newline as separator
			scanner.Scan()
			line := scanner.Text()
			// if line is empty, break the loop
			if len(line) == 0 {
				break
			}

			//append the line to a slice
			lines = append(lines, line)
			lines_count[line]++
		}
		for key, _ := range lines_count {
			fmt.Println(key)
		}

		if scanner.Err() != nil {
			fmt.Println("Error occured")
		}

	}

}
