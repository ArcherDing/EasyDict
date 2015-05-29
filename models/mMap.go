package models

import "github.com/astaxie/beego/orm"

type Map struct {
	Id    int64
	Key   string `orm:"size(10)"`
	Value string `orm:"size(10)"`
}

func AddMap(obj *Map) (int64, error) {
	o := orm.NewOrm()

	id, err := o.Insert(obj)
	if err != nil {
		return 0, err
	}

	return id, err
}

func GetMaps(key string) ([]Map, int64) {
	o := orm.NewOrm()
	objs := make([]Map, 0)
	count, _ := o.Raw("SELECT * FROM map WHERE key = ?", key).QueryRows(&objs)

	return objs, count
}

func DeleteMap(id int64) {

}

func UpdateMap(id int64, value string) {

}
