package test2

import "testing"

func payload() {
	for i := 0; i < 100000; i++ {
		
	}	
}

func BenchmarkDirect(b *testing.B) {
	payload()
}

func BenchmarkDeffer(b *testing.B) {
	success := false
	defer func() {
		success = true	
	}()
	payload()
}