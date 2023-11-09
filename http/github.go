package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	// resp, err := http.Get("https://api.github.com/users/librairica")
	// if err != nil {
	// 	log.Fatalf("error: %s", err)
	// 	// fatalf is the equivalent of:
	// 	// log.Printf("error: %s", err)
	// 	// os.Exit(1)
	// }
	// if resp.StatusCode != http.StatusOK {
	// 	log.Fatalf("error: %s", resp.Status)
	// }
	// fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))
	// headers can be repeated, so we can't use a map
	// Header.Get is also case-insensitive, unlike maps, e.g. content-type == Content-Type

	// initialization and condition
	// if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
	// 	log.Fatalf("error: can't copy - %s", err)
	// }
	// once you read from the body you can't read again

	// var r Reply
	// dec := json.NewDecoder(resp.Body)
	// if err := dec.Decode(&r); err != nil {
	// 	log.Fatalf("error: can't decode = %s", err)
	// }
	// fmt.Println(r)
	// fmt.Printf("%#v\n", r)
	// in browser and terminal, json is formatted bc the user agent is human
	// running go program returns compact format for machines
	// you can use jq to format nicely

	// serialization is the act of taking a data structure in your language and converting it into a sequence of bytes
	// deserialization is the opposite

	name, num, err := githubInfo("librairica")
	if err != nil {
		log.Fatalf("error: can't get GH info: %s", err)
	}
	fmt.Printf("name: %s\n", name)
	fmt.Printf("num repos: %d\n", num)
}

// you can use Go: Add Tags to Struct Fields shortcut in command palette
type Reply struct {
	Name string
	// Public_Repos int
	NumRepos int `json:"public_repos"`
}

// githubInfo returns name and number of public repos for login
func githubInfo(login string) (string, int, error) {
	url := "https://api.github.com/users/" + url.PathEscape(login)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		return "", 0, errors.New(resp.Status)
	}
	var r struct { // anonymous struct example
		Name string
		// Public_Repos int
		NumRepos int `json:"public_repos"`
	}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		log.Fatalf("error: can't decode = %s", err)
	}
	return r.Name, r.NumRepos, nil
}

/* JSON -> Go
true/false -> true/false
string -> string
null -> nil
number -> float64, float32, int8, int16, int32, int64, int, uint8, ...
array -> []any ([]interface{})
object -> map[string]any, struct

JSON -> io.Reader -> Go: json.Decoder
Go -> io.Writer -> JSON: json.Encoder
JSON -> []byte -> Go: json.Unmarshal
Go -> []byte -> JSON json.Marshal
*/
