package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to You", func(t *testing.T) {
		got := Hello("KB", "")
		want := "Hello, KB"

		assetCorrectMessage(t, got, want)
	})

	t.Run("Saying hello to Noone", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assetCorrectMessage(t, got, want)
	})

	t.Run("Saying hello in Spanish", func(t *testing.T) {
		got := Hello("KB", "Spanish")
		want := "Hola, KB"

		assetCorrectMessage(t, got, want)
	})

	t.Run("Saying hello in French", func(t *testing.T) {
		got := Hello("KB", "French")
		want := "Bonjour, KB"

		assetCorrectMessage(t, got, want)
	})
}

func assetCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
