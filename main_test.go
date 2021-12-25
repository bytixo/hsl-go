package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

type HashResult []struct {
	Hsl string `json:"hsl"`
	Req string `json:"req"`
}

func TestSolve(t *testing.T) {
	// result.json contain 5k correct hsl hashe's
	result, err := os.Open("result.json")
	if err != nil {
		panic(err)
	}
	defer result.Close()

	b, _ := ioutil.ReadAll(result)

	var hashResult HashResult
	err = json.Unmarshal(b, &hashResult)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Testing %d hsl hashe's", len(hashResult))
	for _, hash := range hashResult {

		res := Solve(hash.Req)
		z := strings.Split(hash.Hsl, ":")
		w := strings.Split(res, ":")

		// Get the last arg to create hsl hash
		if !(w[5] == z[5]) {
			t.Fatalf(`solve("") = %q, %v, want %s, error`, w, err, hash.Hsl)
		}
	}

}
