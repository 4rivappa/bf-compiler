package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

func handleError(e error) {
	if e != nil {
		panic(e)
	}
}

// TESTCASE: >>+[<+++++++++[->---------<]+>-[--[[-]<->]<[<[->-]>[<+]<]>]<[-<+>]+>,]<<[<+]

type lp struct {
	memIndex     int
	contextIndex int
}

var memory = [30000]uint8{}

var leftP []lp

func main() {
	// fmt.Println("Herewego")

	startTime := time.Now()

	arguments := os.Args
	if len(arguments) <= 1 {
		log.Fatal(errors.New("insufficient command line arguments"))
	}

	fileName := arguments[1]
	// fmt.Println(fileName)

	content, err := os.ReadFile(fileName)
	handleError(err)

	contentLength := len(content)
	memIndex := 0

	for index := 0; index < contentLength; index++ {
		// fmt.Println(index, content[index])

		switch content[index] {
		case 43: // operation (+)
			memory[memIndex]++
		case 45: // operation (-)
			memory[memIndex]--
		case 60: // operation (<)
			if memIndex-1 < 0 {
				// fmt.Println(memIndex, leftP)
				log.Fatal(errors.New("cannot move left beyond 0"))
			}
			memIndex = memIndex - 1
		case 62: // operation (>)
			if memIndex+1 >= 30000 {
				// fmt.Println(memIndex)
				log.Fatal(errors.New("cannot move right beyond 29999"))
			}
			memIndex = memIndex + 1
		case 91: // operation ([)
			if memory[memIndex] == 0 {
				// fmt.Println(memory[memIndex])
				// for content[index] != 93 {
				// 	index++
				// }
				// TODO: Improve the code here .. this is breaking
				opensCount := 0
				index += 1
				// fmt.Println(memIndex, index, "[")
				for content[index] != 93 || opensCount != 0 {
					switch content[index] {
					case 93:
						opensCount -= 1
					case 91:
						opensCount += 1
					}
					// fmt.Println(index, opensCount)
					index += 1
				}
			} else {
				leftP = append(leftP, lp{memIndex, index})
			}
		case 93: // operation (])
			// if index > 50 {
			// 	fmt.Println(index)
			// 	fmt.Println(leftP)
			// }
			if len(leftP) < 1 {
				log.Fatal(errors.New("mismatch []'s"))
			}
			latest := leftP[len(leftP)-1]
			if memory[memIndex] == 0 {
				leftP = leftP[:len(leftP)-1]
			} else {
				index = latest.contextIndex
			}
		case 46: // operation (.)
			fmt.Printf("%c", memory[memIndex])
		case 44: // operation (,)
			fmt.Scan(&memory[memIndex])
		}
		// fmt.Println(memory[:5])
	}

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	seconds := elapsedTime.Seconds()
	fmt.Println(seconds)
}
