package iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	var charRepeated string
	for i := 0; i < 5; i++ {
		charRepeated += "a"
	}
	fmt.Println(charRepeated)
	//Output: aaaaa
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", repeated, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 7)
	}
}
