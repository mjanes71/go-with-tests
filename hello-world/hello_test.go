package main
import "testing"

func TestHello(t *testing.T){
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Megan")
		want := "Hello, Megan!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})
}

//this is a helper function in the test suite that you
//can feed into the main tests
func assertCorrectMessage(t testing.TB, got, want string){
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}