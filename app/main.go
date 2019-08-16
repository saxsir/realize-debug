package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func get() []byte {
	resp, err := http.Get("https://api.github.com/search/repositories?sort=stars&q=atoma")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return b
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// b := get()
		fmt.Println("aaa")
		s := []int{1, 2, 3}
		fmt.Println(s[4])
		// fmt.Println(string(b))
		// fmt.Fprintf(os.Stderr, "aaaaaaaa")
		// fmt.Fprintf(os.Stderr, string(b))
		fmt.Println("bbb")

		fmt.Fprintf(w, string(b))
	})

	http.ListenAndServe(":8081", nil)
}
