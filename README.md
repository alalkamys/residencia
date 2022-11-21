<!-- markdownlint-configure-file {
  "MD033": false,
  "MD041": false,
  "MD051": false
} -->

<div align="center">

<img src="docs/assets/imgs/golden-leaf.png" alt="Residència logo" title="Residència" align="center" height="113"/>

<h1>Residència</h1>

[Key Features](#key-features) •
[Methodology](#methodology) •
[Getting Started](#getting-started) •
[Developed By](#⚔️-developed-by) •
[Author](#book-author)

<img src="https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg" alt="Awesome Badge"/>
<img src="https://img.shields.io/pypi/status/ansicolortags.svg" alt="Stars Badge"/>

</div>

<br />

Welcome to `residencia` repository!

`residencia` is a [Golang][golang] booking web application repository used to manage bookings and reservations.

## Key Features

- Simple to use.
- Follows best practices.
- High maintainabilty.
- Highly flexible.

## Methodology

### 1. Tools and Dependencies

This solution is using:

- Built-in Go version 1.19.
- Routing using [go-chi/chi][pkg-go-chi].
- Session management using [alexedwards/scs/v2][pkg-alexedwards-scs].
- CSRF protection using [justinas/nosurf][pkg-justinas-nosurf].

### 2. Folder Structure

- Folder structure:

  <pre>
  .
  ├── .gitignore
  ├── README.md
  ├── cmd/
  │   └── web/
  │       ├── main.go                   // entrypoint
  │       ├── middleware.go
  │       └── routes.go
  ├── docs/
  ├── go.mod
  ├── go.sum
  ├── pkg/
  │   ├── config/
  │   ├── handlers/
  │   ├── models/
  │   └── render/
  └── templates/
  </pre>

## The Why?

- I wanted to practice my Golang development skills to advance in DevOps field furthermore.

- It would be great to make use of `residencia` in applying DevOps principles by implementing various architectures and tools with `residencia` as it's backend application.

- It's fun!

----

## Getting Started

TO BE ADDED

## ⚔️ Developed By

<a href="https://www.linkedin.com/in/shehab-el-deen/" target="_blank"><img alt="LinkedIn" align="right" title="LinkedIn" height="24" width="24" src="docs/assets/imgs/linkedin.png"></a>

Shehab El-Deen Alalkamy

## :book: Author

Shehab El-Deen Alalkamy

<!--*********************  R E F E R E N C E S  *********************-->

<!-- * Links * -->
[golang]: https://go.dev/
[pkg-go-chi]: https://github.com/go-chi/chi#readme
[pkg-alexedwards-scs]: https://github.com/alexedwards/scs/#readme
[pkg-justinas-nosurf]: https://github.com/justinas/nosurf#readme

<!-- * Images * -->
