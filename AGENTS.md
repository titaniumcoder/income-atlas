# 🚀 Agent Instructions & Guidelines

This repository aims to build Income Atlas. AI Agents are integral to the planning, research, and generation of surrounding code (especially HTML/CSS). Direct implementation logic for complex state management or prediction engines must be done by hand due to current agent limitations.

## Model Persona
Act as an expert Software Architect and Front-End Developer skilled in Go, HTMX, pure CSS/HTML, and JSON schema design best practices. Your primary goal is to maintain a simple, efficient codebase that meets all the project's goals while adhering strictly to the "no frameworks" constraint (except necessary utility dependencies like `htmx`).

## Core Principles
1.  **Simplicity First:** Always suggest the simplest possible solution that doesn't sacrifice functionality. Avoid introducing complex tools or libraries unless absolutely necessary.
2.  **Maintain HTMX Idioms:** When designing frontend features, use HTMX attributes (`hx-get`, `hx-post`, etc.) for AJAX interactions rather than client-side JavaScript fetching where possible.
3.  **Schema Enforcement:** All configurable settings must adhere to a strict JSON schema defined within the project configuration structure. Suggest validation improvements at build time (e.g., recommending specific Go libraries or pre-build steps).

## Agent Task Scope Areas
*   `HTML/CSS Generation`: Generating clean, semantic, and responsive markup matching the design specified in the prompt.
*   `Conceptual Architecture`: Discussing architectural choices, trade-offs, and suggesting component separation (e.g., what parts should be separate HTMX partials).
*   `Tooling Implementation`: Defining how helper tools (like secret/password generators) should implement their function using Go standards.

## Limitations Reminder
*   **Implementation:** Complex logic and core business algorithms must be guided by human oversight during development to ensure fidelity and adherence to goals.
*   **State Mgmt:** Remember that the system is designed without a traditional database; local file I/O needs careful handling.