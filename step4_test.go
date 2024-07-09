package main

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func Test_StartTest4(t *testing.T) {
	fmt.Println()
	log.Println("Testing step 4")
}

func TestValidate_Step4_ValidJSON1_ReturnTrue(t *testing.T) {
	log.Println("TestValidate_Step4_ValidJSON1_ReturnTrue")
	file, _ := os.Open("tests/step4/valid.json")
	defer file.Close()

	isValid := ValidateJSON(file)

	if !isValid {
		t.Errorf("Expected true, got false")
	}
}

func TestValidate_Step4_ValidJSON2_ReturnTrue(t *testing.T) {
	log.Println("TestValidate_Step4_ValidJSON2_ReturnTrue")
	file, _ := os.Open("tests/step4/valid2.json")
	defer file.Close()

	isValid := ValidateJSON(file)

	if !isValid {
		t.Errorf("Expected true, got false")
	}
}

func TestValidate_Step4_InvalidJSON1_ReturnFalse(t *testing.T) {
	log.Println("TestValidate_Step4_InvalidJSON1_ReturnFalse")
	file, _ := os.Open("tests/step4/invalid.json")
	defer file.Close()

	isValid := ValidateJSON(file)

	if isValid {
		t.Errorf("Expected false, got true")
	}
}
