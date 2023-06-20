/**
* Author: Xiangyu Wu
* Date: 2023-06-19
* From: hyperledger/fabric/common/flogging/fabenc/color_test.go
 */

package enc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColorReset(t *testing.T) {
	require.Equal(t, ResetColor(), "\x1b[0m")

	t.Log(ResetColor() + "test color")
}

func TestNormalColors(t *testing.T) {
	require.Equal(t, ColorBlack.Normal(), "\x1b[30m")
	require.Equal(t, ColorRed.Normal(), "\x1b[31m")
	require.Equal(t, ColorGreen.Normal(), "\x1b[32m")
	require.Equal(t, ColorYellow.Normal(), "\x1b[33m")
	require.Equal(t, ColorBlue.Normal(), "\x1b[34m")
	require.Equal(t, ColorMagenta.Normal(), "\x1b[35m")
	require.Equal(t, ColorCyan.Normal(), "\x1b[36m")
	require.Equal(t, ColorWhite.Normal(), "\x1b[37m")

	t.Log(ColorBlack.Normal() + "black" + ResetColor())
	t.Log(ColorRed.Normal() + "red" + ResetColor())
	t.Log(ColorGreen.Normal() + "green" + ResetColor())
	t.Log(ColorYellow.Normal() + "yellow" + ResetColor())
	t.Log(ColorBlue.Normal() + "blue" + ResetColor())
	t.Log(ColorMagenta.Normal() + "magenta" + ResetColor())
	t.Log(ColorCyan.Normal() + "cyan" + ResetColor())
	t.Log(ColorWhite.Normal() + "white" + ResetColor())
}

func TestBoldColors(t *testing.T) {
	require.Equal(t, ColorBlack.Bold(), "\x1b[30;1m")
	require.Equal(t, ColorRed.Bold(), "\x1b[31;1m")
	require.Equal(t, ColorGreen.Bold(), "\x1b[32;1m")
	require.Equal(t, ColorYellow.Bold(), "\x1b[33;1m")
	require.Equal(t, ColorBlue.Bold(), "\x1b[34;1m")
	require.Equal(t, ColorMagenta.Bold(), "\x1b[35;1m")
	require.Equal(t, ColorCyan.Bold(), "\x1b[36;1m")
	require.Equal(t, ColorWhite.Bold(), "\x1b[37;1m")

	t.Log(ColorBlack.Bold() + "bold black" + ResetColor())
	t.Log(ColorRed.Bold() + "bold red" + ResetColor())
	t.Log(ColorGreen.Bold() + "bold green" + ResetColor())
	t.Log(ColorYellow.Bold() + "bold yellow" + ResetColor())
	t.Log(ColorBlue.Bold() + "bold blue" + ResetColor())
	t.Log(ColorMagenta.Bold() + "bold magenta" + ResetColor())
	t.Log(ColorCyan.Bold() + "bold cyan" + ResetColor())
	t.Log(ColorWhite.Bold() + "bold white" + ResetColor())
}