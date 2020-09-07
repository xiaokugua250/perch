package model

type BaseReponse struct {
	Message string `json:"message"`
	Kind    string `json:"kind"`
	Code    int    `json:"code"`
	Total   int64    `json:"total"`
}
type ResultReponse struct {
	BaseReponse
	Spec interface{} `json:"spec"`
}
