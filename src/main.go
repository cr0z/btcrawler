package main

import (
	"crawler"
	"db"
	"fmt"
	"github.com/x-croz/config"
	"github.com/x-croz/log"
	"github.com/x-croz/tpl"
	"net/http"
	_ "net/http/pprof"
	"os"
	"utils"
)

func init() {
	tpl.SetViewsPath("tpl")
}

func main() {
	go utils.SavePID("app.pid")
	go utils.HoldSignal(func(s os.Signal) {
		log.Info("application exited with signal: ", s)
		os.Exit(1)
	})
	spider := crawler.NewDhtCrawler()
	spider.SetHandler(func(torrent *crawler.BitTorrent) {
		bt := &db.BitTorrent{}
		db.Monitor().Send(bt.TableName(), bt.From(torrent))
	})
	go spider.Run()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var bq = db.BTQuery{}
		bts, err := bq.Query()
		if err != nil {
			log.Error(err)
		}
		render := tpl.NewRender()
		render.TplName = "index.tpl"
		render.Data["bts"] = bts
		render.Render(w)
	})
	addr := fmt.Sprintf("%s:%d", config.String("app.host"), config.Int64("app.port"))
	log.Info("http serve on:", addr)
	http.ListenAndServe(addr, nil)
}
