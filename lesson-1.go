package main

import (
	"fmt"
	"strconv"
)

func IsClabeValid(clabe string) bool {

	var multiplication []int = []int{3, 7, 1, 3, 7, 1, 3, 7, 1, 3, 7, 1, 3, 7, 1, 3, 7}

	var clabeSum int
	var modMinus10 int
	const storiPrefix string = "6461802244"
	var indexMult int = 0
	var multNum int

	if len(clabe) != 18 {

		return false

	}

	for _, v := range clabe[:17] {
		clabeNum, err := strconv.Atoi(string(v))
		if err != nil {
			fmt.Println("Fallo tu str to int")
			return false
		}
		multNum = multiplication[indexMult]
		//fmt.Println("Multiplying", clabeNum, "*", multNum)
		clabeMult := clabeNum * multNum
		clabeSum += clabeMult
		indexMult++
	}

	mod10 := clabeSum % 10

	modMinus10 = (10 - mod10) % 10

	finalDigit, err := strconv.Atoi(string(clabe[17]))
	if err != nil {
		return false
	}

	fmt.Println(finalDigit, "is not", modMinus10)
	return int(finalDigit) == modMinus10

}

func main() {
	fmt.Println(IsClabeValid("646180224445015833")) // true
	fmt.Println(IsClabeValid("646180224445015837")) // false
	fmt.Println(IsClabeValid("646"))                // false
}

// testing pull requests
