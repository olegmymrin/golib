package common

import (
	"testing"
)

type message struct {
	Output chan bool
}

var parallel = 100

func payload() {
	s := 0
	for i := 0; i < 10; i++ {
		s += i
	}
}

func BenchmarkNoChannel(b *testing.B) {
	b.StopTimer()
	go func(count int) {
		for i := 0; i < count; i++ {
			payload()
		}
	}(b.N)
	b.StartTimer()
	b.SetParallelism(parallel)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
		}
	})
}

func BenchmarkSingleChannel(b *testing.B) {
	b.StopTimer()
	input := make(chan message)
	commonOutput := make(chan bool)
	go func(input <-chan message, count int) {
		for i := 0; i < count; i++ {
			<-input
			payload()
		}
	}(input, b.N)
	b.StartTimer()
	b.SetParallelism(parallel)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			message := message{commonOutput}
			input <- message
		}
	})
}

func BenchmarkDoubleChannel(b *testing.B) {
	b.StopTimer()
	input := make(chan message)
	commonOutput := make(chan bool)
	go func(input <-chan message, count int) {
		for i := 0; i < count; i++ {
			message := <-input
			payload()
			message.Output <- true
		}
	}(input, b.N)
	b.StartTimer()
	b.SetParallelism(parallel)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			message := message{commonOutput}
			input <- message
			<-message.Output
		}
	})
}

func BenchmarkChannelPerRequest(b *testing.B) {
	b.StopTimer()
	input := make(chan message)
	go func(input <-chan message, count int) {
		for i := 0; i < count; i++ {
			message := <-input
			payload()
			message.Output <- true
		}
	}(input, b.N)
	b.StartTimer()
	b.SetParallelism(parallel)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			message := message{make(chan bool)}
			input <- message
			<-message.Output
		}
	})
}
