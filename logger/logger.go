package logger

import (
	"fmt"
)

type Color string

const (
	ColorDefault Color = "\u001b[0m"
	ColorRed     Color = "\u001b[31m"
	ColorGreen   Color = "\u001b[32m"
	ColorYellow  Color = "\u001b[33m"
	ColorBlue    Color = "\u001b[34m"
	ColorReset   Color = "\u001b[0m"
)

func Log(color Color, msg string) {
	fmt.Print(string(color)+msg, string(ColorReset))
}

func Logln(color Color, msg string) {
	fmt.Println(string(color)+msg, string(ColorReset))
}
