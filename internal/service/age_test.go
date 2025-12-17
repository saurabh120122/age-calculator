package service

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {
	dob := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	age := calculateAge(dob)

	if age <= 0 {
		t.Fatalf("invalid age: %d", age)
	}
}
