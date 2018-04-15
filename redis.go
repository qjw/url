package url

import (
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strconv"
)

type Redis struct {
	Host     string  `json:"host"`
	Port     *int    `json:"port"`
	Db       int     `json:"db"`
	Password *string `json:"password"`
}

// redis://:pwd@localhost:1234/8
// redis://localhost:1234
// redis://localhost
// redis://:password@localhost
/*
{
 "host": "localhost:1234",
 "port": 1234,
 "db": 8,
 "password": "pwd"
}
{
 "host": "localhost:1234",
 "port": 1234,
 "db": 0,
 "password": null
}
{
 "host": "localhost",
 "port": null,
 "db": 0,
 "password": null
}
{
 "host": "localhost",
 "port": null,
 "db": 0,
 "password": "password"
}
*/
func ParseRedis(s string) (*Redis, error) {
	u, err := url.Parse(s)
	if err != nil {
		return nil, err
	}

	if u.Scheme != "redis" {
		return nil, fmt.Errorf("invalid scheme '%s'", u.Scheme)
	}

	r := &Redis{}
	// 端口/port
	r.Host = u.Host
	match, _ := regexp.MatchString(":[0-9]+$", u.Host)
	if match {
		_, port, err := net.SplitHostPort(u.Host)
		if err != nil {
			return nil, err
		}
		portInt, err := strconv.Atoi(port)
		if err != nil {
			return nil, fmt.Errorf("invalid port '%s'", port)
		}
		r.Port = &portInt
	}

	// DB
	if len(u.Path) == 0 {
		r.Db = 0
	} else if len(u.Path) > 1 && u.Path[0] == '/' {
		dbInt, err := strconv.Atoi(u.Path[1:])
		if err != nil {
			return nil, fmt.Errorf("invalid DB '%s'", u.Path[1:])
		}
		r.Db = dbInt
	} else {
		return nil, fmt.Errorf("invalid DB '%s'", u.Path)
	}

	// password
	if u.User != nil {
		if p, ok := u.User.Password(); ok {
			r.Password = &p
		}
	}

	return r, nil
}
