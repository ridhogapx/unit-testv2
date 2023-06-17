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

// func BenchmarkCase(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		*expected = Person{
// 			Name: "John",
// 			Nick: "Doe",
// 			Age:  40,
// 		}
// 	}
// }

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
