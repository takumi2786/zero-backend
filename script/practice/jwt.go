package main

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

/*
HMAC

	受信者と送信者の間でメッセージが改ざんされていないことを検証する仕組み。

ランダム文字列を生成

	openssl rand -hex 32

送信者: ペイロードを秘密鍵で署名し、値と共に送信

	echo -n "value" | openssl dgst -sha256 -mac hmac -macopt hexkey:0a4a2fb5bab4c933135571c483bc15f1ea419fd33fdbbb9feddfe10278f030b1

受信者: ペイロードを秘密鍵で署名し、送信者からの値と比較

	echo -n "value" | openssl dgst -sha256 -mac hmac -macopt hexkey:0a4a2fb5bab4c933135571c483bc15f1ea419fd33fdbbb9feddfe10278f030b1

参考

	https://qiita.com/syamobariyuta/items/184e56909e79bea0af47#hmac
*/
var SECRET_KEY string = "0a4a2fb5bab4c933135571c483bc15f1ea419fd33fdbbb9feddfe10278f030b1"

func main() {
	claims := jwt.MapClaims{
		"user_id": 12345678,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"messgae": "hello",
	}

	/*
		送信者: ペイロードを秘密鍵で署名し、値と共に送信
	*/
	// ヘッダーとペイロードの生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Printf("Header: %#v\n", token.Header) // Header: map[string]interface {}{"alg":"HS256", "typ":"JWT"}
	fmt.Printf("Claims: %#v\n", token.Claims) // CClaims: jwt.MapClaims{"exp":1634051243, "user_id":12345678}

	tokenString, _ := token.SignedString([]byte(SECRET_KEY))
	fmt.Println("tokenString:", tokenString) // tokenString: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTUwMDM3MzYsInVzZXJfaWQiOjEyMzQ1Njc4fQ.4fJLJPc3aZpsDvc6Opgu-EiiUx1zPhSiAKXB9pqfKBQ

	header_decoded, err := base64.StdEncoding.DecodeString("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9")
	if err != nil {
		fmt.Println(err)
	}
	payload_decoded, err := base64.StdEncoding.DecodeString("eyJleHAiOjE2OTUwMDM3MzYsInVzZXJfaWQiOjEyMzQ1Njc4fQ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(
		"header_decoded: %s payload_decoded: %s \n",
		string(header_decoded),
		string(payload_decoded),
	)

	/*
		受信者: ペイロードを秘密鍵で署名し、送信者からの値と比較
	*/
	// jwtライブラリでトークンを読み込む
	token, err = jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(SECRET_KEY), nil
		},
	)
	// トークンを検証する
	if !token.Valid {
		panic("token is invalid")
	}
	// トークンの中身を取得する
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		fmt.Println(claims)
	} else {
		fmt.Println(err)
	}
}
