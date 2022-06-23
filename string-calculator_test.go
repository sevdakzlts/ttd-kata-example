package calculator

import (
	"reflect"
	"testing"
)

func TestAdder(t *testing.T) {

	checkSums := func(t *testing.T, got, want int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of input", func(t *testing.T) {
		got := Add("2,2")
		want := 4

		checkSums(t, got, want)
	})

	t.Run("safetly sum empty string", func(t *testing.T) {
		got := Add("")
		want := 0

		checkSums(t, got, want)
	})

	t.Run("make the sums of new lines between numbers", func(t *testing.T) {
		got := Add("1\n2,3")
		want := 6

		checkSums(t, got, want)
	})

}
