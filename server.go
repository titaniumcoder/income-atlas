package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/fs"
	"math"
	"net/http"
	"strconv"
)

type ServerConfig struct {
	WebFS    fs.FS
	ConfigFS fs.FS
	DataFS   fs.FS
}

type Server struct {
	templates *template.Template
	settings  Settings
	staticFS  http.Handler
	Expenses  []Expenses
}

func templateFuncs() template.FuncMap {
	return template.FuncMap{
		"dict": func(values ...any) map[string]any {
			d := make(map[string]any, len(values)/2)
			for i := 0; i+1 < len(values); i += 2 {
				key, _ := values[i].(string)
				d[key] = values[i+1]
			}
			return d
		},
		"has": func(m any, key string) bool {
			v, ok := m.(map[string]any)
			if !ok {
				return false
			}
			_, ok = v[key]
			return ok
		},
		"get": func(m any, key string) any {
			v, ok := m.(map[string]any)
			if !ok {
				return ""
			}
			return v[key]
		},
		"html": func(v any) template.HTML {
			if v == nil {
				return ""
			}
			return template.HTML(fmt.Sprint(v))
		},

		"toEUR": func(amount any, currency string) float64 {
			value, ok := number(amount)
			if !ok {
				return 0
			}

			switch currency {
			case "EUR", "":
				return value
			case "RON":
				return value * 0.19
			case "USD":
				return value * 0.86
			case "GBP":
				return value * 1.16
			default:
				// Unknown currency: do not silently invent a rate.
				// Returning 0 makes the problem visible in the UI.
				return 0
			}
		},

		"eur": func(amount any, currency string) string {
			value, ok := number(amount)
			if !ok {
				return ""
			}

			eur := value
			switch currency {
			case "EUR", "":
			case "RON":
				eur = value * 0.19
			case "USD":
				eur = value * 0.86
			case "GBP":
				eur = value * 1.16
			default:
				return fmt.Sprintf("unsupported currency %q", currency)
			}

			return formatEUR(eur)
		}, "percent": func(v any) string {
			value, ok := number(v)
			if !ok {
				return ""
			}
			return fmt.Sprintf("%.2f%%", value)
		},
	}
}

func number(v any) (float64, bool) {
	switch x := v.(type) {
	case int:
		return float64(x), true
	case int64:
		return float64(x), true
	case float64:
		return x, true
	case float32:
		return float64(x), true
	case json.Number:
		f, err := x.Float64()
		return f, err == nil
	case string:
		f, err := strconv.ParseFloat(x, 64)
		return f, err == nil
	default:
		return 0, false
	}
}

func formatEUR(v float64) string {
	rounded := math.Round(v*100) / 100

	if math.Abs(rounded-math.Round(rounded)) < 0.000001 {
		return fmt.Sprintf("€%.0f", rounded)
	}

	return fmt.Sprintf("€%.2f", rounded)
}

func NewServer(cfg ServerConfig) (*Server, error) {
	tmpl, err := template.New("").
		Funcs(templateFuncs()).
		ParseFS(cfg.WebFS, "web/templates/*.html")
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

	expenses, err := parseExpenses(cfg.DataFS)
	if err != nil {
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
		Expenses:  expenses,
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
		"AppName":  s.settings.AppName,
		"Expenses": s.Expenses,
	}

	// TODO: move this into NewServer and cache the parsed template
	if err := s.templates.ExecuteTemplate(w, "index.html", data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func parseExpenses(dataFs fs.FS) ([]Expenses, error) {
	dataDir, err := fs.ReadDir(dataFs, "data")
	if err != nil {
		return nil, err
	}
	var expenses []Expenses
	for _, entry := range dataDir {
		if entry.IsDir() {
			continue
		}
		expensesFile, err := fs.ReadFile(dataFs, "data/"+entry.Name())
		if err != nil {
			return nil, err
		}
		var expense Expenses
		if err := expense.UnmarshalJSON(expensesFile); err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}
	return expenses, nil
}
