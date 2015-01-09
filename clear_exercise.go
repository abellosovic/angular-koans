package main

import (
	"io/ioutil"
	"regexp"
)

func main() {
	b, err := ioutil.ReadFile("src/GemStoreModule.js")
	if err != nil {
		panic(err)
	}

	r, err := regexp.Compile(`(?sm)(// -{12} START EDITING HERE -{22})(.*?)(// -{12} STOP EDITING HERE -{22})`)
	if err != nil {
		panic(err)
	}

	newb := r.ReplaceAll(b, []byte("$1\n\n$3"))

	ioutil.WriteFile("src/GemStoreModule.js", newb, 0444)
}
