package random

import (
	"math"
	"math/rand"
	"strings"
	"time"
)

// ref: github.com/labstack/gommon/random/random.go
const (
	// Uppercase ...
	Uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// Lowercase ...
	Lowercase = "abcdefghijklmnopqrstuvwxyz"
	// Alphabetic ...
	Alphabetic = Uppercase + Lowercase
	// Numeric ...
	Numeric = "0123456789"
	// Alphanumeric ...
	Alphanumeric = Alphabetic + Numeric
	// Symbols ...
	Symbols = "`" + `~!@#$%^&*()-_+={}[]|\;:"<>,./?`
	// Hex ...
	Hex = Numeric + "abcdef"
	// HumanReadable ...
	HumanReadable = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"
)

var src = rand.NewSource(time.Now().UnixNano())

// String generates rondom string.
// ref: https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
func String(length int, charsets ...string) string {
	charset := strings.Join(charsets, "")
	if charset == "" {
		charset = Alphanumeric
	}

	charN := len(charset)                                // Alphanumeric なら 62 文字
	idxBits := int(math.Ceil(math.Log2(float64(charN)))) // idxBits = 6; 2 ^ 6 = 64 >= 62文字
	idxMask := int64(1<<idxBits - 1)                     // 00000001 -> 01000000 -> 00111111 -> 下位 6 bit
	idxMax := 63 / idxBits                               // 63/6 = 10 -> 1回の擬似乱数生成で10文字を取得できる

	b := make([]byte, length)
	for i, cache, remain := length-1, src.Int63(), idxMax; i >= 0; {
		if remain == 0 {
			// 1回の擬似乱数生成で取得できる残り文字数が0になったら擬似乱数を再度生成する
			cache, remain = src.Int63(), idxMax
		}
		// 生成した擬似乱数の下位6ビットをビット演算で取得
		if idx := int(cache & idxMask); idx < len(charset) {
			b[i] = charset[idx]
			i-- // 1文字取得する度にインデックスを移動させる
		}
		// 擬似乱数の下位6ビットしか利用していないので、右シフトして次の上位6ビットを利用する
		cache >>= idxBits
		// 6ビット利用する度に残り回数を減らす
		remain--
	}
	return string(b)
}

// StringPreOptimized generates rondom string.
// ref: github.com/labstack/gommon/random/random.go
func StringPreOptimized(length uint8, charsets ...string) string {
	charset := strings.Join(charsets, "")
	if charset == "" {
		charset = Alphanumeric
	}
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[src.Int63()%int64(len(charset))]
	}
	return string(b)
}
