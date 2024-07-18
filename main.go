package main

import (
	"encoding/hex"
	"fmt"

	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
	"github.com/meglicfran/GolangAesEncrptDecrpt/utils"
)

type rButton struct {
	Action string
}

var usedFont declarative.Font = declarative.Font{
	Family:    "",
	PointSize: 20,
	Bold:      false,
	Italic:    false,
	Underline: false,
	StrikeOut: false,
}

func main() {
	var _, outTE *walk.TextEdit
	var text, iv, key *walk.TextEdit
	buttonData := &rButton{"Encrypt"}
	cipherText, err := utils.AesEncrypt([]byte("TestTestTest12345"), []byte("aesEncryptionKey"), []byte("1234567890123456"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hex.EncodeToString(cipherText))
	plaintext, err := utils.AesDecrypt(cipherText, []byte("aesEncryptionKey"), []byte("1234567890123456"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(plaintext))

	Keyinput := declarative.TextEdit{AssignTo: &iv}
	children := []declarative.Widget{
		declarative.VSplitter{
			Children: []declarative.Widget{
				declarative.TextEdit{AssignTo: &text, Font: usedFont},
				declarative.TextEdit{AssignTo: &iv, Font: usedFont},
				declarative.TextEdit{AssignTo: &key, Font: usedFont},
				declarative.RadioButtonGroup{
					DataMember: "Action",
					Buttons: []declarative.RadioButton{
						{
							Font:  usedFont,
							Text:  "Encrypt",
							Value: "Encrypt",
							Name:  "Encrypt",
						},
						{
							Font:  usedFont,
							Text:  "Decrypt",
							Value: "Decrypt",
							Name:  "Decrypt",
						},
					},
				},
				declarative.TextEdit{AssignTo: &outTE, ReadOnly: true, Font: usedFont},
			},
		},
		declarative.PushButton{
			Text: buttonData.Action,
			OnClicked: func() {
				textBytes := []byte(text.Text())
				keyBytes := []byte(key.Text())
				ivBytes := []byte(iv.Text())
				fmt.Print(Keyinput.Font.Family)
				switch buttonData.Action {
				case "Encrypt":
					cipherText, err := utils.AesEncrypt(textBytes, keyBytes, ivBytes)
					if err != nil {
						outTE.SetText(err.Error())
						return
					}
					outTE.SetText("Cipher text:" + hex.EncodeToString(cipherText))
					return
				case "Decrypt":
					hexText, err := hex.DecodeString(text.Text())
					if err != nil {
						outTE.SetText(err.Error())
						return
					}
					plaintext, err := utils.AesDecrypt(hexText, keyBytes, ivBytes)
					if err != nil {
						outTE.SetText(err.Error())
						return
					}
					outTE.SetText("Plain text: " + string(plaintext) + " " + Keyinput.Font.Family)
					return
				}
				//outTE.SetText(strings.ToUpper(fmt.Sprintf("%v : ", buttonData.Action) + text.Text() + iv.Text() + key.Text()))
			},
		},
	}

	declarative.MainWindow{
		Title:   "AES Encryption/Decryption",
		MinSize: declarative.Size{Width: 600, Height: 400},
		DataBinder: declarative.DataBinder{
			DataSource: buttonData,
			AutoSubmit: true,
		},
		Layout:   declarative.VBox{},
		Children: children,
	}.Run()
}
