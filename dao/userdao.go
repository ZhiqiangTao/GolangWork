package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

var db *sql.DB

func init() {
	dsn := `paywriter:RxXSM7IA2iMp7LwP@tcp(10.10.8.97:3307)/test`
	_db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic("数据库连接失败" + err.Error())
	}
	_db.SetMaxOpenConns(1000)
	db = _db
	fmt.Println("mysql sb inited.")
}

func QueryPersonById(id int64) ([]User, error) {
	res := make([]User, 0)
	sqlStr := "select id, name, age from user where id>?"
	var user User
	err := db.QueryRow(sqlStr, id).Scan(&user.Id, &user.Name, &user.Age)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			/*
				具体问题具体分析，sql.ErrNoRows只是没有数据了，业务层面不应该因为没有数据而“报错”，应该直接返回已经填充好的返回结果
				如果返回sql.ErrNoRows, 那么调用层就会依赖sql包，增加暴露的表面积
			*/
			return nil, nil
		} else {
			return nil, errors.Wrap(err, "查询错误巴拉巴拉~~~")
		}
	}

	res = append(res, user)
	return res, nil
}
