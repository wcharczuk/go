package ansi

import (
	"testing"

	"go.charczuk.com/sdk/assert"
)

func TestColorApply(t *testing.T) {
	its := assert.New(t)

	escapedBlack := ColorBlack.Normal()
	its.Equal("\033[0;"+string(ColorBlack), escapedBlack)

	appliedBlack := ColorBlack.Apply("test")
	its.Equal(ColorBlack.Normal()+"test"+ColorReset, appliedBlack)
}

func TestColors(t *testing.T) {
	its := assert.New(t)

	its.Equal(ColorBlack.Apply("foo"), Black("foo"))
	its.Equal(ColorRed.Apply("foo"), Red("foo"))
	its.Equal(ColorGreen.Apply("foo"), Green("foo"))
	its.Equal(ColorYellow.Apply("foo"), Yellow("foo"))
	its.Equal(ColorBlue.Apply("foo"), Blue("foo"))
	its.Equal(ColorPurple.Apply("foo"), Purple("foo"))
	its.Equal(ColorCyan.Apply("foo"), Cyan("foo"))
	its.Equal(ColorWhite.Apply("foo"), White("foo"))
	its.Equal(ColorLightBlack.Apply("foo"), LightBlack("foo"))
	its.Equal(ColorLightRed.Apply("foo"), LightRed("foo"))
	its.Equal(ColorLightGreen.Apply("foo"), LightGreen("foo"))
	its.Equal(ColorLightYellow.Apply("foo"), LightYellow("foo"))
	its.Equal(ColorLightBlue.Apply("foo"), LightBlue("foo"))
	its.Equal(ColorLightPurple.Apply("foo"), LightPurple("foo"))
	its.Equal(ColorLightCyan.Apply("foo"), LightCyan("foo"))
	its.Equal(ColorLightWhite.Apply("foo"), LightWhite("foo"))

	its.Equal(ColorRed.Bold()+"foo"+ColorReset, Bold(ColorRed, "foo"))
	its.Equal(ColorRed.Underline()+"foo"+ColorReset, Underline(ColorRed, "foo"))
}
