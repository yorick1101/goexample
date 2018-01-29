package main
import (

  "fmt"
  "regexp"
)

func main(){

  match,_ :=regexp.MatchString("p[a-z]+ch","peach")
  fmt.Println(match)

  r,_ := regexp.Compile("p([a-z]+)ch")

  fmt.Println(r.MatchString("peach"))

  fmt.Println(r.FindStringIndex("peach bunch"))

  fmt.Println(r.FindStringSubmatch("peach bunch"))

  fmt.Println(r.FindStringSubmatchIndex("peach bunch"))

// Second parameter indicate the maxmium number of match, -1 means no limit
  fmt.Println(r.FindAllString("peach bunch beach", -1))
  fmt.Println(r.FindAllStringSubmatch("peach punch pinch", -1))
  fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))

//throw exception if error
  r = regexp.MustCompile("p([a-z]+)ch")
  fmt.Println(r)

  fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

  in := []byte("a peach")
  out := r.ReplaceAllFunc(in, bytes.ToUpper)
  fmt.Println(string(out))
}
