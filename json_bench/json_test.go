package json_bench

import (
	"encoding/json"
	jsoniter "github.com/json-iterator/go"
	"strconv"
	"testing"
)

var iter = jsoniter.ConfigCompatibleWithStandardLibrary

func BenchmarkStdMarshal(t *testing.B) {
	v := map[string]int{}
	for i := 0; i < 100; i++ {
		v[strconv.Itoa(i)] = i
	}

	for i := 0; i < t.N; i++ {
		if _, err := json.Marshal(v); err != nil {
			panic(err)
		}
	}
}

func BenchmarkIteratorMarshal(t *testing.B) {
	v := map[string]int{}
	for i := 0; i < 100; i++ {
		v[strconv.Itoa(i)] = i
	}

	for i := 0; i < t.N; i++ {
		if _, err := iter.Marshal(v); err != nil {
			panic(err)
		}
	}
}

func BenchmarkStdUnMarshal(t *testing.B) {
	v := map[string]int{}
	for i := 0; i < 100; i++ {
		v[strconv.Itoa(i)] = i
	}

	r, _ := json.Marshal(v)
	var to map[string]int

	for i := 0; i < t.N; i++ {
		if err := json.Unmarshal(r, &to); err != nil {
			panic(err)
		}
	}
}

func BenchmarkIteratorUnMarshal(t *testing.B) {
	v := map[string]int{}
	for i := 0; i < 100; i++ {
		v[strconv.Itoa(i)] = i
	}

	r, _ := json.Marshal(v)
	var to map[string]int

	for i := 0; i < t.N; i++ {
		if err := iter.Unmarshal(r, &to); err != nil {
			panic(err)
		}
	}
}
