package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTableFactorial(t *testing.T) {
tests := []struct{
	name string
	request int
	expected int
}{
	{
	name: "3",
	request: 3,
	expected: 6,
	},
	{
	name: "4",
	request: 4,
	expected: 24,
	},
	{
	name: "5",
	request: 5,
	expected: 120,
	},
}
fmt.Println(tests)

for _, test := range tests {
	t.Run(test.name, func(t *testing.T) {
		result := Factorial(test.request)
		require.Equal(t, test.expected, result)
	})
}

}

// go test -v -run=TestTableFactorial
