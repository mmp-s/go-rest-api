package main

import (
  "path/filepath"
  "testing"
  "github.com/matryer/silk/runner"
)

func TestAPIEndpoints(t *testing.T) {
  startServer()
  runner.New(t, "http://127.0.0.1:8080").RunGlob(filepath.Glob("./docs/*.silk.md"))
}
