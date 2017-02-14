package api

type ApiRequest struct {
	Id     string                 `json:"id"`
	Sign   string                 `json:"sign,omitempty"`
	Client Client                 `json:"client,omitempty"`
	Page   Page                   `json:"page,omitempty"`
	User   User                   `json:"user,omitempty"`
	Data   map[string]interface{} `json:"data"`
}

type Client struct {
	Caller   string            `json:"caller,omitempty"`
	Os       string            `json:"os,omitempty"`
	Ver      string            `json:"ver,omitempty"`
	Platform string            `json:"platform,omitempty"`
	Ch       string            `json:"ch,omitempty"`
	Ex       map[string]string `json:"ex,omitempty"`
}

func (c ApiRequest) CheckData(keys ...string) bool {
	for _, key := range keys {
		if _, ok := c.Data[key]; !ok {
			return false
		}
	}
	return true
}
