package main

import (
	"html/template"
	"io/fs"
	"net/http"
)

type ServerConfig struct {
	WebFS    fs.FS
	ConfigFS fs.FS
}

type Server struct {
	templates *template.Template
	settings  Settings
	staticFS  http.Handler
}

func NewServer(cfg ServerConfig) (*Server, error) {
	tmpl, err := template.ParseFS(cfg.WebFS, "web/templates/*.html")
	if err != nil {
		return nil, err
	}

	settingsFile, err := fs.ReadFile(cfg.ConfigFS, "config/settings.json")
	if err != nil {
		return nil, err
	}

	var settings Settings
	if err := settings.UnmarshalJSON(settingsFile); err != nil {
		return nil, err
	}

	staticSubFS, err := fs.Sub(cfg.WebFS, "web/static")
	if err != nil {
		return nil, err
	}

	return &Server{
		templates: tmpl,
		settings:  settings,
		staticFS:  http.FileServer(http.FS(staticSubFS)),
	}, nil
}

func (s *Server) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.StripPrefix("/static/", s.staticFS))
	mux.HandleFunc("/", s.index)

	return mux
}

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"AppName": s.settings.AppName,
	}

	if err := s.templates.ExecuteTemplate(w, "index.html", data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
