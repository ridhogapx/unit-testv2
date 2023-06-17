package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var expected Person = Person{
	Name: "John",
	Nick: "Doe",
	Age:  40,
}

var testVal *Person = new(Person)

// Benchmark test
func BenchmarkCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChangePerson(&expected)
	}
}

func BenchmarkWithSub(b *testing.B) {
	b.Run("First", func(b *testing.B) {
		*testVal = Person{
			Name: "Dio",
			Nick: "Giorvani",
			Age:  26,
		}
		for i := 0; i < b.N; i++ {
			ChangePerson(testVal)
		}
	})

	b.Run("Second", func(b *testing.B) {
		*testVal = Person{
			Name: "Gerald",
			Nick: "Wijaya",
			Age:  26,
		}
		for i := 0; i < b.N; i++ {
			ChangePerson(testVal)
		}
	})
}

func TestFirst(t *testing.T) {
	// first case

	test_person := Person{
		Name: "Budi Setiawan",
		Nick: "Budi",
		Age:  34,
	}

	pointer_test := &test_person

	result := ChangePerson(pointer_test)
	t.Run("first", func(t *testing.T) {
		assert.Equal(t, &expected, result, "Result is not as expected")
	})
}
