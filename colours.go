package colours

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// --- Colour and Style Structs --- //

type Colour string

type Style struct {
	colour    Colour
	bold      bool
	underline bool
	italic    bool
}

func (c Colour) Bold() *Style {
	return &Style{colour: c, bold: true}
}

func (c Colour) Underline() *Style {
	return &Style{colour: c, underline: true}
}

func (c Colour) Italic() *Style {
	return &Style{colour: c, italic: true}
}

func (c Colour) str() string {
	return string(c)
}

func (s *Style) Bold() *Style {
	newStyle := *s
	newStyle.bold = true
	return &newStyle
}

func (s *Style) Underline() *Style {
	newStyle := *s
	newStyle.underline = true
	return &newStyle
}

func (s *Style) Italic() *Style {
	newStyle := *s
	newStyle.italic = true
	return &newStyle
}

// --- Apply Style to Text --- //

func (c Colour) Apply(text string) string {
	var builder strings.Builder

	builder.WriteString(c.str())
	builder.WriteString(text)
	builder.WriteString(Reset.str())

	return builder.String()
}

func (s *Style) Apply(text string) string {
	var builder strings.Builder

	if s.bold {
		builder.WriteString(Bold)
	}
	if s.underline {
		builder.WriteString(Underline)
	}
	if s.italic {
		builder.WriteString(Italic)
	}

	builder.WriteString(s.colour.str())
	builder.WriteString(text)
	builder.WriteString(Reset.str())

	return builder.String()
}

func AddColour(text string, colour Colour) string {
	return colour.Apply(text)
}

// --- Colour & Style Definitions --- //

const (
	Red      Colour = "\033[31m"
	Green    Colour = "\033[32m"
	Yellow   Colour = "\033[33m"
	Blue     Colour = "\033[34m"
	Magenta  Colour = "\033[35m"
	Cyan     Colour = "\033[36m"
	White    Colour = "\033[37m"
	Orange   Colour = "\033[38;2;255;149;0m"
	DarkBlue Colour = "\033[38;2;0;0;139m"
	Reset    Colour = "\033[0m"
)

const (
	Bold      = "\033[1m"
	Underline = "\033[4m"
	Italic    = "\033[3m"
)

var colours = []Colour{Red, Green, Yellow, Blue, Magenta, Cyan, White, Orange, DarkBlue}

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

// --- Random Colours --- //

func randomDefaultColour() Colour {
	var r = random.Intn(len(colours))
	return colours[r]
}

func randomRGBColour() Colour {
	r := random.Intn(256)
	g := random.Intn(256)
	b := random.Intn(256)
	return Colour(fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b))
}

func ApplyRandomColour(text string) string {
	return randomDefaultColour().Apply(text)
}

func ApplyRandomRGBColour(text string) string {
	return randomRGBColour().Apply(text)
}
