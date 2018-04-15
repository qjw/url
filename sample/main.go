package main

import (
	"encoding/json"
	"fmt"

	"github.com/qjw/url"
)

//func testRedis(s string) {
//	u, _ := url.Parse(s)
//	fmt.Println(u.Scheme)
//	fmt.Println(u.User)
//	//	fmt.Println(u.User.Username())
//	//	p, _ := u.User.Password()
//	//	fmt.Println(p)
//	fmt.Println(u.Host)
//	host, port, _ := net.SplitHostPort(u.Host)
//	fmt.Println(host)
//	fmt.Println(port)
//	fmt.Println(u.Path)
//	fmt.Println(u.Fragment)
//	fmt.Println(u.RawQuery)
//	//	m, _ := url.ParseQuery(u.RawQuery)
//}

func check(data interface{}, err error) {
	if err == nil {
		data, _ := json.MarshalIndent(data, "", " ")
		fmt.Println(string(data))
	} else {
		fmt.Println(err)
	}
}

func main() {
	r, err := url.ParseRedis("redis://:pwd@localhost:1234/8")
	check(r, err)
	r, err = url.ParseRedis("redis://localhost:1234")
	check(r, err)
	r, err = url.ParseRedis("redis://localhost")
	check(r, err)
	r, err = url.ParseRedis("redis://:password@localhost")
	check(r, err)
}
