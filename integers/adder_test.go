package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expexted := 4

	if sum != expexted {
		t.Errorf("expexted %d but got %d", expexted, sum)
	}
}
