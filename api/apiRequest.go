package api

type ApiRequest struct {
	Id     string                 `json:"id"`
	Sign   string                 `json:"sign"`
	Client Client                 `json:"client"`
	Page   Page                   `json:"page"`
	User   User                   `json:"user"`
	Data   map[string]interface{} `json:"data"`
}

type Client struct {
	Caller   string            `json:"caller"`
	Os       string            `json:"os"`
	Ver      string            `json:"ver"`
	Platfrom string            `json:"platform"`
	Ch       string            `json:"ch"`
	Ex       map[string]string `json:"ex"`
}

func (this ApiRequest) CheckData(keys ...string) bool {
	for _, key := range keys {
		if _, ok := this.Data[key]; !ok {
			return false
		}
	}
	return true
}
