package random

import (
	"math"
	"math/rand"
	"strings"
	"time"
)

type (
	// Random .
	Random struct {
	}
)

// Charsets
const (
	Uppercase    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Lowercase    = "abcdefghijklmnopqrstuvwxyz"
	Alphabetic   = Uppercase + Lowercase
	Numeric      = "0123456789"
	Alphanumeric = Alphabetic + Numeric
	Symbols      = "`" + `~!@#$%^&*()-_+={}[]|\;:"<>,./?`
	Hex          = Numeric + "abcdef"
	Human        = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"
)

var (
	gSrc    = rand.NewSource(time.Now().UnixNano())
	gRondom = &Random{}
)

// String generates a random string
// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
func (r *Random) String(length uint8, charsets ...string) string {
	charset := strings.Join(charsets, "")
	if charset == "" {
		charset = Alphanumeric
	}

	charCnt := len(charset)
	idxBits := uint8(math.Ceil(math.Log2(float64(charCnt))))
	idxMask := int64(1<<idxBits - 1)
	idxMax := 63 / idxBits
	// fmt.Println("length:", length, "charCnt:", charCnt, "idxBitx", idxBits, "idxMask", idxMask, "idxMax", idxMax)

	b := make([]byte, length)
	for i, cache, remain := 0, gSrc.Int63(), idxMax; i < int(length); {
		// fmt.Println("i", i)
		if remain == 0 {
			// 1回の擬似乱数生成で取得できる残り文字数が0になったら擬似乱数を再度生成する
			cache, remain = gSrc.Int63(), idxMax
		}
		// 生成した擬似乱数の下位6ビットをビット演算で取得
		if idx := int(cache & idxMask); idx < len(charset) {
			// a := b[i]
			// c := charset[idx]
			// fmt.Println(a, c)
			b[i] = charset[idx]
			i++ // 1文字取得する度にインデックスを移動させる
		}
		// 擬似乱数の下位6ビットしか利用していないので、右シフトして次の上位6ビットを利用する
		cache >>= idxBits
		// 6ビット利用する度に残り回数を減らす
		remain--
	}

	return string(b)
}

// StringOld generates a random string
func (r *Random) StringOld(length uint8, charsets ...string) string {
	charset := strings.Join(charsets, "")
	if charset == "" {
		charset = Alphanumeric
	}
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Int63()%int64(len(charset))]
	}
	return string(b)
}

// String .
func String(length uint8, charsets ...string) string {
	return gRondom.String(length, charsets...)
}
