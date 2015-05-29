package main

import (
	"github.com/ArcherDing/EasyDict/models"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
	"log"
	"strings"
)

const (
	IDI_ICON = 11
)

type MyMainWindow struct {
	*walk.MainWindow
	lbWords   *walk.ListBox
	cbKeys    *walk.ComboBox
	teDesc    *walk.TextEdit
	keyList   []string
	wordModel *WordModel
}

type WordModel struct {
	walk.ListModelBase
	items []models.Dict
}

func (m *WordModel) ItemCount() int {
	return len(m.items)
}

func (m *WordModel) Value(index int) interface{} {
	return m.items[index].Word
}

func (this *MyMainWindow) ToggleStyle(b bool, style int) {
	originalStyle := int(win.GetWindowLongPtr(this.Handle(), win.GWL_STYLE))
	if originalStyle != 0 {
		if b {
			originalStyle |= style
		} else {
			originalStyle ^= style
		}
		win.SetWindowLongPtr(this.Handle(), win.GWL_STYLE, uintptr(originalStyle))
	}
}

func (this *MyMainWindow) SetMaxButtonEnable(b bool) {
	this.ToggleStyle(b, win.WS_MAXIMIZEBOX)
}

func (this *MyMainWindow) SetMinButtonEnable(b bool) {
	this.ToggleStyle(b, win.WS_MINIMIZEBOX)
}

func (this *MyMainWindow) SetSizableEnable(b bool) {
	this.ToggleStyle(b, win.WS_THICKFRAME)
}

func (this *MyMainWindow) Center() {
	sWidth := win.GetSystemMetrics(win.SM_CXFULLSCREEN)
	sHeight := win.GetSystemMetrics(win.SM_CYFULLSCREEN)
	if sWidth != 0 && sHeight != 0 {
		size := this.Size()
		this.SetX(int(sWidth/2) - (size.Width / 2))
		this.SetY(int(sHeight/2) - (size.Height / 2))
	}
}

func (this *MyMainWindow) cbKeys_OnKeyPress(key walk.Key) {
	if key == win.VK_RETURN {
		text := this.cbKeys.Text()
		if len(text) <= 0 {
			return
		}

		newWord, wordType := TransWord(text)
		words := make([]models.Dict, 0)
		if wordType == 1 {
			words = GetDictsByKannji(newWord)
		} else if wordType == 2 {
			words = GetDictsByKana(newWord)
		} else {
			words = GetDictsByEnglish(newWord)
			newWord = words[0].Kana
		}
		this.wordModel = &WordModel{items: words}
		this.lbWords.SetModel(this.wordModel)
		this.lbWords.SetCurrentIndex(0)
		words[0].Meaning = strings.Replace(words[0].Meaning, "\n", "\r\n", -1)
		this.teDesc.SetText(words[0].Word + "\r\n" + words[0].Meaning)

		for index, value := range this.keyList {
			if value == newWord {
				this.cbKeys.SetCurrentIndex(index)
				return
			}
		}

		this.keyList = append(this.keyList, newWord)
		this.cbKeys.SetModel(this.keyList)
		this.cbKeys.SetCurrentIndex(len(this.keyList) - 1)
	}

	if key == win.VK_F1 {
		text := this.cbKeys.Text()
		if len(text) <= 0 {
			return
		}
		words := GetDictsByEnglish(text)
		this.wordModel = &WordModel{items: words}
		this.lbWords.SetModel(this.wordModel)
		this.lbWords.SetCurrentIndex(0)
		words[0].Meaning = strings.Replace(words[0].Meaning, "\n", "\r\n", -1)
		this.teDesc.SetText(words[0].Word + "\r\n" + words[0].Meaning)

		for index, value := range this.keyList {
			if value == words[0].Kana {
				this.cbKeys.SetCurrentIndex(index)
				return
			}
		}

		this.keyList = append(this.keyList, words[0].Kana)
		this.cbKeys.SetModel(this.keyList)
		this.cbKeys.SetCurrentIndex(len(this.keyList) - 1)
	}
}

