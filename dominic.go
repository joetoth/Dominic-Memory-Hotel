// Go implementation of the Dominic Hotel mnemonic memory system 
// found in the book "Mind Performance Hacks". This is a system to help you
// remember up to 10,000 things.
// http://www.lifetrainingonline.com/blog/how-to-remember-100-things.htm
//
// In order to remember long numbers a mnemonic system is needed. The following
// mnemonic associating is the building base.
// 0=O, 1=A, 2=B, 3=C, 4=D, 5=E, 6=S, 7=G, 8=H, 9=N
// 
// Next you need to create a memoryMap of names and actions to represent a 2
// digit number. For instance:
// AE => Albert Einstein: Writing on the chalkboard
// would represent number 15
// HO => Santa Claus: Laughing "HO HO HO"
// would represent number 80

// If you want to represent the number 1580 you would take the person from
// the first number and action from the second number, yielding: Albert Einstein
// laughing "HO HO HO"

// Yes, this means you need to come up with 100 name/action combinations.

// This program helps you see your 10,000 combinations
// The dominic.txt file is of the format Mnemonic:Name:Action, where a colon is the 
// separator and shouldn't be in the name or action.

package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var rangeFlag *string = flag.String("r", "", "Print the contents of cells. For example '0-9'. Cell numbers start at 0.")

type NameAction struct {
	name   string
	action string
}

// Read a whole file into the memory and store it as array of lines
func parse(path string) (memoryMap map[int]NameAction, err error) {
	var (
		file       *os.File
		part       []byte
		prefix     bool
		lineNumber int
	)
	memoryMap = make(map[int]NameAction)

	if file, err = os.Open(path); err != nil {
		return
	}
	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 1024))
	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(part)
		if !prefix {

			line := strings.Split(buffer.String(), ":")
			memoryMap[lineNumber] = NameAction{line[1], line[2]}
			buffer.Reset()
		}

		lineNumber++
	}
	//	fmt.Errorf("zzzzzerr%s", err)
	//	if err == os.ErrInvalid {
	//		err = nil
	//	}
	err = nil
	return
}

/**
 * 0=O, 1=A, 2=B, 3=C, 4=D, 5=E, 6=S, 7=G, 8=H, 9=N
 */
func mnemonic(num int) (first string, second string) {

	toChar := func(i int) string {

		var char string

		switch i {
		case 0:
			char = "O"
		case 1:
			char = "A"
		case 2:
			char = "B"
		case 3:
			char = "C"
		case 4:
			char = "D"
		case 5:
			char = "E"
		case 6:
			char = "S"
		case 7:
			char = "G"
		case 8:
			char = "H"
		case 9:
			char = "N"
		}

		return char
	}

	var firstNum, secondNum int

	if num >= 10 {
		firstNum = num / 10
	}

	secondNum = num % 10

	first = toChar(firstNum)

	second = toChar(secondNum)

	return
}

func MergedCell(memoryMap map[string]NameAction, name, action string) (ret NameAction) {
	return NameAction{memoryMap[name].name, memoryMap[action].action}
}

func main() {

	flag.Parse()

	// filename, start, end
	memoryMap, err := parse("src/dominic.txt")
	if err != nil {
		fmt.Println("Error: %s\n", err)
		return
	}

	if *rangeFlag != "" {
		var cellsRange string = *rangeFlag
		var splitRange []string = strings.Split(cellsRange, "-")
		begin, err := strconv.Atoi(splitRange[0])
		end, err := strconv.Atoi(splitRange[1])

		if err != nil {
			fmt.Println("error", err)
			return
		}

		for i := begin; i <= end; i++ {
			// Determine first cell
			firstCellNumber := i / 100
			secondCellNumber := i
			if firstCellNumber > 0 {
				secondCellNumber = i % (firstCellNumber * 100)
			}

			var firstCell, secondCell NameAction = memoryMap[firstCellNumber], memoryMap[secondCellNumber]
			fmt.Println(i, firstCell.name, "=>", secondCell.action)
		}
	}

}
