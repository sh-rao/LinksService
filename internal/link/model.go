package link

import "time"

type LinkData struct {
	LinkID      string      `json:"link_id"`
	UserID      string      `json:"user_id"`
	DateCreated time.Time   `json:"date_created"`
	Type        string      `json:"type"`
	Data        interface{} `json:"data"`
}

type Classic struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

type ShowsList struct {
	SoldOut   Show   `json:"sold_out"`
	NotOnSale Show   `json:"not_on_sale"`
	OnSale    []Show `json:"on_sale"`
}

type MusicPlayer struct {
	Platform    string `json:"platform"`
	AudioPlayer string `json:"audio_player"`
}

type Show struct {
	Title string    `json:"title"`
	Date  time.Time `json:"date"`
	Venue string    `json:"venue"`
}

type CreateLinkRequest struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type CreateLinkResponse struct {
	LinkData LinkData `json:"link_data"`
}

type GetLinksResponse struct {
	LinksData []LinkData `json:"links_data"`
}
