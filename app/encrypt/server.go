package encrypt

import (
	"bytes"
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	// "strings"
	"DistributedComputingSystem/app/encrypt/helpers"
)
type Request struct{
	Text string `json:"text"`
	Key  string `json:"key"`

}
type Response struct{
	Text string `json:"text"`
	Err  string `json:"error"`

}

func RegisterHandlers() {
	handler := new(encryptHandler)
	http.Handle("/encrypt", handler)
	http.Handle("/decrypt", handler)
}

type encryptHandler struct{}

func (eh encryptHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var response string
	srv:=helpers.EncryptServiceInstance{}
	// pathSegments := strings.Split(r.URL.Path, "/")
	request,err:=eh.decode(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	response,err= srv.Encrypt(request.Key,request.Text)
	// if pathSegments[0]=="encrypt" {
	// 	respose,err=srv.Encrypt(request.Text,request.Key)
	// 	if err != nil {
	// 		return
	// 	}

	// }else{
	// 	respose,err=srv.Decrypt(request.Text,request.Key)
	// 	if err != nil {
	// 		return
	// 	}
	// }
	w.WriteHeader(http.StatusOK)
	data,err := eh.toJSON(response)
	w.Header().Add("content-type", "application/json")
	w.Write(data)
}
func (encryptHandler) decode(r *http.Request) (Request, error) {
	var request Request
	json.NewDecoder(r.Body).Decode(&request)
	return request, nil
}
func (encryptHandler) encode(w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
func (encryptHandler) toJSON(obj interface{}) ([]byte, error) {
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	err := enc.Encode(obj)
	if err != nil {
		return nil, fmt.Errorf("Failed to serialize response: %q", err)
	}
	return b.Bytes(), nil
}