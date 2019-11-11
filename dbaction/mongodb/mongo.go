package main

import (
	"fmt"
	"github.com/globalsign/mgo"
)

const (
// USER string = "user"
// MSG  string = "msg"
)

var (
	session      *mgo.Session
	ip           = "127.0.0.1"
	port         = "27017"
	databaseName = "movies"
	maxPoolSize  = "10"
	poolLimit    = 10
)

func Session() *mgo.Session {
	if session == nil {
		var err error
		session, err = mgo.Dial(ip + ":" + port + "?maxPoolSize=" + maxPoolSize)
		session.SetPoolLimit(poolLimit)
		if err != nil {
			panic(err) // no, not really
		}
	}
	return session.Clone()
}

// 可以指定collection，database使用配置中的值
func MC(collection string, f func(*mgo.Collection) error) error {
	session := Session()
	defer func() {
		session.Close()
		// if err = recover(); err != nil {
		// Log("M", err)
		// }
	}()
	c := session.DB(databaseName).C(collection)
	// 关于return 和 defer 执行的优先级参看:https://studygolang.com/articles/4809
	return f(c)
}

// 可以指定database和collection
func MDC(dbName string, collection string, f func(*mgo.Collection) error) error {
	session := Session()
	defer func() {
		session.Close()
		if err := recover(); err != nil {
			// Log("M", err)
		}
	}()
	c := session.DB(dbName).C(collection)
	return f(c)
}

func main() {
	session := Session()
	defer func() {
		session.Close()
	}()
	c := session.DB("runoob").C("coll1")
	fmt.Println(c.FullName)
}
