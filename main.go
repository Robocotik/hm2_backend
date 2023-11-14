package main

import (
	"flag"
	"fmt"
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
	fmt.Println(flags)



}
