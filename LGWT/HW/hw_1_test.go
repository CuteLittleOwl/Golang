package main

import (
	"testing"
)

func TestHW(t *testing.T) {
	t.Run("Country is Russia and the ruller is Putin", func(t *testing.T) {
		got := ruller("USA", "Putin")
		want := "Putin is the President"
		if got != want {
			assertCorrectMessage(t, got, want)
		}
	})
	t.Run("Country is Germany and the ruller is Merkel", func(t *testing.T) {
		got := ruller("Germany", "Merkel")
		want := "Merkel is the Chancellor"
		if got != want {
			assertCorrectMessage(t, got, want)
		}
	})
	t.Run("Country is Britain and the ruller is Elisabeth", func(t *testing.T) {
		got := ruller("Britain", "Elisabeth")
		want := "Elisabeth is the Queen"
		if got != want {
			assertCorrectMessage(t, got, want)
		}
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	t.Errorf("got %q want %q", got, want)
}
