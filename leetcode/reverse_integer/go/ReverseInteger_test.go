package main

import (
	"fmt"
	"testing"
	"time"
)

func TestReverseInteger(test *testing.T){

	if ReverseInteger(123) != 321{
		test.Errorf("Failed")
	}
	if ReverseInteger(-123) != -321{
		test.Errorf("Failed")
	}
	if ReverseInteger(120) != 21{
		test.Errorf("Failed")
	}
	if ReverseInteger(0) != 0{
		test.Errorf("Failed")
	}

	// function execution time
	startTime := time.Now().UnixNano()
	ReverseInteger(123)
	endTime := time.Now().UnixNano()
	fmt.Printf("Time elapsed: %d nanoseconds\n", endTime - startTime)
}
