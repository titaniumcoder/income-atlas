package main

import "embed"

//go:embed web/templates/*.html web/static/*
var webFS embed.FS

//go:embed config/*.json
var configFS embed.FS
