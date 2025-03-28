package tipcalc_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/windscale-stephen/ex4prog_go/tipcalc/internal/tipcalc"
)

func TestDisplayPrompt(t *testing.T) {
	t.Parallel()
	prompt := "What is the bill? $"
	var w bytes.Buffer
	err := tipcalc.DisplayPrompt(&w, prompt)
	if err != nil {
		t.Fatal(err)
	}
	got := w.String()
	if got != prompt {
		t.Errorf("expected prompt \"%s\", got \"%s\"\n", prompt, got)
	}
}

func TestReadFPNumber(t *testing.T) {
	t.Parallel()
	type testCase struct {
		fp   string
		want float64
	}
	testCases := []testCase{
		{fp: "10.00\n", want: 10.0},
		{fp: "0.0\n", want: 0.0},
		{fp: "1.5\n", want: 1.5},
		{fp: "-1.0\n", want: -1.0},
		{fp: "10\n", want: 10.0},
		{fp: "1\n", want: 1.0},
		{fp: " 1.0 \n", want: 1.0},
		{fp: "1.0 ", want: 1.0},
	}
	for _, tc := range testCases {
		reader := strings.NewReader(tc.fp)
		got, err := tipcalc.ReadFPNumber(reader)
		if err != nil {
			t.Fatal(err)
		}
		if got != tc.want {
			t.Errorf("for fp %s want %.2f, got %f", tc.fp,
				tc.want, got)
		}
	}
}

func TestReadFPNumberInvalid(t *testing.T) {
	t.Parallel()
	testCases := []string{
		"abcd\n",
		"",
		"\n",
	}
	for _, tc := range testCases {
		reader := strings.NewReader(tc)
		_, err := tipcalc.ReadFPNumber(reader)
		if err == nil {
			t.Fatalf("for input \"%s\" expected error, got nil", tc)
		}
	}
}

func TestTip(t *testing.T) {
	t.Parallel()
	type testCase struct {
		bill   float64
		tipPct float64
		want   float64
	}
	testCases := []testCase{
		{bill: 10.00, tipPct: 0.0, want: 0.0},
		{bill: 10.00, tipPct: 15.0, want: 1.50},
		{bill: 11.25, tipPct: 15.0, want: 1.69},
		{bill: 15.00, tipPct: 100.0, want: 15.00},
	}
	for _, tc := range testCases {
		got, err := tipcalc.Tip(tc.bill, tc.tipPct)
		if err != nil {
			t.Fatal(err)
		}
		if got != tc.want {
			t.Errorf("bill: %.2f, tipPct: %.2f, want tip %.2f, got %f",
				tc.bill, tc.tipPct, tc.want, got)
		}
	}
}

func TestTipInvalid(t *testing.T) {
	t.Parallel()
	type testCase struct {
		bill   float64
		tipPct float64
	}
	testCases := []testCase{
		{bill: 0.0, tipPct: 0.0},
		{bill: -1.0, tipPct: 0.0},
		{bill: 10.0, tipPct: -1.0},
	}
	for _, tc := range testCases {
		_, err := tipcalc.Tip(tc.bill, tc.tipPct)
		if err == nil {
			t.Fatalf("bill: %.2f, tipPct: %.2f, expecting error, got nil",
				tc.bill, tc.tipPct)
		}
	}
}

func TestTotal(t *testing.T) {
	t.Parallel()
	type testCase struct {
		bill float64
		tip  float64
		want float64
	}
	testCases := []testCase{
		{bill: 100.00, tip: 10.00, want: 110.00},
		{bill: 1.00, tip: 0.01, want: 1.01},
	}
	for _, tc := range testCases {
		got, err := tipcalc.Total(tc.bill, tc.tip)
		if err != nil {
			t.Fatal(err)
		}
		if got != tc.want {
			t.Errorf("bill: %.2f, tip: %.2f, want total %.2f, got %f",
				tc.bill, tc.tip, tc.want, got)
		}
	}
}

func TestTotalInvalid(t *testing.T) {
	t.Parallel()
	type testCase struct {
		bill float64
		tip  float64
	}
	testCases := []testCase{
		{bill: 0.0, tip: 0.0},
		{bill: -1.0, tip: 0.0},
		{bill: 10.00, tip: -1.0},
	}
	for _, tc := range testCases {
		_, err := tipcalc.Total(tc.bill, tc.tip)
		if err == nil {
			t.Fatalf("bill: %.2f, tip: %.2f, expecting error, got nil",
				tc.bill, tc.tip)
		}
	}
}
