package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var whatIsIt string
	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	sd, _ := base64.StdEncoding.DecodeString(secret)

	whatIsIt = string(sd)
	for i := len(whatIsIt) - 1; i > -1; i-- {
		fmt.Printf("%c", whatIsIt[i])
	}
}
