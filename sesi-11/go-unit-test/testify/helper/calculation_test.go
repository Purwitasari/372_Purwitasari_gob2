package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFailSum(t *testing.T) {
	result := Sum(10, 10)

	require.Equal(t, 40, result, "Result has to be 40")

	fmt.Println("TestFailSum Eksekusi Terhenti")
}

func TestSum(t *testing.T) {
	result := Sum(10, 10)

	assert.Equal(t, 20, result, "Result has to be 20")

	fmt.Println("TestSum Eksekusi Terhenti")
}

/* Running

go test ./helper
go test -v ./helper
go test -v ./helper -run=TestSum

*/
