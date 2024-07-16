package main

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
	"github.com/meglicfran/GolangAesEncrptDecrpt/utils"
)

func main() {
	var inTE, outTE *walk.TextEdit
	cipherText, err := utils.AesEncrypt("TestTestTest12345", "aesEncryptionKey", []byte("1234567890123456"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hex.EncodeToString(cipherText))
	declarative.MainWindow{
		Title:   "SCREAMO",
		MinSize: declarative.Size{Width: 600, Height: 400},
		Layout:  declarative.VBox{},
		Children: []declarative.Widget{
			declarative.HSplitter{
				Children: []declarative.Widget{
					declarative.TextEdit{AssignTo: &inTE},
					declarative.TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			declarative.PushButton{
				Text: "SCREAM",
				OnClicked: func() {
					outTE.SetText(strings.ToUpper(inTE.Text()))
				},
			},
		},
	}.Run()
}
