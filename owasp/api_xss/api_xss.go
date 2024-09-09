package api_xss

import (
	"net/http"
	"text/template"
)

type Server struct {
}
type Person struct {
	Name string
}

func (s *Server) HelloHandler(w http.ResponseWriter, req *http.Request) {
	helloTemplate := `{{.}}`
	tmpl := template.New("hello")
	tmpl, _ = tmpl.Parse(helloTemplate)
	name := req.URL.Query().Get("name")

	_ = tmpl.Execute(w, name)
}
