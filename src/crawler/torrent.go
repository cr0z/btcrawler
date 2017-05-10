/**
 * Created by 93201 on 2017/5/10.
 */
package crawler

import "encoding/json"

type File struct {
	Path   []interface{} `json:"path"`
	Length int           `json:"length"`
}

type BitTorrent struct {
	InfoHash string `json:"infohash"`
	Name     string `json:"name"`
	Files    []File `json:"files,omitempty"`
	Length   int    `json:"length,omitempty"`
}

func (b *BitTorrent) Json() []byte {
	by, _ := json.Marshal(b)
	return by
}
func (b *BitTorrent) JsonString() string {
	return string(b.Json())
}
