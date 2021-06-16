package helpers

type Request struct{
	Text string `json:"text"`
	Key  string `json:"key"`

}
type Response struct{
	Text string `json:"text"`
	Err  string `json:"error"`

}
