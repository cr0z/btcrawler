/**
 * Created by 93201 on 2017/5/10.
 */
package db

import (
	"crawler"
	"encoding/json"
)

type File struct {
	Path   []interface{} `json:"path"`
	Length int           `json:"length"`
}

/**
CREATE TABLE `t_bt`(
	`id`  int NOT NULL AUTO_INCREMENT ,
	`infohash` varchar(255) NOT NULL,
	`name` varchar(255),
	`length` int,
	`files` varchar(2047),
	PRIMARY KEY (`id`)
)
*/

type BitTorrent struct {
	ID       int    `gorm:"columns:id" json:"id"`
	Infohash string `gorm:"columns:infohash" json:"infohash"`
	Name     string `gorm:"columns:name" json:"name"`
	Length   int    `gorm:"columns:length" json:"length,omitempty"`

	Files string  `gorm:"columns:files" json:"-"` //[{"path":["12 Nyan Nyan Final Attack Frontier Greatest Hits!.mp3"],"length":18070628}]
	F     []*File `gorm:"-" json:"files"`
}

func (b *BitTorrent) TableName() string {
	return "t_bt"
}

func (b *BitTorrent) From(t *crawler.BitTorrent) *BitTorrent {
	fb, _ := json.Marshal(t.Files)

	b.Infohash = t.InfoHash
	b.Name = t.Name
	b.Length = t.Length
	b.Files = string(fb)
	return b
}

func (b *BitTorrent) Insert() error {
	return mdb.Create(b).Error
}

type BTQuery struct {
	qStruct
}

func (b *BTQuery) Query() ([]*BitTorrent, error) {
	db := mdb
	if b.Count > 0 {
		if b.Page == 0 {
			b.Page = 1
		}
		db = db.Limit(b.Count).Offset(b.Count * (b.Page - 1))
	}
	if b.Order != "" {
		db = db.Order(b.Order)
	}

	var bts = []*BitTorrent{}
	db = db.Find(&bts)
	return bts, db.Error
}
