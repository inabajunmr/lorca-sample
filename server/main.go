package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/zserge/lorca"
)

var ui lorca.UI

func main() {
	ui, _ = lorca.New("", "", 480, 320)
    	defer ui.Close()
	ui.Bind("hello", hello)
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	go http.Serve(ln, http.FileServer(http.Dir("./www")))
	ui.Load(fmt.Sprintf("http://%s", ln.Addr()))
	defer ui.Close()
	// Wait until UI window is closed
	<-ui.Done()
}

func hello() {
  ui.Eval("alert('hello lorca')")
}

