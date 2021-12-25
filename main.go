package main

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

/*

	Code come from here https://gist.github.com/post04/4e8f60c94661979bcbef1ab60232891e
	I just fixed it and cleaned it

	HSL.js:   https://newassets.hcaptcha.com/c/3de5319d/hsl.js



*/

type payload struct {
	S int    `json:"s"`
	T string `json:"t"`
	D string `json:"d"`
	L string `json:"l"`
	E int    `json:"e"`
}

var (
	x = "0123456789/:abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func Solve(req string) string {
	JWT := strings.Split(req, ".")

	pd := &payload{}

	_ = json.Unmarshal([]byte(atob(JWT[1])), &pd)

	d := pd.D
	s := pd.S

	e := get(s, d)

	formattedDate := getFormattedDate()
	return "1:" + fmt.Sprint(s) + ":" + formattedDate + ":" + d + "::" + e
}
func get(r int, t string) string {

	for e := 0; e < 25; e++ {

		n := func() (arr []int) {
			for i := 0; i < e; i++ {
				arr = append(arr, 0)
			}
			return
		}()
		for a(n) {
			u := t + "::" + ix(n)
			if o(r, u) {
				return ix(n)
			}
		}
	}
	return ""
}
func a(r []int) bool {
	for t := len(r) - 1; t > -1; t-- {
		if r[t] < len(x)-1 {
			r[t] += 1
			return true
		}
		r[t] = 0
	}
	return false
}

func ix(r []int) string {
	e := ""
	for n := 0; n < len(r); n++ {
		e += string(x[r[n]])
	}
	return e
}

func o(r int, t string) bool {
	hashed := sha1.New()
	hashed.Write([]byte(t))
	d := hashed.Sum(nil)
	var e int
	var o []int
	for n := 0; n < 8*len(d); n++ {
		e = int(d[n/8] >> (n % 8) & 1) // Yup this is correct, Operators precedence in Go is different than in python : https://www.tutorialspoint.com/go/go_operators_precedence.htm
		o = append(o, e)
	}
	a := o[:r]

	return 0 == a[0] && index2(a, 1) >= int(r)-1 || -1 == index2(a, 1)
}
func index2(arr []int, item int) int {
	for position, thing := range arr { // Post if you see this you made me lost 2 hours
		if thing == item {
			return position
		}
	}
	return -1
}

func atob(input string) string {
	data, err := base64.RawStdEncoding.DecodeString(input)
	if err != nil {
		return ""
	}
	return string(data)
}

func getFormattedDate() string {
	timeNow := getTimeISO()
	timeNow = timeNow[:len(timeNow)-5]
	timeNow = strings.ReplaceAll(timeNow, "-", "")
	timeNow = strings.ReplaceAll(timeNow, ":", "")
	timeNow = strings.ReplaceAll(timeNow, "T", "")
	return timeNow
}
func getTimeISO() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05.999Z")
}
