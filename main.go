package main

import (
	"encoding/hex"

	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
	"github.com/meglicfran/GolangAesEncrptDecrpt/utils"
)

type radioButton struct {
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

var outputText, inputText, iv, key *walk.TextEdit
var actionButton *walk.PushButton
var encryptRadioButton, decryptRadioButton *walk.RadioButton
var radioButtonData *radioButton = &radioButton{"Encrypt"}
var actionEnabled = walk.NewMutableCondition()

func pushButtonClicked() {
	textBytes := []byte(inputText.Text())
	keyBytes := []byte(key.Text())
	ivBytes := []byte(iv.Text())
	switch radioButtonData.Action {
	case "Encrypt":
		cipherText, err := utils.AesEncrypt(textBytes, keyBytes, ivBytes)
		if err != nil {
			outputText.SetText(err.Error())
			break
		}
		outputText.SetText("Cipher text:" + hex.EncodeToString(cipherText))
	case "Decrypt":
		hexText, err := hex.DecodeString(inputText.Text())
		if err != nil {
			outputText.SetText(err.Error())
			break
		}
		plaintext, err := utils.AesDecrypt(hexText, keyBytes, ivBytes)
		if err != nil {
			outputText.SetText(err.Error())
			break
		}
		outputText.SetText("Plain text: " + string(plaintext))
	}
	actionEnabled.SetSatisfied(true)
}

func pushButtonAsync() {
	actionEnabled.SetSatisfied(false)
	go pushButtonClicked()
}

func updateActionText() {
	actionButton.SetText(radioButtonData.Action)
}

func main() {
	declarative.MustRegisterCondition("isEnabled", actionEnabled)
	actionEnabled.SetSatisfied(true)
	children := []declarative.Widget{
		declarative.VSplitter{
			Children: []declarative.Widget{
				declarative.TextEdit{AssignTo: &inputText, Font: usedFont},
				declarative.TextEdit{AssignTo: &iv, Font: usedFont},
				declarative.TextEdit{AssignTo: &key, Font: usedFont},
				declarative.RadioButtonGroup{
					DataMember: "Action",
					Buttons: []declarative.RadioButton{
						{
							AssignTo:  &encryptRadioButton,
							Font:      usedFont,
							Text:      "Encrypt",
							Value:     "Encrypt",
							Name:      "Encrypt",
							OnClicked: updateActionText,
						},
						{
							AssignTo:  &decryptRadioButton,
							Font:      usedFont,
							Text:      "Decrypt",
							Value:     "Decrypt",
							Name:      "Decrypt",
							OnClicked: updateActionText,
						},
					},
				},
				declarative.TextEdit{AssignTo: &outputText, ReadOnly: true, Font: usedFont},
			},
		},
		declarative.PushButton{
			Text:      radioButtonData.Action,
			OnClicked: pushButtonAsync,
			AssignTo:  &actionButton,
			Enabled:   declarative.Bind("isEnabled"),
		},
	}

	declarative.MainWindow{
		Title:   "AES Encryption/Decryption",
		MinSize: declarative.Size{Width: 600, Height: 400},
		DataBinder: declarative.DataBinder{
			DataSource: radioButtonData,
			AutoSubmit: true,
		},
		Layout:   declarative.VBox{},
		Children: children,
	}.Run()
}
