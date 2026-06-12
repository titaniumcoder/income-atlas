# 🌐 Income Atlas

Income Atlas is a project designed to track, predict income streams based on specific entries (like toggl entries), and compare living conditions across different countries using key metrics such as taxes, social contributions, and cost of living.

## ✨ Goal
To create a simple, powerful web application that provides comprehensive financial tracking and comparative analysis of global economic data. We aim for an intuitive interface built with pure HTML/CSS enhanced by HTMX for modern interactivity.

## 🛠️ Technology Stack
*   **Backend:** Go
*   **Frontend:** Pure HTML + CSS / JavaScript (with heavy reliance on HTMX for inter-component communication)
*   **Database:** None (Local files or in-memory state preferred, keeping the structure minimal).

## 🧩 Core Features Planned
1.  Income Tracking and Prediction: Inputting data to predict future financial standing.
2.  Global Comparison Tool: Comparing countries based on tax burdens/social contributions and living standards.

## ⚙️ Development Setup
Key development tools include `air` for hot-reloading during development, ensuring the server restarts automatically when global settings change. Client-side validation relies heavily on JSON Schema.

### Getting Started
1.  Clone the repository.
2.  Run the local development command (TBD).
3.  Manage global settings via the `settings/` directory, ensuring schema validation is utilized.

## 🗺️ Future State
*   **Live Demo:** https://income-atlas.fly.dev
*   The project structure prioritizes simplicity: a single startup command and small helper utilities for tasks like generating session secrets or passwords.