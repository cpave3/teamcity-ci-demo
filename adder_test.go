package main

import "testing"

func TestAdd(t *testing.T) {
  result := Add(2, 3)
  if result != 5 {
    t.Errorf("Result was incorrect, got %d, want %d", result, 5)
  }
}
