/**
 * Created by 93201 on 2017/5/10.
 */
package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/x-croz/config"
	"github.com/x-croz/log"
)

var mdb *gorm.DB

func init() {
	if mdb != nil {
		return
	}
	var err error
	mdb, err = gorm.Open("mysql", config.String("mysql.conn"))
	if err != nil {
		log.Panic(err)
	}
	monitor = newDBTaskMonitor()
	monitor.addTask((&BitTorrent{}).TableName(), &btSaveTask{})
	monitor.init()
	go monitor.start()
}

func Close() {
	mdb.Close()
}
