package service

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gobuffalo/packr"
	"github.com/gorilla/mux"
)

// Config TODO
type Config struct {
	IPAddress        string
	Lifecycle        string
	Name             string
	Port             string
	StaticFileServer http.Handler
}

// Lifecycle TODO
type Lifecycle struct {
	Name      string
	Example   []byte
	Structure []byte
	Version   string
}

// Page TODO
type Page struct {
	HTML     []byte
	Template []byte
	Config   PageTemplate
}

// PageTemplate TODO
type PageTemplate struct {
	Title string
	Route string
}

// Start TODO
func (c *Config) Start() {
	address := c.IPAddress + ":" + c.Port
	assetdir := "./assets/" + c.Name

	// create an in memory instance of the service web assets
	box := packr.NewBox(assetdir)
	c.StaticFileServer = http.FileServer(box)

	// initialize handler for service home page
	pageTemplate, err := box.Find("index.html")
	if err != nil {
		log.Fatal(err)
	}

	title := "SDE " + strings.Title(c.Name)
	page := &Page{
		Template: pageTemplate,
		Config:   PageTemplate{Title: title, Route: "127.0.0.1"},
	}

	if err := page.generate(); err != nil {
		log.Fatal(err)
	}

	// initialize handler for service lifecycle configuration
	exampleLifecycle, err := box.Find("lifecycle.json")
	if err != nil {
		log.Fatal(err)
	}

	lifecycle := &Lifecycle{
		Example:   exampleLifecycle,
		Name:      c.Name,
		Structure: exampleLifecycle,
		Version:   "0.1",
	}

	if c.Lifecycle != "" {
		structure, err := ioutil.ReadFile(c.Lifecycle)
		if err != nil {
			log.Fatal(err)
		}
		lifecycle.Structure = structure
	}

	// configure routing and launch the service
	router := mux.NewRouter()
	router.HandleFunc("/", page.handler).Methods("GET")
	router.HandleFunc("/lifecycle", lifecycle.handler).Methods("GET")
	router.HandleFunc("/ping", ping).Methods("GET")
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(box)))

	log.Printf("service:%s started on %s\n", c.Name, address)
	http.ListenAndServe(address, router)
}

func (l *Lifecycle) handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(l.Structure)
}

func (p *Page) handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "max-age=0")
	w.Write(p.HTML)
}

func (p *Page) generate() error {
	page, err := template.New("template").Funcs(getFuncMaps()).Parse(string(p.Template))
	if err != nil {
		log.Printf("error parsing template %s %v\n", p.Config.Title, err)
		return err
	}

	buf := new(bytes.Buffer)
	if err := page.ExecuteTemplate(buf, "Template", p.Config); err != nil {
		log.Printf("error executing template %s %v\n", p.Config.Title, err)
		return err
	}

	html, err := ioutil.ReadAll(buf)
	if err != nil {
		return err
	}
	p.HTML = html

	return nil
}

func ping(w http.ResponseWriter, r *http.Request) {
	status := map[string]string{"alive": "true"}

	js, err := json.Marshal(status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func getFuncMaps() template.FuncMap {
	return template.FuncMap{
		"capitalize": func(word string) string {
			return strings.Title(word)
		},
	}
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
