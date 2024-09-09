package html_xss

import (
	"fmt"
	"html"
	"net/http"
)

type Server struct {
}

func (s *Server) HelloHandler(w http.ResponseWriter, req *http.Request) {
	helloTemplate := `
	<html>
	<head></head>
	<body>
		<p>Hello, <b>%s</b></p>
	</body>
	</html>
	`
	name := req.URL.Query().Get(html.EscapeString("name"))
	fmt.Fprintf(w, helloTemplate, name)
}
