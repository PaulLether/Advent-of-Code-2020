package main
import (
    "bufio"
    "fmt"
    "log"
	"os"
	"strconv"
)

// Part One: finding if two numbers in the slice equals the wanted number
func PartOne(expenseNumbers []int, targetNumber int) int {
	for firstNumber:= 0 ; firstNumber < len(expenseNumbers) ; firstNumber++{
		for secondNumber := 0 ; secondNumber < len(expenseNumbers) ; secondNumber++{
			if(expenseNumbers[firstNumber] + expenseNumbers[secondNumber] == targetNumber){
				return (expenseNumbers[firstNumber]*expenseNumbers[secondNumber])
			}
		}
	}
	return 0
}

// Part Two: finding if three numbers in the slice equals the wanted number
func PartTwo(expenseNumbers []int, targetNumber int) int {
	for firstNumber:= 0 ; firstNumber < len(expenseNumbers) ; firstNumber++{
		for secondNumber := 0 ; secondNumber < len(expenseNumbers) ; secondNumber++{
			for thridNumber := 0 ; thridNumber < len(expenseNumbers) ; thridNumber++{
				if(expenseNumbers[firstNumber] + expenseNumbers[secondNumber] + expenseNumbers[thridNumber] == targetNumber){
					return (expenseNumbers[firstNumber]*expenseNumbers[secondNumber]*expenseNumbers[thridNumber])
				}
			}
		}
	}
	return 0
}

func main() {
	// reading file
	file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
	scanner := bufio.NewScanner(file)
	expenseNumbers := make([]int, 0)

	// loading the numbers into slice
    for scanner.Scan() {
		currentNumber := 0;
		currentNumber, _  = strconv.Atoi(scanner.Text())
		expenseNumbers = append(expenseNumbers, currentNumber)
	}
	
	// Part One 
	targetNumber := 2020
	partOneValue := PartOne(expenseNumbers, targetNumber)
	fmt.Printf("Part One number: %d\n", partOneValue)

	// Part Two 
	partTwoValue := PartTwo(expenseNumbers, targetNumber)
	fmt.Printf("Part Two number: %d\n", partTwoValue)
}