package infra

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type SqlHandler struct {
	Conn *gorm.DB
}

func NewSqlHandler() SqlHandler {
	// conn, err := gorm.Open("mysql", "root:root@tcp(mysql)/sample?charset=utf8&parseTime=True&loc=Local")
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(echo-db)"
	DBNAME := "app"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	conn, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error)
	}
	sqlHandler := SqlHandler{}
	sqlHandler.Conn = conn
	return sqlHandler
}

func (handler *SqlHandler) Find(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.Find(out, where...)
}

func (handler *SqlHandler) Exec(sql string, values ...interface{}) *gorm.DB {
	return handler.Conn.Exec(sql, values...)
}

func (handler *SqlHandler) First(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.First(out, where...)
}

func (handler *SqlHandler) Raw(sql string, values ...interface{}) *gorm.DB {
	return handler.Conn.Raw(sql, values...)
}

func (handler *SqlHandler) Create(value interface{}) *gorm.DB {
	return handler.Conn.Create(value)
}

func (handler *SqlHandler) Save(value interface{}) *gorm.DB {
	return handler.Conn.Save(value)
}

func (handler *SqlHandler) Delete(value interface{}) *gorm.DB {
	return handler.Conn.Delete(value)
}

func (handler *SqlHandler) Where(query interface{}, args ...interface{}) *gorm.DB {
	return handler.Conn.Where(query, args...)
}

func (handler *SqlHandler) Debug() *gorm.DB {
	return handler.Conn.Debug()
}

func (handler *SqlHandler) Do() *gorm.DB {
	return handler.Conn
}
