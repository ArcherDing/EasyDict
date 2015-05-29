package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func EnToJp(text string) string {
	/*
		proxy := func(_ *http.Request) (*url.URL, error) {
			return url.Parse("http://172.28.25.121:8080")
		}
		transport := &http.Transport{Proxy: proxy}
		client := &http.Client{Transport: transport}
	*/
	client := &http.Client{}
	response, _ := client.PostForm("http://fanyi.baidu.com/multitransapi",
		url.Values{"from": {"en"}, "to": {"jp"}, "query": {text}, "raw_trans": {""}, "count": {"5"}})
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	if strings.Contains(string(body), "cands") == false {
		return text
	}

	var result map[string]interface{}
	json.Unmarshal(body, &result)
	var data map[string]interface{}
	data = result["data"].(map[string]interface{})
	cands := data["cands"].([]interface{})
	for _, value := range cands {
		if IsKatakana(value.(string)) {
			return value.(string)
		}
	}
	return text
}

func JpToEn(text string) string {

	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://172.28.25.121:8080")
	}
	transport := &http.Transport{Proxy: proxy}
	client := &http.Client{Transport: transport}

	//client := &http.Client{}
	response, _ := client.PostForm("http://fanyi.baidu.com/multitransapi",
		url.Values{"from": {"jp"}, "to": {"en"}, "query": {text}, "raw_trans": {""}, "count": {"5"}})
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	if strings.Contains(string(body), "cands") == false {
		log.Println(string(body))
		return text
	}

	var result map[string]interface{}
	json.Unmarshal(body, &result)
	var data map[string]interface{}
	data = result["data"].(map[string]interface{})
	cands := data["cands"].([]interface{})
	log.Println(cands)
	return cands[1].(string)
}

func JpToCh(text string) string {

	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://172.28.25.121:8080")
	}
	transport := &http.Transport{Proxy: proxy}
	client := &http.Client{Transport: transport}

	//client := &http.Client{}
	response, _ := client.PostForm("http://fanyi.baidu.com/multitransapi",
		url.Values{"from": {"jp"}, "to": {"zh"}, "query": {text}, "raw_trans": {""}, "count": {"5"}})
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	if strings.Contains(string(body), "cands") == false {
		log.Println(string(body))
		return text
	}

	var result map[string]interface{}
	json.Unmarshal(body, &result)
	var data map[string]interface{}
	data = result["data"].(map[string]interface{})
	cands := data["cands"].([]interface{})
	var kannji string
	for _, value := range cands {
		if text == value.(string) {
			return text
		}
		kannji += value.(string) + ","
	}
	log.Println(kannji)
	return kannji
}

func IsKatakana(value string) bool {
	runes := []rune(value)
	for _, key := range runes {
		if key < '\u30A0' || key > '\u30FF' {
			return false
		}
	}
	return true
}
