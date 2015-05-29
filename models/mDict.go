package models

import "github.com/astaxie/beego/orm"
import "fmt"

type Dict struct {
	Id      int64
	Word    string `orm:"size(40)"`
	Kannji  string `orm:"size(20)"`
	Kana    string `orm:"size(20)"`
	Meaning string `orm:"size(300)"`
}

func AddDict(obj *Dict) (int64, error) {
	o := orm.NewOrm()

	id, err := o.Insert(obj)
	if err != nil {
		return 0, err
	}

	return id, err
}

func UpdateDict(obj *Dict) (int64, error) {
	o := orm.NewOrm()
	_obj := make(orm.Params)
	_obj["Word"] = obj.Word
	_obj["Kannji"] = obj.Kannji
	_obj["Kana"] = obj.Kana
	_obj["Meaning"] = obj.Meaning

	num, err := o.QueryTable("dict").Filter("Id", obj.Id).Update(_obj)
	return num, err
}

func GetDict(word string, size int64) ([]Dict, int64) {
	o := orm.NewOrm()
	objs := make([]Dict, 0)
	count, err := o.Raw("SELECT * FROM dict WHERE kannji LIKE ? LIMIT ?", word+"%", size).QueryRows(&objs)
	if err != nil {
		fmt.Println(err)
	}

	return objs, count
}

func GetDictByKannji(kannji string, size int64) ([]Dict, int64) {
	o := orm.NewOrm()
	objs := make([]Dict, 0)
	count, err := o.Raw("SELECT id,word,kana,meaning FROM dict WHERE kannji LIKE ? LIMIT ?", kannji+"%", size).QueryRows(&objs)
	if err != nil {
		fmt.Println(err)
	}

	return objs, count
}

func GetDictByKannjiEqual(kannji string) ([]Dict, int64) {
	o := orm.NewOrm()
	objs := make([]Dict, 0)
	count, err := o.Raw("SELECT id,word,kana,meaning FROM dict WHERE kannji = ?", kannji).QueryRows(&objs)
	if err != nil {
		fmt.Println(err)
	}

	return objs, count
}

func GetDictByKana(kana string, size int64) ([]Dict, int64) {
	o := orm.NewOrm()
	objs := make([]Dict, 0)
	count, err := o.Raw("SELECT id,word,kana,meaning FROM dict WHERE kana LIKE ? LIMIT ?", kana+"%", size).QueryRows(&objs)
	if err != nil {
		fmt.Println(err)
	}

	return objs, count
}

func GetDictByKanaEqual(kana string) ([]Dict, int64) {
	o := orm.NewOrm()
	objs := make([]Dict, 0)
	count, err := o.Raw("SELECT id,word,kana,meaning FROM dict WHERE kana = ?", kana).QueryRows(&objs)
	if err != nil {
		fmt.Println(err)
	}

	return objs, count
}

func GetDictById(id int64) Dict {
	o := orm.NewOrm()
	obj := Dict{}
	err := o.Raw("SELECT id,word,kana,meaning FROM dict WHERE id= ?", id).QueryRow(&obj)
	if err != nil {
		fmt.Println(err)
	}

	return obj
}
