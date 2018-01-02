package factory

import (
	"testing"
)

/*
go test -v
*/

func Test_factory(t *testing.T) {
	v := getFormat("JSON").Run()
	t.Log(v)
}

func Test_factory2(t *testing.T) {
	v := getFormat("XML").Run()
	t.Log(v)
}

/*
go test -test.bench=".*"
*/

func Benchmark_factory2(b *testing.B) {
	b.StopTimer()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		getFormat("XML").Run()
	}
}
