package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
)

func IsClabeValid(clabe string) bool {

	var multiplication []int = []int{3, 7, 1, 3, 7, 1, 3, 7, 1, 3, 7, 1, 3, 7, 1, 3, 7}

	var clabeSum int
	var modMinus10 int
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

	return int(finalDigit) == modMinus10

}

type (
	validationRequest struct {
		Clabe string `json:"clabe"`
	}

	validationResponse struct {
		Valid bool `json:"valid"`
	}
)

func HandleRequest(ctx context.Context, event validationRequest) (validationResponse, error) {
	fmt.Printf("%v\n", event)
	isValid := IsClabeValid(event.Clabe)
	resp := validationResponse{
		Valid: isValid,
	}
	return resp, nil
}

func main() {

	lambda.Start(HandleRequest)

}
