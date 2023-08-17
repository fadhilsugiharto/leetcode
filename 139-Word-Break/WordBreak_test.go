package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
)

//run test:
//go test
//
//details:
//go test -v
//
//specific:
//go test -v -run=TestNamaFunction

// To run benchmark:
// go test -v -bench=.
// To run only benchmark
// go test -v -run=NonExistantFunc -bench=.

func TestMain(m *testing.M) {
	// performs operation before the test
	fmt.Println("Test starting...")

	m.Run()

	// performs operation after the test
	fmt.Println("Test finished")
}

func BenchmarkWordBreakTable(b *testing.B) {
	benchmarks := []struct {
		name          string
		requestString string
		requestDict   []string
	}{
		{
			name:          "true-test",
			requestString: "cars",
			requestDict:   []string{"car", "ca", "rs"},
		},
		{
			name:          "false-test",
			requestString: "carss",
			requestDict:   []string{"car", "ca", "rs"},
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				wordBreak(benchmark.requestString, benchmark.requestDict)
			}
		})
	}
}

func TestWordBreakTable(t *testing.T) {
	tests := []struct {
		name          string
		requestString string
		requestDict   []string
		expected      bool
		errMsg        string
	}{
		{
			name:          "true-test",
			requestString: "cars",
			requestDict:   []string{"car", "ca", "rs"},
			expected:      true,
			errMsg:        "result should be true",
		},
		{
			name:          "false-test",
			requestString: "carss",
			requestDict:   []string{"car", "ca", "rs"},
			expected:      false,
			errMsg:        "result should be false",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := wordBreak(test.requestString, test.requestDict)
			assert.Equal(t, test.expected, res, test.errMsg)
		})
	}
}

func TestWordBreakSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Test cannot run on windows")
	}

	res := wordBreak("cars", []string{"car", "ca", "rs"})

	assert.Equal(t, true, res, "result should be true")

	fmt.Println("this will still be executed if error occurs")
}

/////////////// To run subtest partially ////////////////////////
//    go test -v -run TestWordBreakSubTest/true-test ////////////
//func TestWordBreakSubTest(t *testing.T) {
//	// subtest true with assert
//	t.Run("true-test", func(t *testing.T) {
//		res := wordBreak("cars", []string{"car", "ca", "rs"})
//		assert.Equal(t, true, res, "result should be true")
//		fmt.Println("this will still be executed even though test fails")
//	})
//
//	// subtest false with require
//	t.Run("false-test", func(t *testing.T) {
//		res := wordBreak("carss", []string{"car", "ca", "rs"})
//		require.Equal(t, false, res, "result should be false")
//		fmt.Println("this will not be executed if test fails")
//	})
//}

//func TestWordBreakRequire(t *testing.T) {
//	res := wordBreak("carss", []string{"car", "ca", "rs"})
//
//	require.Equal(t, true, res, "result should be true")
//
//	fmt.Println("this will not be executed if error occurs")
//}
//
//func TestWordBreakAssert(t *testing.T) {
//	res := wordBreak("carss", []string{"car", "ca", "rs"})
//
//	assert.Equal(t, true, res, "result should be true")
//
//	fmt.Println("this will still be executed if error occurs")
//}
//
//func TestWordBreakManualTrue(t *testing.T) {
//	res := wordBreak("cars", []string{"car", "ca", "rs"})
//
//	if res != true {
//		t.Errorf("expected result is %t but got %t", true, res)
//	}
//
//}
//
//func TestWordBreakManualFalse(t *testing.T) {
//	res := wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})
//
//	if res != false {
//		t.Errorf("expected result is %t but got %t", false, res)
//	}
//
//}
