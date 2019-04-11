package main

import (
	"fmt"
	b64 "encoding/base64"
)

func main() {
	//data := "abc123!?$*&()'-=@~"
	data := "g01@ng_enc0de@pr@ctice"

	// encoder requires type []byte
	// standard base64 encoder
	encode_data := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(encode_data)

	decode_data, _ := b64.StdEncoding.DecodeString(encode_data)
	fmt.Println(string(decode_data))
	fmt.Println("===== ===== ===== ===== =====")

	// URL base64 encoder
	url_encode_data := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(url_encode_data)

	url_decode_data, _ := b64.URLEncoding.DecodeString(url_encode_data)
	fmt.Println(string(url_decode_data))

}
