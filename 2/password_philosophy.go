package main
import (
    "bufio"
    "log"
	"os"
	"regexp"
	"strings"
	"strconv"
	"fmt"
)

func main() {
	// reading file
	file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
	scanner := bufio.NewScanner(file)
	r:= regexp.MustCompile(`^(\d+)-(\d+) ([a-z]): ([a-z]+)$`)

	partOneValid := 0
	partTwoValid := 0
    for scanner.Scan() {
		// part one 
		matches := r.FindStringSubmatch(scanner.Text())
		min, _ := strconv.Atoi(matches[1])
		max, _ := strconv.Atoi(matches[2])
		letter := matches[3]
		password := matches[4]

		if  min <= strings.Count(password, letter) && strings.Count(password, letter) <= max{
			partOneValid++
		}

		// part two
		if (max <= len(password) && string(password[min -1]) == letter && string(password[max -1]) != letter) || 
		(max <= len(password) && string(password[min -1]) != letter && string(password[max -1]) == letter){
			partTwoValid++
		}
	}
	fmt.Printf("PartOne valid= %d\n", partOneValid)
	fmt.Printf("PartTwo valid= %d\n", partTwoValid)
}