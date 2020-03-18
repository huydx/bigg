

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	_ "net/http/pprof"
	"github.com/huydx/bigg/trie"
)

var (
	trie = trie.NewTrie()
)

type KipalogPost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type KipalogPosts []KipalogPost

func index() error {
	k := make([]KipalogPost, 0)
	f, err := os.Open("/tmp/kipalog.json")
	if err != nil {
		return err
	}
	c, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	enc := json.NewDecoder(bytes.NewReader(c))
	err = enc.Decode(&k)
	if err != nil {
		return err
	}

	var count = 0

	for _, p := range k {
		if p.Title != "" {
			t := p.Title
			trie.Insert(normalize(p.Title), p.Title)
			count++
			for i, c := range t {
				if string(c) == " " {
					trie.Insert(normalize(p.Title[i+1:]), p.Title)
					count++
				}
			}
		}
	}

	fmt.Printf("indexed %d documents\n", count)
	fmt.Printf("indexed %d nodes\n", trie.Count())
	return nil
}

func main() {
	if err := index(); err != nil {
		panic(err)
	}
	debug()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if r.Method != "POST" {
			return
		}
		j := make([]string, 0)
		r.ParseForm()
		query := r.Form.Get("q")
		results := trie.FindPrefix(query, 5)
		for _, result := range results {
			j = append(j, result.Origin())
		}
		json.NewEncoder(w).Encode(j)
	})
	http.ListenAndServe(":9999", nil)
}

var (
	vietnameseMap = map[string][]string{
		"a": {"à", "á", "ạ", "ả", "ã", "â", "ầ", "ấ", "ậ", "ẩ", "ẫ", "ă", "ằ", "ắ", "ặ", "ẳ", "ẵ"},
		"e": {"è", "é", "ẹ", "ẻ", "ẽ", "ê", "ề", "ế", "ệ", "ể", "ễ"},
		"i": {"ì", "í", "ị", "ỉ", "ĩ"},
		"o": {"ò", "ó", "ọ", "ỏ", "õ", "ô", "ồ", "ố", "ộ", "ổ", "ỗ", "ơ", "ờ", "ớ", "ợ", "ở", "ỡ"},
		"u": {"ù", "ú", "ụ", "ủ", "ũ", "ư", "ừ", "ứ", "ự", "ử", "ữ"},
		"y": {"ỳ", "ý", "ỵ", "ỷ", "ỹ"},
		"d": {"đ"},
		"A": {"À", "Á", "Ạ", "Ả", "Ã", "Â", "Ầ", "Ấ", "Ậ", "Ẩ", "Ẫ", "Ă", "Ằ", "Ắ", "Ặ", "Ẳ", "Ẵ"},
		"E": {"È", "É", "Ẹ", "Ẻ", "Ẽ", "Ê", "Ề", "Ế", "Ệ", "Ể", "Ễ"},
		"I": {"Ì", "Í", "Ị", "Ỉ", "Ĩ"},
		"O": {"Ò", "Ó", "Ọ", "Ỏ", "Õ", "Ô", "Ồ", "Ố", "Ộ", "Ổ", "Ỗ", "Ơ", "Ờ", "Ớ", "Ợ", "Ở", "Ỡ"},
		"U": {"Ù", "Ú", "Ụ", "Ủ", "Ũ", "Ư", "Ừ", "Ứ", "Ự", "Ử", "Ữ"},
		"Y": {"Ỳ", "Ý", "Ỵ", "Ỷ", "Ỹ"},
		"D": {"Đ"},
	}
)

// normalize a string with Vietnamese and space
// into non-space non-accent mark
func normalize(in string) string {
	var out string
	var invertMap = map[string]string{}
	for k, v := range vietnameseMap {
		for _, v2 := range v {
			invertMap[v2] = k
		}
	}

	for _, i := range in {
		if _, ok := invertMap[string(i)]; ok  {
			out += invertMap[string(i)]
			continue
		}
		if string(i) == " " {
			continue
		}
		out += string(i)
	}

	return out
}

func debug() {
	go func() {
		http.ListenAndServe(":8080", nil)
	}()
}
