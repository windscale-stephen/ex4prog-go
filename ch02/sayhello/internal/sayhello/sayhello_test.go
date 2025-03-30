package sayhello_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/windscale-stephen/ex4prog_go/ch02/sayhello/internal/sayhello"
)

func TestDisplayPrompt(t *testing.T) {
	t.Parallel()
	prompt := "What is your name? "
	var w bytes.Buffer
	err := sayhello.DisplayPrompt(&w, prompt)
	if err != nil {
		t.Fatal(err)
	}
	got := w.String()
	if got != prompt {
		t.Errorf("expected prompt \"%s\", got \"%s\"\n", prompt, got)
	}
}

func TestMakeGreeting(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name string
		want string
	}
	testCases := []testCase{
		{name: "Brian", want: "Hello, Brian, nice to meet you!\n"},
		{name: "Fred Flintstone", want: "Hello, Fred Flintstone, nice to meet you!\n"},
	}
	for _, tc := range testCases {
		got := sayhello.MakeGreeting(tc.name)
		if got != tc.want {
			t.Errorf("for name \"%s\" expected greeting \"%s\", got \"%s\"", tc.name, tc.want, got)
		}
	}
}

func TestReadName(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name string
		want string
	}
	testCases := []testCase{
		{name: "Brian", want: "Brian"},
		{name: " \tBrian", want: "Brian"},
		{name: "Brian \t", want: "Brian"},
		{name: " \tBrian \t", want: "Brian"},
		{name: "Fred Flintstone", want: "Fred Flintstone"},
		{name: "\t Fred Flintstone", want: "Fred Flintstone"},
		{name: "Fred Flintstone\t ", want: "Fred Flintstone"},
		{name: " \tFred \t Flintstone\t ", want: "Fred Flintstone"},
		{name: " \tFred\t \tStoneAge \t Flintstone", want: "Fred StoneAge Flintstone"},
	}
	for _, tc := range testCases {
		reader := strings.NewReader(tc.name)
		got, err := sayhello.ReadName(reader)
		if err != nil {
			t.Fatal(err)
		}
		if got != tc.want {
			t.Errorf("for name \"%s\" want \"%s\", got \"%s\"", tc.name, tc.want, got)
		}
	}
}
