# Catball Url Shortener

This is an url shortener and doc hosting tool written in go, with minimal footprint.  
It is designed to run in stand-alone mode, or, you can copy the binary and run it at anywhere.

## Features

- Shorten URL
- Host rendered markdown document

## Usage

There's no front end for catball, but it provides a simple api.

1. Copy `config.example.json` to `example.json` and change the information within.
2. run `export GIN_MODE=release`
3. ./catball

## APIs

See main.go
