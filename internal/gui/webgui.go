package gui

import (
	"net/url"
	"strconv"

	"github.com/zserge/webview"
)

const (
	view_main_start = `
<!DOCTYPE html>
<html><head></head>
  <body>
    <center><h1 unselectable="on">Free port:</h1><h2>`
	view_main_end = `<h2></center>
  </body>
</html>
	`
)

// StartGUI launch gui
func StartGUI(port int) {
	datauri := `data:text/html,`
	datauri += url.PathEscape(view_main_start + strconv.Itoa(port) + view_main_end)
	webview.Open("Show free port",
		datauri, 300, 300, true)

}
