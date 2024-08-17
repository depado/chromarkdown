# Chromarkdown

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)[![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com)

![Go Version](https://img.shields.io/badge/Go%20Version-latest-brightgreen.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/depado/chromarkdown)](https://goreportcard.com/report/github.com/depado/chromarkdown)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/depado/chromarkdown/blob/master/LICENSE)
[![Godoc](https://godoc.org/github.com/depado/chromarkdown?status.svg)](https://godoc.org/github.com/depado/chromarkdown)
[![Say Thanks!](https://img.shields.io/badge/Say%20Thanks-!-1EAEDB.svg)](https://saythanks.io/to/Depado)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fdepado%2Fchromarkdown.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fdepado%2Fchromarkdown?ref=badge_shield)

Build single-file static HTML page with Chroma (syntax highlighting) and Markdown

![screenshot](https://github.com/depado/chromarkdown/blob/master/img/screenshot.png)

## Introduction

Chromarkdown is a tool to generate single-file static HTML pages from a
Markdown file input. This tool has no external dependecies and performs the
following operations:

- Syntax highlighting thanks to [chroma](https://github.com/alecthomas/chroma)
- Markdown rendering using [blackfriday](https://github.com/russross/blackfriday)
- Embedded [Roboto](https://fonts.google.com/specimen/Roboto) and
[Roboto-mono](https://fonts.google.com/specimen/Roboto+Mono) fonts
- Dynamic CSS for Syntax Highlighter according to the chosen theme
- Single-file (one HTML file) output with embedded styles and fonts (no network
call)
- Responsive page with simple design according to
[bettermotherfuckingwebsite](http://bettermotherfuckingwebsite.com/)

## Install

Pre-compiled binaries can be found in the [releases](https://github.com/depado/chromarkdown/releases)
tab.

```
$ wget https://github.com/depado/chromarkdown/releases/download/v1.0.0/chromarkdown_linux_amd64
```

## Usage

Once installed you can run the `chromarkdown` command:

```
$ chromarkdown --help

Chromarkdown uses a combination of blackfriday and chroma to render an input markdown file.
It generates standalone HTML files that includes fonts, a grid system and extra CSS.

Usage:
  chromarkdown [input file] [flags]

Flags:
  -h, --help            help for chromarkdown
      --no-toc          Disable the table of content
  -o, --output string   specify the path of the output HTML (default "out.html")
      --theme string    Specify the theme for syntax highlighting (default "monokai")
  -t, --title string    Specify the title of the HTML page (default "Ouput")
```


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fdepado%2Fchromarkdown.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fdepado%2Fchromarkdown?ref=badge_large)
