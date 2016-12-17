package api

type Page struct {
	Page      int64 `json:"page,omitempty"`
	Size      int64 `json:"size,omitempty"`
	TotalSize int64 `json:"totalSize,omitempty"`
}

type User struct {
	Uid int64  `json:"uid,omitempty"`
	Sid string `json:"sid,omitempty"`
}
