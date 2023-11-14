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
	i bool
	f int
	s int
	
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
	var lines_count = make(map[string]int)

	if BoolToInt(flags.c)+BoolToInt(flags.d)+BoolToInt(flags.u) > 1 {
		fmt.Println("Incorrect combination of flags (you can't choose more than 1 from --c, --d, --u)")
		return

	} else {
		scanner := bufio.NewScanner(os.Stdin)

		for {
			// read line from stdin using newline as separator
			scanner.Scan()
			line := scanner.Text()
			// if line is empty, break the loop
			if len(line) == 0 {
				break
			}

			//append the line to a slice
			lines_count[line]++
		}

		if scanner.Err() != nil {
			fmt.Println("Error occured")
		}
		
	}

	if flag.NFlag() == 0 {
		//fmt.Println("I'm in ZERO!")
		for key, _ := range lines_count {
			fmt.Println(key)
		}

	} else if flags.c {
		//fmt.Println("I'm in C!")
		for key, count := range lines_count {
			fmt.Printf("%d %s\n", count, key)
		}
	}

}
