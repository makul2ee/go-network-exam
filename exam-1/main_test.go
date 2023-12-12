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

<<<<<<< HEAD

func TestAdd(t *testing.T) {
    expecred :=5
    if got := Add(2,3); got != expecred {
        t.Errorf("Add(2,3) = %v, want %v" , got, expecred)
=======
func TestAdd(t *testing.T) {
    expected := 5
    if got := Add(2, 3); got != expected {
        t.Errorf("Add(2, 3) = %v, want %v", got, expected)
>>>>>>> 8fe3951eaf58456ef6f7cecfed9fab5c4c54cc8f
    }
}