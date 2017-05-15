/**
 * Created by 93201 on 2017/5/10.
 */
package db

import (
	"github.com/x-croz/log"
	"time"
)

type btSaveTask struct {
	btCache []*BitTorrent
	dChan   chan *BitTorrent
	sChan   chan int
}

func (t *btSaveTask) init() {
	t.btCache = []*BitTorrent{}
	t.dChan = make(chan *BitTorrent)
	t.sChan = make(chan int)
}

func (t *btSaveTask) start() {
	tick := time.Tick(time.Minute * 30)
	for {
		select {
		case d := <-t.dChan:
			t.btCache = append(t.btCache, d)
		case <-tick:
			t.save()
		case <-t.sChan: //stop
			t.save()
			return
		}
	}
}

func (t *btSaveTask) stop() {
	t.sChan <- -1
}

func (t *btSaveTask) onCall(d interface{}) {
	if bt, ok := d.(*BitTorrent); ok {
		t.dChan <- bt
	}
}

var savedCount = 0

func (t *btSaveTask) save() {
	if len(t.btCache) == 0 || savedCount > 1000000 {
		return
	}
	temp := t.btCache
	t.btCache = []*BitTorrent{}
	tx := mdb.Begin()
	for _, v := range temp {
		tx.Create(v)
	}
	err := tx.Commit().Error
	if err == nil {
		log.Info("save bt success, save count:", len(temp))
		savedCount += len(temp)
	} else {
		log.Warning("save bt faild: ", err)
	}
}
