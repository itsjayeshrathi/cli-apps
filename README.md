# Go CLI Apps – Learning by Building

This repository contains a collection of command-line applications built using Go, as part of my learning journey inspired by the book **“Powerful Command-Line Applications in Go”**.

Each folder contains a standalone Go CLI tool with its own `go.mod`, source code, and tests.

## 📂 Projects

| Project     | Description                                                  |
|-------------|--------------------------------------------------------------|
| [`mdp-cli`](./mdp-cli)   | A markdown previewer that renders `.md` files to HTML.          |
| [`todo-cli`](./todo-cli) | A simple todo list manager using file-based storage.            |
| [`walk-cli`](./walk-cli) | Walks a directory tree and filters/list/deletes files based on extension, size, etc. |
| [`wc-cli`](./wc-cli)     | A basic word count utility similar to the Unix `wc` command.    |
| [`colstat-cli`](./colstat-cli)| Simple tool for analyzing csvs in order to understand benchmarking and profiling |
| [`unarchive-cli`](./unarchive-cli)| A tool for unarchiving zipped files at particular location. |
| [`goci-cli`](./goci-cli)| Similar to ci-cd tool for golang projects. |

> ✅ Each project includes tests and is written using idiomatic Go.

## 🔧 Requirements

- Go 1.20+
- (Optional) `make`, `bat`, or `just` if you want to script CLI usage

## 🚀 Getting Started

Clone the repo and `cd` into any project directory to run it:

```bash
git clone https://github.com/yourusername/cli-apps.git
cd cli-apps/mdp-cli
go run main.go testdata/test.md
