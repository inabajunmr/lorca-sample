package main

import (
	"net/url"

	"github.com/zserge/lorca"
)

var ui lorca.UI

func main() {
	ui, _ = lorca.New("", "", 480, 320)
    	defer ui.Close()
	ui.Bind("hello", hello)
	// Create UI with basic HTML passed via data URI
	ui.Load("data:text/html,"+url.PathEscape(`
	<html>
		<head><title>Hello</title></head>
		<body>
			<h1>Hello, world!</h1>
			<input type='button' onclick='hello()' value='hello lorca'>
		</body>
	</html>
	`))
	defer ui.Close()
	// Wait until UI window is closed
	<-ui.Done()
}

func hello() {
  ui.Eval("alert('hello lorca')")
}

