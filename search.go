package goT411

const (
	offset = iota
	limit
	cat
	subcat
	lang
	quality
	ep
	season
)

//Torrent represent a torrent
type Torrent struct {
	ID string `json:"id"`
}

//Torrents represent a list of Torrent
type Torrents []Torrent

//SearchTorrents return a list of torrent that corresponds to query and params
func (c Client) SearchTorrents(query string, params map[int]string) (Torrents, error) {
	var out Torrents

	return out, nil
}
