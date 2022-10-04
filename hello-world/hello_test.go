package main
import "testing"

func TestHello(t *testing.T){
	got := Hello("Megan")
	want := "Hello, Megan!"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}