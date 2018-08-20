
package models

import (
	"database/sql"
	. "excel/db"
	"time"
	"fmt"
)

const TimeFormat = "2006-01-02 15:04:05"

type JsonTime time.Time

// 实现它的json序列化方法
func (this JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

type Tag struct {
	Id     int64 `json:"id"`
	Name  string `json:"name"`
	Frequency int `json:"frequency"`
	AddTime JsonTime `json:"add_time"`
	UpdateTime JsonTime `json:"update_time"`
}


func initConnPool() (*sql.DB,error) {
	Db := &MySQLClient{Host:"localhost",User:"root",Pwd:"123456",DB:"blog2",Port:3306,MaxOpen:300,MaxIdle:200}
	Db.Init()
	return Db.Pool,nil
}

//insert
func InsertTag(name,frequency string) (int64, error) {
	db,err :=initConnPool()
	stmt, err := db.Prepare("INSERT INTO tbl_tag(name,frequency) VALUES(?,?)")
	defer stmt.Close()
	if err != nil {
		return -1, err
	}

	res, err := stmt.Exec(name,frequency)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

//getall
func QueryAll() ([]Tag, error) {
	db,err :=initConnPool()
	rows, err := db.Query("SELECT * FROM tbl_tag limit 1,30")
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var tags []Tag
	for rows.Next() {
		var id int64
		var name string
		var frequency int
		var addTime JsonTime  //为了"2018-06-24T12:02:28+08:00" 转换为 可读的形式
		var updateTime JsonTime

		//updateTime.Format("2006-01-02 15:04:05")
		err = rows.Scan(&id, &name, &frequency,&addTime,&updateTime)
		if err != nil {
			return nil, err
		}
		tag:= Tag{id, name, frequency,addTime,updateTime}
		//updateTime.Format("2006-01-02 15:04:05")
		//addTime = addTime.Format("2006-01-02 15:04:05")
		tags = append(tags, tag)
	}
	return tags, nil
}


//修改
func FinishTag(tagId int64, name string) (int64, error) {
	db,err :=initConnPool()
	stmt, err := db.Prepare("UPDATE tbl_tag SET name=? WHERE id=?")
	defer stmt.Close()
	if err != nil {
		return 0, nil
	}

	res, err := stmt.Exec(name, tagId)
	if err != nil {
		return 0, nil
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return 0, nil
	}
	return affect, nil
}


//删除
func DeleteTag(todoId int64) (int64, error) {
	db,err := initConnPool()
	stmt, err := db.Prepare("DELETE FROM tbl_tag WHERE id=?")
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(todoId)
	if err != nil {
		return 0, nil
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return 0, nil
	}
	return affect, nil
}


//获取单个
func GetTagName(todoId int64) (string,int, error) {

	db,_ := initConnPool()
	// 只查询一行数据
	var name string
	var frequency int
	err := db.QueryRow("SELECT name ,frequency FROM tbl_tag WHERE id=?", todoId).Scan(&name,&frequency)
	if err != nil {
		return "",0, err
	}
	return name,frequency, nil
}
