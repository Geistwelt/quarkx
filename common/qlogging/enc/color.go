/**
* Author: Xiangyu Wu
* Date: 2023-06-19
* From: hyperledger/fabric/common/flogging/fabenc/color.go
 */

package enc

import "fmt"

type Color uint8

const ColorNone Color = 0

const (
	ColorBlack   Color = 30
	ColorRed     Color = 31
	ColorGreen   Color = 32
	ColorYellow  Color = 33
	ColorBlue    Color = 34
	ColorMagenta Color = 35 // 品红
	ColorCyan    Color = 36 // 青色
	ColorWhite   Color = 37
)

func (c Color) Normal() string {
	return fmt.Sprintf("\x1b[%dm", c)
}

func (c Color) Bold() string {
	if c == ColorNone {
		return c.Normal()
	}
	return fmt.Sprintf("\x1b[%d;1m", c)
}

func ResetColor() string {
	return ColorNone.Normal()
}
