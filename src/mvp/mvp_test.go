package mvp

import (
	"testing"
)

func TestMVP(t *testing.T) {
	tests := []struct {
		name      string
		operation string
		key       string
		value     string
	}{
		{
			name:      "put fred",
			operation: "put",
			key:       "fred",
			value:     "wilma",
		},
		{
			name:      "get fred",
			operation: "get",
			key:       "fred",
			value:     "wilma",
		},
	}
	var cache = New()
	defer cache.Close()

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			var value string
			var present bool
			var err error

			switch tt.operation {
			case "put":
				err = cache.Put(tt.key, tt.value)
				if err != nil {
					t.Fatalf("put failed and returned %v\n", err)
				}
				t.Logf("put: key = %q, value = %q\n", tt.key, tt.value)

			case "get":
				value, present = cache.Get(tt.key)
				t.Logf("get: key = %q, present = %t, value = %q\n", tt.key, present, value)
				if present != true {
					t.Fatalf("get didn't return present == true, value was %q\n", value)
				}
				if value != tt.value {
					t.Fatalf("get didn't return wilma, but instaed %q\n", value)
				}

			}

		})
	}
}

func BenchmarkMVP(b *testing.B) {
	var cache = New()
	defer cache.Close()

	cache.Put("fred", "wilma")
	for i := 0; i < b.N; i++ {
		cache.Get("fred")
	}
}

func TestCrash(t *testing.T) {
	var cache = New()
	defer cache.Close()

	cache.Put("fred", "wilma")
	for i := 0; i < 2; i++ {
		cache.Get("fred")
	}
}
