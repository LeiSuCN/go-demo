# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

A minimal Go HTTP server: `GET /` returns `Hello, world!`.

Single module at the repo root, standard library only, no third-party dependencies.

## Build & Run

```bash
# Run the server (listens on :8080)
go run ./app/http
```

## Structure

- `app/` — Application layers, each a self-contained app with its own `main.go`.
  - `app/http/` — The HTTP service.
    - `main.go` — Entry point. Registers routes and serves on `:8080`.
    - `handler/` — HTTP handlers, one file per handler.
- `model/` — Shared data models used across application layers.
