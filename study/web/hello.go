/**
 * @Author: frankstar
 * @Email: frankstar007@163.com
 * @Date: 2020-07-22 11:08
 * @Project : go-improve
 * @Desc: 
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloWorld(w http.ResponseWriter, r *http.Request)  {

	r.ParseForm() //解析参数
	fmt.Println(r.Form)
	fmt.Println("URL:", r.URL.Path)
	fmt.Println("Schema:", r.URL.Scheme)
	for k,v := range r.Form {
		fmt.Println(k, ":", strings.Join(v," "))
	}
	fmt.Fprintf(w, "hello world")

}

func main()  {
	http.HandleFunc("/", sayHelloWorld)
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
