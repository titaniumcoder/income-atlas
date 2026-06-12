# 🌐 Income Atlas

Income Atlas is a project designed to track, predict income streams based on specific entries (like toggl entries), and compare living conditions across different countries using key metrics such as taxes, social contributions, and cost of living.

## ✨ Goal
To create a simple, powerful web application that provides comprehensive financial tracking and comparative analysis of global economic data. We aim for an intuitive interface built with pure HTML/CSS enhanced by HTMX for modern interactivity.

## 🛠️ Technology Stack
*   **Backend:** Go
*   **Frontend:** Pure HTML + CSS / JavaScript (with heavy reliance on HTMX for inter-component communication)
*   **Database:** None (File-based configuration with JSON Schema validation)

## 📁 Project Structure
```
income-atlas/
├── config/           # Configuration files (JSON)
├── schema/           # JSON Schema definitions
├── web/
│   ├── templates/    # HTML templates
│   └── static/       # Static assets (CSS, JS)
├── data/             # User data (later)
├── main.go           # Application entry point
├── server.go         # HTTP server and routes
├── embed.go          # Go embed directives
└── generated.go      # Auto-generated code from JSON schema
```

## 🧩 Core Features Planned
1.  Income Tracking and Prediction: Inputting data to predict future financial standing.
2.  Global Comparison Tool: Comparing countries based on tax burdens/social contributions and living standards.

## ⚙️ Development Setup
Key development tools include `air` for hot-reloading during development, ensuring the server restarts automatically when files change (configured in `.air.toml`).

### Getting Started
1.  Clone the repository.
2.  Install dependencies: `go mod tidy`
3.  Run the server: `make run` or `air`

### VSCode Integration
The project includes `.vscode/settings.json` to link JSON schema definitions with config files for validation and autocomplete.

## 🗺️ Future State
*   **Live Demo:** https://income-atlas.fly.dev