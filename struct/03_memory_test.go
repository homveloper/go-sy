package struct_test

import (
	"testing"
	"unsafe"
)

type User03 struct {
	age   int32
	score float64
}

type MemoryWaste struct {
	a int8  // 1	pad 7byte
	b int64 // 8
	c int8  // 1
	d int64 // 8	pad 7byte
} // total 32byte

type MemoryPadding struct {
	a int8  // 1
	c int8  // 1	pad 6byte
	b int64 // 8
	d int64 // 8
} // total 24byte

/*
	고언어도 C언어와 동일하게 메모리 패딩을 한다.
	이는 CPU 동작 방식과 성능 향상을 위해서 이다.

	struct 멤버 순서에 따른 메모리 최적화는 컴파일 타임에 수행되지 않는다.
	메모리 최적화를 수행 하려면 직접 멤버 변수의 선언 순서를 알맞게 배치해야 한다.
*/
func TestStructSize(t *testing.T) {
	user := User03{
		age:   50,
		score: 50.0,
	}

	t.Log(unsafe.Sizeof(user)) // 16

	t.Log(unsafe.Sizeof(MemoryWaste{}))   // 32byte
	t.Log(unsafe.Sizeof(MemoryPadding{})) // 24byte
}
