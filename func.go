package main

import (
	"github.com/ArcherDing/EasyDict/models"
	"log"
	"strings"
)

/* wordType: 0x01 汉字
             0x02 假名
			 0x03 英文
*/
func TransWord(word string) (newWord string, wordType int) {
	chars := ""
	wordType = 2
	runes := []rune(word)

	for _, key := range runes {
		if key >= '\u4e00' && key <= '\u9fa5' {
			objs, count := models.GetMaps(string(key))
			if count > 0 {
				newWord += objs[0].Value
			} else {
				newWord += string(key)
			}
			wordType = 1
		} else if key >= '\u0040' && key <= '\u00ff' || key == '-' {
			chars += string(key)
			objs, count := models.GetMaps(chars)
			if count > 0 {
				newWord += objs[0].Value
				chars = ""
			}
		} else {
			newWord += string(key)
		}

	}

	if len(chars) > 0 {
		newWord = word
		wordType = 3
	}
	log.Println(newWord)
	return newWord, wordType
}

func GetDictsByKannji(value string) []models.Dict {
	words := make([]models.Dict, 0)
	for i := 0; i < len(value); i++ {
		words, count := models.GetDictByKannji(value[:len(value)-i], 20)
		if count > 0 {
			return words
		}
	}

	word := models.Dict{Word: value + "【" + JpToCh(value) + "】", Kannji: value, Kana: value}
	words = append(words, word)
	return words
}

func GetDictsByKana(value string) []models.Dict {
	words := make([]models.Dict, 0)
	if IsKatakana(value) {
		words, count := models.GetDictByKanaEqual(value)
		if count > 0 {
			return words
		}
		word := models.Dict{Word: value + "【" + JpToEn(value) + "】", Kannji: value, Kana: value}
		words = append(words, word)
		return words
	} else {
		for i := 0; i < len(value); i++ {
			words, count := models.GetDictByKana(value[:len(value)-i], 20)
			if count > 0 {
				return words
			}
		}
	}

	word := models.Dict{Word: value + "【" + JpToCh(value) + "】", Kannji: value, Kana: value}
	words = append(words, word)
	return words
}

func GetDictsByEnglish(value string) []models.Dict {
	words := make([]models.Dict, 0)
	words, count := models.GetDictByKannjiEqual(strings.ToLower(value))
	if count > 0 {
		return words
	}

	kana := EnToJp(value)
	words, count = models.GetDictByKanaEqual(strings.ToLower(kana))
	if count > 0 {
		return words
	}
	kannji := value
	meaning := JpToCh(kana)
	word := models.Dict{Word: kana + "【" + kannji + "】", Kannji: kannji, Kana: kana, Meaning: meaning}
	if meaning == kana {
		words = append(words, word)

	} else {
		models.AddDict(&word)
	}
	_words, _ := models.GetDictByKannji(strings.ToLower(value), 20)
	if meaning == kana {
		words = append(words, _words...)
	} else {
		words = _words
	}
	return words
}
