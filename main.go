package main

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
type count_num struct{
	count int
	num int
}
var lines_count = make(map[string]count_num)
var txtCount = 0

func BoolToInt(el bool) int {
	if el {
		return 1
	} else {
		return 0
	}
}

func CheckI() {
	fmt.Println("\n________after I___________")
	var changed_list_count = make(map[string]count_num)
	//fmt.Printf("i got at start %v\n", lines_count)
	for key, val := range lines_count {
		//fmt.Printf("for |%v| i did |%v|\n", key, strings.ToLower(key))
		arg := strings.ToLower(key)
		_, exist := changed_list_count[arg]
		if exist{
			tmp := changed_list_count[arg]
			tmp.count += val.count
			changed_list_count[arg] = tmp
		} else{
			changed_list_count[arg] = count_num{val.count,val.num}
		}
		fmt.Println(changed_list_count)

	}
	//fmt.Println(changed_list_count)
	lines_count = changed_list_count
	//fmt.Println(lines_count)
	fmt.Println("-----------I'M IN I----------")
}

func CheckU() {

	fmt.Println("\n________after U___________")
	for key, val := range lines_count {
		if val.count > 1 {
			delete(lines_count, key)
		}
	}
	//fmt.Println(lines_count)
	fmt.Println("------------I'M IN U-----------")
}

func CheckD() { // --d flag
	fmt.Println("\n________after D___________")
	for key, val := range lines_count {
		if val.count == 1 {
			delete(lines_count, key)
		}
	}
	fmt.Println(lines_count)
	fmt.Println("----------I'M IN D---------")
}

func CheckF() { // --f= flag
	fmt.Println("\n________after F___________")
	var changed_list_count = make(map[string]count_num)
	for key, val := range lines_count {
		arg := ""
		if len(strings.Fields(key)) > flags.f{
			arg = strings.Join(strings.Fields(key)[flags.f:], " ")
		} 
		_, exist := changed_list_count[arg]
		if exist {
			tmp := changed_list_count[arg]
			tmp.count += val.count
			changed_list_count[arg] = tmp
		}else{
			changed_list_count[arg] = count_num{val.count,val.num}
		}
	}
	lines_count = changed_list_count
	fmt.Println(lines_count)
	fmt.Println("----------I'M IN F-------------")
}

func CheckS() { // --s= flag
	fmt.Println("\n________after S___________")
	var changed_list_count = make(map[string]count_num)
	for key, val := range lines_count {
		arg := ""
		if len(key) > flags.f{
			arg = key[flags.f:]
		} 
		_, exist := changed_list_count[arg]
		if exist {
			tmp := changed_list_count[arg]
			tmp.count += val.count
			changed_list_count[arg] = tmp
		}else{
			changed_list_count[arg] = count_num{val.count,val.num}
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

func InitialConsoleInput() {
	if txtCount == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		i := 0
		for {
			// read line from stdin using newline as separator
			scanner.Scan()
			
			// if line is empty, break the loop
			line := strings.TrimSpace(scanner.Text())
			_, exist := lines_count[line]
			if exist{
				tmp := lines_count[line]
				tmp.count ++
				lines_count[line] = tmp
			} else{
				lines_count[line] = count_num{1,i}
			}
			i++

			if len(line) == 0 {
				break
			}
			//append the line to a slice

		}
		//fmt.Println(lines_count)

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
	i:= 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		_, exist := lines_count[line]
		if exist{
			tmp := lines_count[line]
			tmp.count ++
			lines_count[line] = tmp
		} else{
			lines_count[line] = count_num{1,i}
		}
		i++
	}
	
	//fmt.Println(lines_count)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
func CheckFlags() {
	flagFuncs := map[string]func(){
		"f": CheckF,
		"s": CheckS,
		"i": CheckI,
		"c": CheckC,
		"d": CheckD,
		"u": CheckU,
	}

	flag.Visit(func(f *flag.Flag) { //theoretically the flag can be false!
		//fmt.Printf("%s  |  %s", f.Name, f.Value)
		flagFuncs[f.Name]()
	})
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
	fmt.Println(lines_count)
	if txtCount == 2 { // check for needing in output.txt
		//fmt.Println("OUTPUTED IN FILE")

		fileOutput, err := os.OpenFile(flag.Args()[1], os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			fmt.Println("Error: Occured problem during opening the file")
			return

		}
		//fmt.Println(lines_count)
		for key, val := range lines_count {
			if flags.c {
				fileOutput.WriteString(fmt.Sprint(val))
				fileOutput.WriteString(" " + key + "\n")
			} else {
				fileOutput.WriteString(key + "\n")
			}
		}

		fileOutput.Close()
		return
	}

	for key, val := range lines_count {
		if flags.c {
			fmt.Printf("%v %v\n", val, key)
		} else {
			fmt.Println(key)
		}
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
	showResult()
}
