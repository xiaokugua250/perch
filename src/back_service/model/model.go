package model

type EmailBody struct {
	From     string      `json:"from"`
	To       string      `json:"to"`
	CopyTo   string      `json:"copy_to"`
	ReplyTo  string      `json:"reply_to"`
	Subject  string      `json:"subject"`
	Body     string      `json:"body"`
	BodyHTML string      `json:"body_html"`
	Attach   interface{} `json:"attach"`
}

type ResultReponse struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	ErrMsg  error       `json:"errMsg"`
	Spec    interface{} `json:"spec"`
}
