package msgcrypt

import (
	"fmt"
	"testing"
)

func TestAesEncrypt(t *testing.T) {
	var str = []string{"hello", "world", "steven", "你好"}
	for _, v := range str {
		res, err := AesEncrypt(v)
		if err != nil {
			panic(err)
		}
		fmt.Println(res)
	}
}

func ExampleAesEncrypt() {
	fmt.Println(AesEncrypt("Encrypt the string with AES crypt"))
}

func BenchmarkAesEncrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res, err := AesEncrypt("hello world")
		if err != nil {
			panic(err)
		}
		fmt.Println(res)
	}
}
