package base62

import (
	"encoding/base64"
	"math"
	"strconv"
	"strings"
)

const base = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const b = 62


// Function encodes the given database ID to a base62 string
func ToBase62(num int) string{
	r := num % b
	res := string(base[r])
	div := num / b
	q := int(math.Floor(float64(div)))
	for q != 0 {
		r = q % b
		temp := q / b
		q = int(math.Floor(float64(temp)))
		res = string(base[int(r)]) + res
	}
	return string(res)
}

// Function decodes a given base62 string to database ID
func ToBase10(str string) int{
	res := 0
	for _, r := range str {
		res = (b * res) + strings.Index(base, string(r))
	}
	return res
}

func TBE62(i int) string {
	res := base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(i)))

	return res
}

func TBD62(str string) int {
	res, _ := base64.StdEncoding.DecodeString(str)
	resi, _ := strconv.Atoi(string(res))
	return resi
}