package api

type Page struct {
	Page      int `json:"page,omitempty"`
	Size      int `json:"size,omitempty"`
	TotalSize int `json:"totalSize,omitempty"`
}

type User struct {
	Uid string `json:"uid,omitempty"`
	Sid string `json:"sid,omitempty"`
}
