package helper

import (
	"fmt"
	"testing"
)

// 1
// func TestFailSum(t *testing.T) {
// 	result := Sum(10, 10)

// 	if result != 40 {
// 		t.Fail()
// 	}

// 	fmt.Println("TestFailSum Eksekusi Terhenti")
// }

// 2
// func TestFailSum(t *testing.T) {
// 	result := Sum(10, 10)

// 	if result != 40 {
// 		t.FailNow()
// 	}

// 	fmt.Println("TestFailSum Eksekusi Terhenti")

// }

//3

// func TestFailSum(t *testing.T) {
// 	result := Sum(10, 10)

// 	if result != 40 {
// 		t.Error("Result has to be 40")
// 	}

// 	fmt.Println("TestFailSum Eksekusi Terhenti")
// }

//4

func TestFailSum(t *testing.T) {
	result := Sum(10, 10)

	if result != 40 {
		t.Fatal("Result has to be 40")
	}

	fmt.Println("TestFailSum Eksekusi Terhenti")
}

func TestSum(t *testing.T) {
	result := Sum(10, 10)

	if result != 20 {
		// panic("Result should be 20")
		t.FailNow()
	}

	fmt.Println("TestSum Eksekusi Terhenti")
}

/* Running

go test ./helper
go test -v ./helper
go test -v ./helper -run=TestSum
*/