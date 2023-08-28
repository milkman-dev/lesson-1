package main

import (
	"fmt"
)

func IsClabeValid(clabe string) bool {

	multiplication := []int{3, 7, 1, 3, 7, 1, 3, 7, 1, 3, 7, 1, 3, 7, 1, 3, 7}

	var clabeSum int
	var modMinus10 int
	const storiPrefix string = "6461802244"
	var indexMult int = 0
	var multNum int

	if len(clabe) != 18 {

		return false

	}

	for a := range clabe[:17] {
		clabeNum := int(a)
		multNum = multiplication[indexMult]
		fmt.Println("Multiplying", clabeNum, "*", multNum)
		clabeMult := clabeNum * multNum
		clabeSum += clabeMult
		indexMult++
		if indexMult > 17 {
			break
		}
	}

	mod10 := clabeSum % 10
	// var clabePrefix bool = clabe[0:10] == storiPrefix

	modMinus10 = (10 - mod10) % 10

	return int(clabe[18]-'0') == modMinus10

}

func main() {
	fmt.Println(IsClabeValid("646180224445015833")) // true
	fmt.Println(IsClabeValid("646180224445015837")) // false
	fmt.Println(IsClabeValid("646"))                // false
}
