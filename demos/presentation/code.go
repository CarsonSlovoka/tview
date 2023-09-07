package main

import (
	"fmt"

	"github.com/rivo/tview"
)

// The width of the code window.
const codeWidth = 56

// Code returns a primitive which displays the given primitive (with the given
// size) on the left side and its source code on the right side.
// 左邊是展示的內容, 右邊是可以放原始代碼
func Code(p tview.Primitive, width, height int, code string) tview.Primitive {
	// Set up code view.
	codeView := tview.NewTextView().
		SetWrap(false).
		SetDynamicColors(true)
	codeView.SetBorderPadding(1, 1, 2, 0)
	fmt.Fprint(codeView, code)

	return tview.NewFlex().
		AddItem(Center(width, height, p), 0, 1, true). // 這是左側的內容
		AddItem(codeView, codeWidth, 1, false)         // 右側內容, 可以將原始代碼放入
}
