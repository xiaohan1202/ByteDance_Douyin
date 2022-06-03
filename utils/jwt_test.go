package utils

import (
	"fmt"
	"testing"
)

func TestJwt(t *testing.T) {
	//var userId int64
	//userId = 3
	//token, err := SignToken(userId)
	//time.Sleep(time.Second * 2)
	//if err != nil {
	//	panic(err)
	//}
	id, err := ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjQsImV4cCI6MTY1NDE3MDEyOCwiaWF0IjoxNjU0MTQ0OTI4LCJpc3MiOiJkb3V5aW4iLCJzdWIiOiJ1c2VyIHRva2VuIn0.EMYwYqCNVb7VeMWENVLyIvXQ3lclw4HLaL5DHcx-tyg")
	if err != nil {
		panic(err)
	}
	fmt.Printf("userId: %#v", id)
}
