package main

// -f -s (delete)-> -i (registr) -> --c --d --u

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type Flags struct {
	f int
	s int
	i bool
	c bool
	d bool
	u bool
}

var flags Flags

func BoolToInt(el bool) int {
	if el {
		return 1
	} else {
		return 0
	}
}

func CheckI() {
	fmt.Println("\n________after I___________")
	changed_list_count := make(map[string]int)
	fmt.Printf("i got at start %v\n", lines_count)
	for key, val := range lines_count {
		fmt.Printf("for |%v| i did |%v|\n", key, strings.ToLower(key))
		changed_list_count[strings.ToLower(key)] += val
	}
	fmt.Println(changed_list_count)
	lines_count = changed_list_count
	//fmt.Println(lines_count)
	fmt.Println("-----------I'M IN I----------")
}

func CheckU() {

	fmt.Println("\n________after U___________")
	for key, count := range lines_count {
		if count > 1 {
			delete(lines_count, key)
		}
	}
	fmt.Println(lines_count)
	fmt.Println("------------I'M IN U-----------")
}

func CheckD() { // --d flag
	fmt.Println("\n________after D___________")
	for key, count := range lines_count {
		if count == 1 {
			delete(lines_count, key)
		}
	}
	fmt.Println(lines_count)
	fmt.Println("----------I'M IN D---------")
}

func CheckF() { // --f= flag
	fmt.Println("\n________after F___________")
	changed_list_count := make(map[string]int)
	for key, val := range lines_count {
		if len(strings.Fields(key)) > flags.f {
			changed_list_count[strings.Join(strings.Fields(key)[flags.f:], " ")] += val
			fmt.Printf("for |%v| i did |%v|\n", key, strings.Join(strings.Fields(key)[flags.f:], " "))
		} else {
			changed_list_count[" "] += val
			fmt.Printf("for |%v| i did |' '|\n", key)
		}
	}
	lines_count = changed_list_count
	fmt.Println(lines_count)
	fmt.Println("----------I'M IN F-------------")
}

func CheckS() { // --s= flag
	fmt.Println("\n________after S___________")
	changed_list_count := make(map[string]int)
	for key, val := range lines_count {
		if len(key) > flags.s {
			fmt.Printf("for |%v| i did |%v|\n", key, key[flags.s:])
			changed_list_count[key[flags.s:]]+=val
		} else {
			changed_list_count[" "] += val
			fmt.Printf("for |%v| i did |' '|\n", key)
		}

	}
	lines_count = changed_list_count
	fmt.Println(lines_count)
	fmt.Println("---------I'M IN S-----------")

}

func CheckC() { // --c flag
	// it already works in default
	fmt.Println("\n________after C___________")
	fmt.Println("-----------I'M IN C------------")
}

var lines_count = make(map[string]int)
var txtCount = 0

func InitialConsoleInput() {
	if txtCount == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for {
			// read line from stdin using newline as separator
			scanner.Scan()
			line := scanner.Text()
			// if line is empty, break the loop
			lines_count[strings.TrimSpace(line)]++

			if len(line) == 0 {
				break
			}
			//append the line to a slice

		}
		fmt.Println(lines_count)

		if scanner.Err() != nil {
			fmt.Println("Error: reading error occured")
		}

	}

}

func ChechTxtCount() {
	for _, el := range flag.Args() {
		if strings.Contains(el, ".txt") {
			txtCount++
		}
	}
}

func CheckFlagCorrectness() {
	if BoolToInt(flags.c)+BoolToInt(flags.d)+BoolToInt(flags.u) > 1 {
		fmt.Println("Error: Incorrect combination of flags (you can't choose more than 1 from --c, --d, --u)")
		return

	}
}

func CheckForAdditionalInput() {
	if txtCount > 2 {
		fmt.Println("Error: so many .txt arguments")
		return
	}

	fileInput, err := os.Open(flag.Args()[0])
	if err != nil {
		fmt.Println("Error: occured")
	}
	defer fileInput.Close()

	scanner := bufio.NewScanner(fileInput)
	for scanner.Scan() {
		lines_count[strings.TrimSpace(scanner.Text())]++
	}
	fmt.Println(lines_count)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if txtCount == 2 { // check for needing in output.txt
		fmt.Println("OUTPUTED IN FILE")

		fileOutput, err := os.OpenFile(flag.Args()[1], os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			fmt.Println("Error: Occured problem during opening the file")
			return

		}
		fileOutput.WriteString("WORKS")
		fileOutput.Close()
	}
}
func CheckFlags() {
	flagFuncs := map[string]func(){
		"c": CheckC,
		"d": CheckD,
		"i": CheckI,
		"u": CheckU,
		"f": CheckF,
		"s": CheckS,
	}

	flag.Visit(func(f *flag.Flag) { //theoretically the flag can be false!
		fmt.Printf("%s  |  %s", f.Name, f.Value)
		flagFuncs[f.Name]()
	})

	if flag.NFlag() == 0 {
		fmt.Println("I'm in ZERO!")
		for key, _ := range lines_count {
			fmt.Println(key)
		}

	}
}

func InitFlags() {

	flag.BoolVar(&flags.c, "c", false, "counting the frequency of each line in the text")
	flag.BoolVar(&flags.d, "d", false, "show only REPEATED lines")
	flag.BoolVar(&flags.u, "u", false, "show only UNIQUE lines")
	flag.IntVar(&flags.f, "f", 0, "to not read first N words")
	flag.IntVar(&flags.s, "s", 0, "to not read first N symbols")
	flag.BoolVar(&flags.i, "i", false, "to not focus on registr")
	flag.Parse()

}

func showResult() {
	for _, val := range lines_count {
		fmt.Println(val)
	}
}
func main() {
	InitFlags()
	CheckFlagCorrectness()
	ChechTxtCount()
	if txtCount > 0 {
		CheckForAdditionalInput()
	}
	InitialConsoleInput()
	CheckFlags()
}
