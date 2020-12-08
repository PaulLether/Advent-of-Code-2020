package main
import (
    "bufio"
    "log"
	"os"
	"fmt"
)

// Part One: finding if two numbers in the slice equals the wanted number
func routeInfo	(vMove int, hMove int, lines[][] string) (int, int){
	vIndex, hIndex, numberOfTrees, openSquares := 0, 0, 0 , 0
	
	for vIndex + vMove < len(lines){
		vIndex += vMove
		hIndex = (hIndex + hMove)%len(lines[0]) // i have loaded it into an slice simliar to the input file rather than one wide slice

		if lines[vIndex][hIndex] == "#"{
			numberOfTrees++
		} else {
			openSquares++
		} 
	}
	return numberOfTrees, openSquares
}

// Part Two: Making movement multiplications
func calcaultingMovement(movement[][] int, lines[][] string) (int, int){
	totalNumberOfTrees, totalNumberOpenSquares := 1, 1
	for i := 0 ; i < len(movement) ; i++{
		numberOfTrees, openSquares := routeInfo(movement[i][0], movement[i][1], lines)
		totalNumberOfTrees *= numberOfTrees
		totalNumberOpenSquares *= openSquares
	}
	return totalNumberOfTrees, totalNumberOpenSquares 
}


func main() {
	// reading file
	file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
	scanner := bufio.NewScanner(file)

	// loading data into a 2d slice
	lines := make([][]string, 0)
    for scanner.Scan() {
		currentLine := scanner.Text()
		sliceOfChars := make([]string, 0)
		
		for charIndex := 0 ; charIndex < len(currentLine) ; charIndex++{
			sliceOfChars = append(sliceOfChars, string(currentLine[charIndex]))
		}
				lines = append(lines, sliceOfChars)
	}

	// Part One 
	numberOfTrees, openSquares := routeInfo(1, 3, lines)
	fmt.Printf("PartOne: \n number of tress= %d\n number of open squares = %d \n", numberOfTrees, openSquares)

	// Part Two
	movement := [][]int{
		{1, 1},
		{1, 3},
		{1, 5},
		{1, 7},
		{2, 1}}

	numberOfTrees, openSquares = calcaultingMovement(movement, lines)
	fmt.Printf("PartTwo: \n number of tress= %d\n number of open squares = %d \n", numberOfTrees, openSquares)
}