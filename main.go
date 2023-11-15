package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
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

func CheckI(el bool) {
	changed_list_count := make(map[string]int)
	for key, _ := range lines_count {
		changed_list_count[strings.ToLower(key)]++
	}
	for key, _ := range changed_list_count {
		fmt.Println(key)
	}
}

func CheckU(el bool) {
	for key, count := range lines_count {
		if count == 1 {
			fmt.Println(key)
		}

	}
}

func CheckD(el bool) {
	for key, count := range lines_count {
		if count > 1 {
			fmt.Println(key)
		}

	}
}

func CheckC(el bool) {
	for key, count := range lines_count {
		fmt.Printf("%d %s\n", count, key)
	}
}

var lines_count = make(map[string]int)

func InitialConsoleInput() {
	if len(flag.Args()) == 0 {
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
			fmt.Println("Error: reading error occured")
		}

	}

}

func CheckForAdditionalInput() {
	if len(flag.Args()) > 2 {
		fmt.Println("Error: so many command-line arguments")
		return
	}

	if !strings.Contains(flag.Args()[0], ".txt") { // by te way input file.txt exist
		fmt.Printf("Error: incorrect command-line argument")
		return
	}
	fileInput, err := os.Open(flag.Args()[0])
	if err != nil {
		fmt.Println("Error: Occured problem during opening the file")
		panic(err)
	}
	data, err := io.ReadAll(fileInput)
	fmt.Println(string(data))
	defer fileInput.Close()

	if len(flag.Args()) == 2 { // check for needing in output.txt
		if !strings.Contains(flag.Args()[1], ".txt") {
			fmt.Printf("Error: incorrect command-line argument")
			return
		}
		fileOutput, err := os.OpenFile(flag.Args()[1], os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			fmt.Println("Error: Occured problem during opening the file")
			panic(err)
		}
		fileOutput.WriteString("WORKS")
		fileOutput.Close()
	}
}

func InitFlags() {
	// var args = []string{}

	flag.BoolVar(&flags.c, "c", false, "counting the frequency of each line in the text")
	flag.BoolVar(&flags.d, "d", false, "show only REPEATED lines")
	flag.BoolVar(&flags.u, "u", false, "show only UNIQUE lines")
	flag.IntVar(&flags.f, "f", 0, "to not read first N words")
	flag.IntVar(&flags.s, "s", 0, "to not read first N symbols")
	flag.BoolVar(&flags.i, "i", false, "to not focus on registr")
	flag.Parse()

	// flag.Visit(func(f *flag.Flag) {
	// 	fmt.Println("Flag:", f.Name, "Value:", f.Value)
	// })

}

func main() {
	InitFlags()
	if len(flag.Args()) > 0{
		CheckForAdditionalInput()
	}
	InitialConsoleInput()
	// An artificial input source.
	if BoolToInt(flags.c)+BoolToInt(flags.d)+BoolToInt(flags.u) > 1 {
		fmt.Println("Incorrect combination of flags (you can't choose more than 1 from --c, --d, --u)")
		return 

	}

	if flag.NFlag() == 0 {
		//fmt.Println("I'm in ZERO!")
		for key, _ := range lines_count {
			fmt.Println(key)
		}

	}
	if flags.c {
		CheckC(flags.c)
	}
	if flags.d {
		CheckD(flags.d)
	}
	if flags.u {
		CheckU(flags.u)
	}
	if flags.i {
		CheckI(flags.i)
	}

}
