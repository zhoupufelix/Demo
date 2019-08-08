package back

type Response struct{
	Code int `json:"code"`
	Msg string `json:"msg"`
}

type Controller struct {
	response Response
	template string
}


