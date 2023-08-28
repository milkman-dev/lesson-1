package main

import (
	"fmt"
)

func IsClabeValid(clabe string) bool {
	// pass
}

func main() {
	fmt.Print(IsClabeValid("646180224445015833")) // true
	fmt.Print(IsClabeValid("646180224445015837")) // false
	fmt.Print(IsClabeValid("646"))                // false
}
