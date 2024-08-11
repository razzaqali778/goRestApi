package database

import (
	"time"

	log "github.com/sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"users.com/common"
	"users.com/models"
)

type MongoDB struct {
	MgDbSession  *mgo.Session
	DatabaseName string
}

func (db *MongoDB) init() error {
	db.DatabaseName = common.Config.MgDbName

	dialInfo := &mgo.DialInfo{
		Addrs:    []string{common.Config.MgAddrs},
		Timeout:  60 * time.Second,
		Database: db.DatabaseName,
		Username: common.Config.MgDbUsername,
		Password: common.Config.MgDbPassword,
	}
	var err error

	db.MgDbSession, err = mgo.DialWithInfo(dialInfo)

	if err != nil {
		log.Debug("can't connect to database", err)
		return err
	}

	return db.initData()
}

func (db *MongoDB) initData() error {
	var err error
	var count int

	sessionCopy := db.MgDbSession.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB(db.DatabaseName).C(common.ColUsers)
	count, err = collection.Find(bson.M{}).Count()

	if count < 1 {
		var user models.User
		user = models.User{bson.NewObjectId(), "admin", "admin"}
		err = collection.Insert(&user)
	}

	return err

}

func (db *MongoDB) Close() {
	if db.MgDbSession != nil {
		db.MgDbSession.Close()
	}
}
