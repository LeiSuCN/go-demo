# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

A minimal Go HTTP server: `GET /` returns `Hello, world!`.

Single module at the repo root, standard library only, no third-party dependencies.

## Build & Run

```bash
# Run the server (listens on :8080)
go run .
```

## Structure

- `main.go` — Entry point. Registers the `hello` handler at `/` and serves on `:8080`.
