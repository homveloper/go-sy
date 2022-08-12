package struct_test

import (
	"testing"
)

type User struct {
	name string
	id   string
	age  int
}

type VIPUser struct {
	user  User // 내장 멤버 변수 방식
	level int
	price int
}

type ClientUser struct {
	User // 포함된 필드 방식
}

func TestStructEmbedding(t *testing.T) {
	vip := VIPUser{
		user: User{
			name: "방구석 개발자",
			id:   "homveloper",
			age:  50,
		},
		level: 5,
		price: 400,
	}

	client := ClientUser{
		User: User{
			name: "방구석 개발자",
			id:   "homveloper",
			age:  50,
		},
	}

	t.Log(vip)
	t.Log(client)
}
