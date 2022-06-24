package calculator

import (
	"errors"
	"reflect"
	"testing"
)

func TestAdder(t *testing.T) {

	checkSums := func(t *testing.T, got, want int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	checkErrors := func(t *testing.T, got, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}

		if got.Error() != want.Error() {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("make the sums of input", func(t *testing.T) {
		got, _ := Add("2,2")
		want := 4

		checkSums(t, got, want)
	})

	t.Run("safetly sum empty string", func(t *testing.T) {
		got, _ := Add("")
		want := 0

		checkSums(t, got, want)
	})

	t.Run("make the sums of new lines between numbers", func(t *testing.T) {
		got, _ := Add("1\n2,3")
		want := 6

		checkSums(t, got, want)
	})

	t.Run("the input '\n' is not valid ", func(t *testing.T) {
		_, err := Add("1,\n")
		want := errors.New("the input is not valid")

		checkErrors(t, err, want)
	})

	t.Run("make the sums of different delimiters", func(t *testing.T) {
		got, _ := Add("//;\n1;2")
		want := 3

		checkSums(t, got, want)
	})
	t.Run("negative number will throw an exception", func(t *testing.T) {
		_, err := Add("-2")
		want := errors.New("negatives not allowed")

		checkErrors(t, err, want)
	})

}
