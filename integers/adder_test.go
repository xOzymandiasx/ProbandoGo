package integers

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	//Output: 6
}

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expexted := 4

	if sum != expexted {
		t.Errorf("expexted %d but got %d", expexted, sum)
	}
}
