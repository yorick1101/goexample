package main

import (
	b64 "encoding/base64"
	"fmt"
)

var p = fmt.Println

func main() {

	data := "abc123!?$*&()'-=@~'"

	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	uEnc2 := b64.RawURLEncoding.EncodeToString([]byte(data))

	p(uEnc)
	p(uEnc2)

}