func (this *MyMainWindow) teDesc_OnKeyPress(key walk.Key) {
	if key == win.VK_INSERT {
		text := this.teDesc.Text()
		if len(text) <= 0 {
			return
		}
		word := models.Dict{}
		start := strings.Index(text, "【")
		end := strings.Index(text, "】")
		word.Word = text[:end+3]
		word.Kannji = text[start+3 : end]
		word.Kana = text[:start]
		word.Meaning = text[end+5:]
		log.Println(word)
		models.AddDict(&word)
		words := GetDictsByKana(word.Kana)
		this.wordModel = &WordModel{items: words}
		this.lbWords.SetModel(this.wordModel)
		this.lbWords.SetCurrentIndex(0)
		words[0].Meaning = strings.Replace(words[0].Meaning, "\n", "\r\n", -1)
		this.teDesc.SetText(words[0].Word + "\r\n" + words[0].Meaning)
	}

	if key == win.VK_HOME {
		text := this.teDesc.Text()
		if len(text) <= 0 {
			return
		}
		index := this.lbWords.CurrentIndex()
		id := this.wordModel.items[index].Id
		start := strings.Index(text, "【")
		end := strings.Index(text, "】")
		word := models.Dict{}
		word.Id = id
		word.Word = text[:end+3]
		word.Kannji = text[start+3 : end]
		word.Kana = text[:start]
		word.Meaning = text[end+5:]
		log.Println(word)
		models.UpdateDict(&word)
		words := GetDictsByKana(word.Kana)
		this.wordModel = &WordModel{items: words}
		this.lbWords.SetModel(this.wordModel)
		this.lbWords.SetCurrentIndex(0)
		words[0].Meaning = strings.Replace(words[0].Meaning, "\n", "\r\n", -1)
		this.teDesc.SetText(words[0].Word + "\r\n" + words[0].Meaning)
	}
}

func (this *MyMainWindow) lbWords_OnCurrentIndexChanged() {
	index := this.lbWords.CurrentIndex()
	if index >= 0 {
		word := models.GetDictById(this.wordModel.items[index].Id)
		word.Meaning = strings.Replace(word.Meaning, "\n", "\r\n", -1)
		this.teDesc.SetText(word.Word + "\r\n" + word.Meaning)
	}
}

func InitUI() {
	mw := new(MyMainWindow)
	if err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "速查日语词典 ArcherDing@163.com",
		MinSize:  Size{700, 400},
		Layout:   VBox{SpacingZero: true},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 7, MarginsZero: true},
				Children: []Widget{
					ListBox{
						AssignTo:   &mw.lbWords,
						ColumnSpan: 2,
						Font:       Font{Family: "微软雅黑", PointSize: 10},
						OnCurrentIndexChanged: mw.lbWords_OnCurrentIndexChanged,
					},
					TextEdit{
						AssignTo:   &mw.teDesc,
						ColumnSpan: 5,
						MinSize:    Size{100, 50},
						Font:       Font{Family: "微软雅黑", PointSize: 10},
						OnKeyPress: mw.teDesc_OnKeyPress,
					},
					ComboBox{
						AssignTo:   &mw.cbKeys,
						ColumnSpan: 2,
						Editable:   true,
						Font:       Font{Family: "微软雅黑", PointSize: 9},
						OnKeyPress: mw.cbKeys_OnKeyPress,
					},
					Label{
						ColumnSpan: 5,
						Text:       "速查日语词典  By：丁琪 QQ：739367060 [日语小学馆词典版] F1:查英语  Enter：查日语",
					},
				},
			},
		},
	}).Create(); err != nil {
		log.Fatal(err)
	}

	if ico, err := walk.NewIconFromResourceId(IDI_ICON); err == nil {
		mw.SetIcon(ico)
	}

	mw.SetMaxButtonEnable(false)
	mw.SetSizableEnable(false)
	mw.Center()

	mw.Show()
	mw.Run()
}
