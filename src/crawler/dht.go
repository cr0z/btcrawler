/**
 * Created by 93201 on 2017/5/10.
 */
package crawler

import (
	"encoding/hex"
	"github.com/shiyanhui/dht"
	"log"
)

type DhtCrawler struct {
	wire   *dht.Wire
	c      chan int
	handle func(torrent *BitTorrent)
}

func NewDhtCrawler() *DhtCrawler {
	c := &DhtCrawler{}
	c.wire = dht.NewWire(65536, 1024, 256)
	c.c = make(chan int)
	c.handle = defaultHandler
	return c
}

func (c *DhtCrawler) Run() {
	go func() {
		for {
			select {
			case resp := <-c.wire.Response():
				metadata, err := dht.Decode(resp.MetadataInfo)
				if err != nil {
					log.Println(err)
					continue
				}
				info := metadata.(map[string]interface{})
				if _, ok := info["name"]; !ok {
					continue
				}
				bt := BitTorrent{
					InfoHash: hex.EncodeToString(resp.InfoHash),
					Name:     info["name"].(string),
				}
				if v, ok := info["files"]; ok {
					files := v.([]interface{})
					bt.Files = make([]File, len(files))

					for i, item := range files {
						f := item.(map[string]interface{})
						bt.Files[i] = File{
							Path:   f["path"].([]interface{}),
							Length: f["length"].(int),
						}
					}
				} else if _, ok := info["length"]; ok {
					bt.Length = info["length"].(int)
				}
				c.handle(&bt)
			case <-c.c:
				return
			}
		}
	}()
	go c.wire.Run()
	config := dht.NewCrawlConfig()
	config.OnAnnouncePeer = func(infoHash, ip string, port int) {
		c.wire.Request([]byte(infoHash), ip, port)
	}
	d := dht.New(config)
	d.Run()
}

func (c *DhtCrawler) SetHandler(h func(torrent *BitTorrent)) *DhtCrawler {
	c.handle = h
	return c
}

func defaultHandler(bt *BitTorrent) {
	log.Printf("%s\n\n", bt.JsonString())
}
