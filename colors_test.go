package hexcolor_test

import (
	"image/color"

	"github.com/GodsBoss/hexcolor"

	"testing"
)

func TestNonParsableRGBStrings(t *testing.T) {
	inputs := []string{
		"asdfasdfasdf",
		"999",
		"ffaabb",
		"#ff",
		"#4435",
		"#12345",
		"#bb0099dd",
	}

	for i := range inputs {
		input := inputs[i]
		t.Run(
			input,
			func(t *testing.T) {
				col, err := hexcolor.ParseRGB(input)
				if col != nil {
					t.Errorf("expected no color, but got %+v", col)
				}
				if err == nil {
					t.Errorf("missing error")
				}
			},
		)
	}
}

func TestNonParsableRGBAStrings(t *testing.T) {
	inputs := []string{
		"asdfasdfasdf",
		"999",
		"ffaabb",
		"#ff",
		"#443",
		"#12345",
		"#bb0099d",
	}

	for i := range inputs {
		input := inputs[i]
		t.Run(
			input,
			func(t *testing.T) {
				col, err := hexcolor.ParseRGBA(input)
				if col != nil {
					t.Errorf("expected no color, but got %+v", col)
				}
				if err == nil {
					t.Errorf("missing error")
				}
			},
		)
	}
}

func TestParsing(t *testing.T) {
	testcases := []parsingTestCase{
		{
			input: "#fFfF",
			parse: hexcolor.ParseRGBA,
			R:     0xFFFF,
			G:     0xFFFF,
			B:     0xFFFF,
			A:     0xFFFF,
		},
		{
			input: "#12345678",
			parse: hexcolor.ParseRGBA,
			R:     0x0880,
			G:     0x1890,
			B:     0x28a0,
			A:     0x7878,
		},
		{
			input: "#abc",
			parse: hexcolor.ParseRGB,
			R:     0xaaaa,
			G:     0xbbbb,
			B:     0xcccc,
			A:     0xffff,
		},
		{
			input: "#fedcba",
			parse: hexcolor.ParseRGB,
			R:     0xfefe,
			G:     0xdcdc,
			B:     0xbaba,
			A:     0xffff,
		},
	}
	for i := range testcases {
		testcases[i].Run(t)
	}
}

type parsingTestCase struct {
	input      string
	parse      func(string) (color.Color, error)
	R, G, B, A uint32
}

func (testcase parsingTestCase) Run(t *testing.T) {
	t.Run(
		testcase.input,
		func(t *testing.T) {
			col, err := testcase.parse(testcase.input)
			if err != nil {
				t.Errorf("expected no error, but got %+v", err)
			}
			if col == nil {
				t.Fatalf("expected color")
			}
			R, G, B, A := col.RGBA()
			assertColor(t, "red", testcase.R, R)
			assertColor(t, "green", testcase.G, G)
			assertColor(t, "blue", testcase.B, B)
			assertColor(t, "alpha", testcase.A, A)
		},
	)
}

func assertColor(t *testing.T, componentName string, expected, actual uint32) {
	if expected != actual {
		t.Errorf("expected color component %s to be %d, but got %d", componentName, expected, actual)
	}
}
