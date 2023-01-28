package test

import (
	"awesomeProject/wangdejiang/src/service"
	"fmt"
	"testing"
)

func TestToken(t *testing.T) {
	claim := service.UserClaims{
		Name:     "aa",
		Identity: "user_1",
		IsAdmin:  1,
	}
	tokenStr, err := claim.GenerateToken()
	if err != nil {
		t.Fail()
	}
	fmt.Println("tokenStr: ", tokenStr)
	claim2 := new(service.UserClaims)
	err2 := claim2.ParseToken(tokenStr)
	if err2 != nil {
		t.Fail()
	}
	fmt.Println("UserClaim", claim2)
}
