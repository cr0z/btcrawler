/**
 * Created by 93201 on 2017/5/10.
 */
package db

import "time"

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
	tick := time.Tick(time.Minute)
	for {
		select {
		case d := <-t.dChan:
			t.btCache = append(t.btCache, d)
		case <-tick:
			t.save()
		case <-t.sChan:
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

func (t *btSaveTask) save() {
	temp := t.btCache
	t.btCache = []*BitTorrent{}
	//TODO save to db
	tx := mdb.Begin()
	for _,v:=range temp{
		tx.Create(v)
	}
	tx.Commit()
}
