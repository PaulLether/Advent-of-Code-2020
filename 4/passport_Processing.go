package main
import (
    "bufio"
    "log"
	"os"
	"fmt"
	"strings"
	"regexp"
	"strconv"
)

// checks if the sting contains certain sub strings 
func Checker(potentialPassport string)  bool{
	mustContain:= []string {"byr","iyr", "eyr", "hgt", "hcl", "ecl", "pid"}	
	for s := range mustContain{
		if !strings.Contains(potentialPassport, mustContain[s]){
			return false
		}
	}
	return true
}

func valid(potentialPassport string) bool{
	r:= regexp.MustCompile(`(([a-z]{3}?):([\d]+))*`)
	matches := r.FindStringSubmatch(potentialPassport)
	// for i:= 0 ; i < len(matches) ; i++{
	// fmt.Printf("matches[%d] = %s ", i, matches[i])
	// }
	
	// fmt.Printf("size of matches = %d\n", len(matches))
	for  i := 0; i + 1< len(matches); i++{
		if strings.Contains(matches[i],"byr"){
			// fmt.Printf("%s :", matches[i])
			matches[i] = strings.ReplaceAll(matches[i], "byr:", "")
			intValue, nullValue := strconv.Atoi(matches[i])
			// fmt.Printf("int value = %d\n", intValue)
			if (intValue > 2002 && 1920 > intValue )|| nil != nullValue{
				return false
			} 
		} else if strings.Contains(matches[i],"iyl"){
			matches[i] = strings.ReplaceAll(matches[i], "iyl:", "")
			intValue, nullValue := strconv.Atoi(matches[i])
			if intValue > 2020 && 2010 > intValue || nil != nullValue{
				return false
			}
		} else if strings.Contains(matches[i],"eyr"){
			matches[i] = strings.ReplaceAll(matches[i], "eyr:", "")
			intValue, nullValue := strconv.Atoi(matches[i])
			if intValue > 2020 && 2010 > intValue || nullValue != nil{
				return false
			}
		} else if strings.Contains(matches[i],"hgt"){
			matches[i] = strings.ReplaceAll(matches[i], "hgt:", "")
			if strings.Contains(matches[i], "cm"){
				matches[i] = strings.ReplaceAll(matches[i], "cm", "")
			 	intValue, nullValue:= strconv.Atoi(matches[i])
				if intValue > 193 && 150 > intValue|| nullValue != nil{
					return false
				}
			} else if strings.Contains(matches[i], "in"){
				matches[i] = strings.ReplaceAll(matches[i], "in:", "")
			 	intValue, nullValue := strconv.Atoi(matches[i])
				if intValue > 76 && 59 > intValue || nullValue != nil{
					return false
				}
			} 
		} else if strings.HasPrefix(matches[i],"hcl:"){
				matches[i] = strings.ReplaceAll(matches[i], "hcl:", "")
				checkExpression:= regexp.MustCompile(`#(([a-z]{6})|([\d]{6}))*`)
				if !checkExpression.MatchString(matches[i]){
					return false
				}
		}	else if strings.HasPrefix(matches[i],"ecl:"){
			matches[i] = strings.ReplaceAll(matches[i], "ecl:", "")
			matches[i] = strings.ReplaceAll(matches[i], " ", "")

			mustContain:= []string {"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}	
			for s := range mustContain{
				if mustContain[s] != matches[i]{
				return false
				}
			}
		} else if strings.HasPrefix(matches[i],"pid"){
			matches[i] = strings.ReplaceAll(matches[i], "pid:", "")
			convertedToNum, err := strconv.Atoi(matches[i])
			convertedToNum++
			if len(matches[i]) != 9 && err != nil{
				return false 
			}
		}
	}
	return true
}

func main() {
	// reading file
	file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
	scanner := bufio.NewScanner(file)

	// Part One: checking the number of valid passports
	numOfValidPassports := 0
	numOfDataVPassports := 0 
	potentialPassport := ""
	index := 0
	for scanner.Scan(){
		if scanner.Text() != ""{
			if potentialPassport == ""{
				potentialPassport = scanner.Text();
			}
			potentialPassport += " " + scanner.Text();
		} else {
			// fmt.Printf("%s\n\n", potentialPassport)
			if Checker(potentialPassport){
				// fmt.Printf("%s\n\n", potentialPassport)
				if valid(potentialPassport){
					numOfDataVPassports++
				}
				numOfValidPassports += 1
			}
			potentialPassport = ""
		}
		index++
	}
	fmt.Printf("Part One:\n Number of valid passports = %d\n", numOfValidPassports)
	
	// Part Two: checking the number of valid passports
	fmt.Printf("Part Two:\n Number of data valid passports = %d\n", numOfDataVPassports)


	
	
	
	
}