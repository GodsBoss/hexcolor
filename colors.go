package hexcolor

import (
	"fmt"
	"image/color"
	"regexp"
)

// ParseRGB parses a string of the form "#RGB" or "#RRGGBB" as color, where
// R, G and B are in the range of 0-f. Both lowercase and uppercase hex letters
// are accepted.
func ParseRGB(rgbHexString string) (color.Color, error) {
	if !rgbRE.MatchString(rgbHexString) {
		return nil, fmt.Errorf("RGB hex string must match #RRGGBB or #RGB, not '%s'", rgbHexString)
	}
	if len(rgbHexString) == 4 { // "#RGB"
		return ParseRGBA(rgbHexString + "F")
	}
	return ParseRGBA(rgbHexString + "FF") // "#RRGGBB"
}

// ParseRGBA is like ParseRGB, but the string is something like "#RGBA" or
// "#RRGGBBAA", with A being the alpha channel.
func ParseRGBA(rgbaHexString string) (color.Color, error) {
	if !rgbaRE.MatchString(rgbaHexString) {
		return nil, fmt.Errorf("RGB hex string must match #RRGGBBAA or #RGBA, not '%s'", rgbaHexString)
	}
	if len(rgbaHexString) == 5 { // "#RGBA"
		return color.NRGBA{
			R: singleHexToInt(rgbaHexString[1], rgbaHexString[1]),
			G: singleHexToInt(rgbaHexString[2], rgbaHexString[2]),
			B: singleHexToInt(rgbaHexString[3], rgbaHexString[3]),
			A: singleHexToInt(rgbaHexString[4], rgbaHexString[4]),
		}, nil
	}
	// "#RRGGBBAA"
	return color.NRGBA{
		R: singleHexToInt(rgbaHexString[1], rgbaHexString[2]),
		G: singleHexToInt(rgbaHexString[3], rgbaHexString[4]),
		B: singleHexToInt(rgbaHexString[5], rgbaHexString[6]),
		A: singleHexToInt(rgbaHexString[7], rgbaHexString[8]),
	}, nil
}

func singleHexToInt(high, low byte) uint8 {
	return hexToIntMap[high]*16 + hexToIntMap[low]
}

var hexToIntMap = map[byte]uint8{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'a': 10,
	'A': 10,
	'b': 11,
	'B': 11,
	'c': 12,
	'C': 12,
	'd': 13,
	'D': 13,
	'e': 14,
	'E': 14,
	'f': 15,
	'F': 15,
}

var rgbRE = regexp.MustCompile("^#(([0-9a-fA-F]{6})|([0-9a-fA-F]{3}))$")
var rgbaRE = regexp.MustCompile("^#(([0-9a-fA-F]{8})|([0-9a-fA-F]{4}))$")
