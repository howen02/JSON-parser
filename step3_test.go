package main

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func Test_StartTest3(t *testing.T) {
	fmt.Println()
	log.Println("Testing step 3")
}

func TestValidate_Step3_ValidJSON1_ReturnTrue(t *testing.T) {
	log.Println("TestValidate_Step3_ValidJSON1_ReturnTrue")
	file, _ := os.Open("tests/step3/valid.json")
	defer file.Close()

	isValid := ValidateJSON(file)

	if !isValid {
		t.Errorf("Expected true, got false")
	}
}

func TestValidate_Step3_InvalidJSON1_ReturnTrue(t *testing.T) {
	log.Println("TestValidate_Step3_InvalidJSON1_ReturnTrue")
	file, _ := os.Open("tests/step3/invalid.json")
	defer file.Close()

	isValid := ValidateJSON(file)

	if isValid {
		t.Errorf("Expected false, got true")
	}
}
