package main

import (
	"encoding/base64"
	"fmt"
	"hash/fnv"
)

func generateUid(body string) (string, error) {
	h := fnv.New64a()
	_, err := h.Write([]byte(body))
	if err != nil {
		return "", err
	}
	uid := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return uid, nil
}
func main() {

	uid, _ := generateUid("1234567")
	fmt.Println(uid)
}
