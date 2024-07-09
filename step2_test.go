package main

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func Test_StartTest2(t *testing.T) {
	fmt.Println()
	log.Println("Testing step 2")
}

func TestValidate_Step2_ValidJSON1_ReturnTrue(t *testing.T) {
	log.Println("TestValidate_Step2_ValidJSON1_ReturnTrue")
	file, _ := os.Open("tests/step2/valid.json")
	defer file.Close()

	isValid := ValidateJSON(file)

	if !isValid {
		t.Errorf("Expected true, got false")
	}
}

func TestValidate_Step2_ValidJSON2_ReturnTrue(t *testing.T) {
	log.Println("TestValidate_Step2_ValidJSON2_ReturnTrue")
	file, _ := os.Open("tests/step2/valid2.json")
	defer file.Close()

	isValid := ValidateJSON(file)

	if !isValid {
		t.Errorf("Expected true, got false")
	}
}

func TestValidate_Step2_InvalidJSON1_ReturnFalse(t *testing.T) {
	log.Println("TestValidate_Step2_InvalidJSON1_ReturnFalse")
	file, _ := os.Open("tests/step2/invalid.json")
	defer file.Close()

	isValid := ValidateJSON(file)

	if isValid {
		t.Errorf("Expected false, got true")
	}
}

func TestValidate_Step2_InvalidJSON2_ReturnFalse(t *testing.T) {
	log.Println("TestValidate_Step2_InvalidJSON2_ReturnFalse")
	file, _ := os.Open("tests/step2/invalid2.json")
	defer file.Close()

	isValid := ValidateJSON(file)

	if isValid {
		t.Errorf("Expected false, got true")
	}
}
