package main

import "testing"
import 	"fmt"

func TestMultiplica(t *testing.T) {
	req, err := http.NewRequest("GET", "v1/multiplica/4/5", nil)
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}
		
}
