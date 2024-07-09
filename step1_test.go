package main

import (
	"log"
	"os"
	"testing"
)

func Test_StartTest1(t *testing.T) {
	log.Println("Testing step 1")
}

func TestValidate_Step1_ValidJSON1_ReturnTrue(t *testing.T) {
	log.Println("TestValidate_Step1_ValidJSON1_ReturnTrue")
	file, _ := os.Open("tests/step1/valid.json")
	defer file.Close()

	isValid := ValidateJSON(file)

	if !isValid {
		t.Errorf("Expected true, got false")
	}
}

func TestValidate_Step1_InvalidJSON1_ReturnFalse(t *testing.T) {
	log.Println("TestValidate_Step1_InvalidJSON1_ReturnFalse")
	file, _ := os.Open("tests/step1/invalid.json")
	defer file.Close()

	isValid := ValidateJSON(file)

	if isValid {
		t.Errorf("Expected false, got true")
	}
}
