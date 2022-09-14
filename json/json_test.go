package json

import (
	"bytes"
	"encoding/json"
	"testing"
)

var testCase = map[string]interface{}{
	"foo": []string{"1", "2", "3"},
	"bar": "baz",
	"test": []interface{}{
		[]string{"test1", "test2", "test3"},
		[]string{"test1", "test2", "test3"},
		[]string{"test1", "test2", "test3"}},
}

func BenchmarkMarshal(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(testCase)
	}
}

func BenchmarkEncoder(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf360 := new(bytes.Buffer)
		enc360 := json.NewEncoder(buf360)
		enc360.Encode(testCase)
	}
}
