package potto

type Response struct {
	Text string `json:"text"`
}

func NewResponse(text string) *Response {
	return &Response{
		Text: text,
	}
}
