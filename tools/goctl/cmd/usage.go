package cmd

import (
	"github.com/gookit/color"
)

var colorRender = []func(v any) string{
	func(v any) string {
		return color.LightRed.Render(v)
	},
	func(v any) string {
		return color.LightGreen.Render(v)
	},
	func(v any) string {
		return color.LightYellow.Render(v)
	},
	func(v any) string {
		return color.LightBlue.Render(v)
	},
	func(v any) string {
		return color.LightMagenta.Render(v)
	},
	func(v any) string {
		return color.LightCyan.Render(v)
	},
}

func blue(s string) string { _ = "STUB: not implemented"; return "" }

func green(s string) string { _ = "STUB: not implemented"; return "" }

func rainbow(s string) string { _ = "STUB: not implemented"; return "" }

func rpadx(s string, padding int) string { _ = "STUB: not implemented"; return "" }
