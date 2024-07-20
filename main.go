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
	PointSize: 10,
	Bold:      false,
	Italic:    false,
	Underline: false,
	StrikeOut: false,
}

var outputText, inputText, iv, key *walk.LineEdit
var inputTextLabel, outputTextLabel *walk.Label
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
		outputText.SetText(hex.EncodeToString(cipherText))
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
		outputText.SetText(string(plaintext))
	}
	actionEnabled.SetSatisfied(true)
}

func pushButtonAsync() {
	actionEnabled.SetSatisfied(false)
	go pushButtonClicked()
}

func radioButtonPressed() {
	actionButton.SetText(radioButtonData.Action)
	if radioButtonData.Action == "Encrypt" {
		inputTextLabel.SetText("Plain text: ")
		outputTextLabel.SetText("Cipher text: ")
	} else {
		inputTextLabel.SetText("Cipher text: ")
		outputTextLabel.SetText("Plain text: ")
	}
}

func main() {
	declarative.MustRegisterCondition("isEnabled", actionEnabled)
	actionEnabled.SetSatisfied(true)
	children := []declarative.Widget{
		declarative.Composite{
			Layout: declarative.Grid{
				Columns:   2,
				Alignment: declarative.AlignHCenterVCenter,
			},
			Children: []declarative.Widget{
				declarative.Label{Text: "Plain text: ", Font: usedFont, AssignTo: &inputTextLabel},
				declarative.LineEdit{AssignTo: &inputText, Font: usedFont},
				declarative.Label{Text: "IV: ", Font: usedFont},
				declarative.LineEdit{AssignTo: &iv, Font: usedFont},
				declarative.Label{Text: "Key: ", Font: usedFont},
				declarative.LineEdit{AssignTo: &key, Font: usedFont},
				declarative.RadioButtonGroupBox{
					ColumnSpan: 2,
					DataMember: "Action",
					Title:      "Action",
					Font:       usedFont,
					Layout:     declarative.HBox{},
					Buttons: []declarative.RadioButton{
						{
							AssignTo:  &encryptRadioButton,
							Font:      usedFont,
							Text:      "Encrypt",
							Value:     "Encrypt",
							OnClicked: radioButtonPressed,
						},
						{
							AssignTo:  &decryptRadioButton,
							Font:      usedFont,
							Text:      "Decrypt",
							Value:     "Decrypt",
							OnClicked: radioButtonPressed,
						},
					},
				},
				declarative.PushButton{
					ColumnSpan: 2,
					Text:       radioButtonData.Action,
					OnClicked:  pushButtonAsync,
					AssignTo:   &actionButton,
					Enabled:    declarative.Bind("isEnabled"),
					Font:       usedFont,
				},
				declarative.Label{Text: "Cipher text: ", Font: usedFont, AssignTo: &outputTextLabel},
				declarative.LineEdit{AssignTo: &outputText, ReadOnly: true, Font: usedFont},
			},
		},
	}

	declarative.MainWindow{
		Title:   "AES Encryption/Decryption",
		MinSize: declarative.Size{Width: 800, Height: 400},
		Size:    declarative.Size{Width: 800, Height: 400},
		DataBinder: declarative.DataBinder{
			DataSource: radioButtonData,
			AutoSubmit: true,
		},
		Layout:   declarative.VBox{},
		Children: children,
	}.Run()
}
