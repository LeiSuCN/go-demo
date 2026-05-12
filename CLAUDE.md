# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

A Go web application for student information management. Built with the standard library `net/http`, serving static HTML pages and RESTful JSON APIs.

Monorepo with multiple Go modules tied together by `go.work`.

## Build & Run

```bash
# Run the server (listens on :8080)
go run ./apps/api
```

No build scripts or Makefile — `go run` / `go build` are the only commands.

## Architecture

- **`go.work`** — Workspace definition tying together `pkgs/models` and `apps/api` modules.
- **`pkgs/models/`** — Shared module containing the `Student` data model. Imported by all apps that need the student definition.
- **`apps/api/`** — The HTTP API service module.
  - `main.go` — Entry point. Embeds static files from `public/` via `//go:embed`, initializes the in-memory store, registers API routes at `/api/submit-student` and `/api/get-students`.
  - `handler/` — HTTP handlers and business logic.
    - `common.go` — In-memory storage (`map[string]models.Student` protected by `sync.Mutex`) and JSON response helpers.
    - `submit_student_handler.go` — POST handler that validates and persists student data.
    - `get_students_handler.go` — GET handler with optional `?studentId=` query param for single-student lookup; returns all students otherwise.
    - `test_data.go` — Seeds three test students.
  - `public/` — Static HTML/CSS pages with Apple-style UI (gradient backgrounds, sidebar navigation, glassmorphism effects). Embedded at compile time.
