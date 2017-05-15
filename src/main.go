package main

import (
	"crawler"
	"db"
	"github.com/x-croz/log"
	"github.com/x-croz/tpl"
	_ "net/http/pprof"
	"os"
	"time"
	"utils"
)

func init() {
	log.SetOutput(logg)
	tpl.SetViewsPath("tpl")
}

func main() {
	go utils.SavePID("app.pid")
	//go utils.HoldSignal(func(s os.Signal) {
	//	log.Info("application exited with signal: ", s)
	//	os.Exit(1)
	//})
	crawler.NewDhtCrawler().SetHandler(func(torrent *crawler.BitTorrent) {
		bt := &db.BitTorrent{}
		db.DBTaskMonitor().Send(bt.TableName(), bt.From(torrent))
	}).Run()
}
var logg = newLogger()
type logger struct {
	f    *os.File
	name string
}

func newLogger() *logger {
	l := &logger{}
	l.name = time.Now().Format("2006_01_02.log")
	l.f, _ = os.Create(l.name)
	return l
}

func (l *logger) Write(p []byte) (int, error) {
	name := time.Now().Format("2006_01_02.log")
	if name != l.name && l.f != nil {
		l.f.Close()
		l.f, _ = os.Create(name)
		l.name = name
	}
	n, _ := l.f.Seek(0, os.SEEK_END)
	return l.f.WriteAt(p, n)
}

func (l *logger) Close() {
	if l.f != nil {
		l.f.Close()
	}
}
