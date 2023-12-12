package main

import "testing"

func TestExampleFunction(t *testing.T) {
    // TODO: Implement tests for the student's code
    // Example test (adjust according to the actual assignment)
    expected := "Hello"
    if got := HelloFunction(); got != expected {
        t.Errorf("HelloFunction() = %v, want %v", got, expected)
    }
}


func TestAdd(t *testing.T) {
    expecred :=5
    if got := Add(2,3); got != expecred {
        t.Errorf("Add(2,3) = %v, want %v" , got, expecred)
    }
}