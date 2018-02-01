package main

import (
  "fmt";
  "net";
  "net/url"
)

var p = fmt.Println

func main(){

  s := "postgres://user:pass@host.com:5432/path?k=v&k=v2#f"
  u, err := url.Parse(s)
  if err != nil {
    panic(err)
  }
  user := u.User.Username()
  passwd, _ := u.User.Password()
  p(user, passwd)
  host, port, _ := net.SplitHostPort(u.Host)
  p(host, port)

  p(u.Path)
  p(u.Fragment)
  p(u.RawQuery)
  m, _ :=url.ParseQuery(u.RawQuery)
  p(m)
  p(m["k"])
}
