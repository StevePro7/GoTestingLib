package adriana

import (
	"testing"
)

func TestAdd(t *testing.T) {

	actual := Add(3, 2)
	expect := 5

	if actual != expect {
		t.Errorf("actual %q expect %q", actual, expect)
	}
}
