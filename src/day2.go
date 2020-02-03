package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
	"unicode"
	"unicode/utf8"
)

var (
	inputFile = flag.String("input", "input/day2.txt", "File with input of day2")
)

func ScanNumber(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// skip leading non-digit
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if unicode.IsDigit(r) {
			break
		}
	}
	// scan whole number
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if !unicode.IsDigit(r) {
			return i + width, data[start:i], nil
		}
	}

	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	return start, nil, nil
}

func readInput(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var intcode []int
	scanner := bufio.NewScanner(file)
	scanner.Split(ScanNumber)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		intcode = append(intcode, n)
	}
	return intcode
}

func execIntcode(intcode []int) bool {
	offset := 0
	for offset < len(intcode) {
		switch opcode := intcode[offset]; opcode {
		case 1:
			pos1 := intcode[offset + 1]
			pos2 := intcode[offset + 2]
			pos3 := intcode[offset + 3]
			intcode[pos3] = intcode[pos1] + intcode[pos2]
			offset += 4
		case 2:
			pos1 := intcode[offset + 1]
			pos2 := intcode[offset + 2]
			pos3 := intcode[offset + 3]
			intcode[pos3] = intcode[pos1] * intcode[pos2]
			offset += 4
		case 99:
			return true
		}
	}
	return false
}

func findNumbers(intcode []int, output int) (int, int) {
	// copy original intcode
	original_intcode := make([]int, len(intcode))
	copy(original_intcode, intcode)

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			copy(intcode, original_intcode)
			intcode[1] = noun
			intcode[2] = verb

			execIntcode(intcode)

			if intcode[0] == output {
				return noun, verb
			}
		}
	}
	return -1, -1
}

func main() {
	log.SetFlags(0)
	flag.Parse()

	intcode := readInput(*inputFile)

	noun, verb := findNumbers(intcode, 19690720)

	log.Println(noun * 100 + verb)
}
