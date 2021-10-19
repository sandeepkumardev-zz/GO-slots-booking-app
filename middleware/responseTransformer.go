package middleware

type ResponseTransformer struct {
	Message   string      `json:"message"`
	Result    interface{} `json:"result"`
	IsSuccess bool        `json:"isSuccess"`
}
