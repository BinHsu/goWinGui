package main

import (
	"bytes"
	"fmt"
	"goWinGui/icon"
	"image/png"
	"strings"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	var inTE, outTE *walk.TextEdit
	/*pngFile, err := os.Open("icon/DeviceCenter.png")
	if err != nil {
		fmt.Println(err)
	}
	defer pngFile.Close()
	iconImage, err := png.Decode(pngFile)
	if err != nil {
		fmt.Println(err)
	}
	iconFromFile, err := walk.NewIconFromImage(iconImage)
	if err != nil && iconFromFile != nil {
		fmt.Println(err)
	}*/

	img, err := png.Decode(bytes.NewReader(icon.AboutPng))
	if err != nil {
		fmt.Println(err)
	}
	iconFromBytes, err := walk.NewIconFromImage(img)
	if err != nil {
		fmt.Println(err)
	}

	MainWindow{
		Icon:    iconFromBytes,
		Title:   "SCREAMO",
		MinSize: Size{600, 400},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE},
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			PushButton{
				Text: "SCREAM",
				OnClicked: func() {
					outTE.SetText(strings.ToUpper(inTE.Text()))
				},
			},
		},
	}.Run()
}
