package controllers

import (
	"gopkg.in/mgo.v2"
	"github.com/vineetdaniel/taskmanager/common"
)

//Struct used for maintiaing HTTP Request Context

type Context struct {
	MongoSession *mgo.Session
}

//close mgo session
func (c *Context) Close() {
	c.MongoSession.Close()
}

Return mgo collection for the given name

func (c *Context) DbCollection(name string) *mgo.Collection {
	return c.MongoSession.DB(common.AppConfig.Database).C(name)
}

//Create a new Context object for each HTTP request

func NewContext() *Context {
	session := common.GetSession().Copy()
	context := &Context{
		MongoSession: session,
	}
	return context
}

