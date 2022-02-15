# E-commerce online platform project with Golang

- [E-commerce Project](#e-commerce-project)
  - [About the project](#about-the-project)
    - [API docs](#api-docs)
    - [Status](#status)
  - [Getting started](#getting-started)
    - [Layout](#layout)
  - [Notes](#notes)
  - [Instructions](#instructions)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# E-commerce Project

## About the project

This is an e-commerce project, allowing to make operations connected to the platform user (Create/Update/Delete), and lets the user to make purchases with 
individual or subscription plans, checkout via a card and generate purchased item invoice PDF's.

### API docs

The project doesn't have API docs. It will be added shortly with the Documentation of a swaggerfile.

### Status

This is a test project with the beta version.

## Getting started

Below is described the conventions or tools specific to this project.

### Layout

```tree
├── air_api.toml
├── air_invoice.toml
├── air_web.toml
├── cmd
│   ├── api
│   │   ├── api.go
│   │   ├── handlers-api.go
│   │   ├── helpers.go
│   │   ├── mailer.go
│   │   ├── middleware.go
│   │   ├── routes.api.go
│   │   └── templates
│   │       ├── password-reset.html.tmpl
│   │       └── password-reset.plain.tmpl
│   ├── micro
│   │   └── invoice
│   │       ├── email-templates
│   │       │   ├── invoice.html.tmpl
│   │       │   └── invoice.plain.tmpl
│   │       ├── invoice.go
│   │       ├── invoice-handlers.go
│   │       ├── invoice-helpers.go
│   │       ├── invoice-mailer.go
│   │       └── invoice-routes.go
│   └── web
│       ├── handlers.go
│       ├── home.page.gohtml
│       ├── main.go
│       ├── middleware.go
│       ├── render.go
│       ├── routes.go
│       ├── templates
│       │   ├── all-sales.page.gohtml
│       │   ├── all-subscriptions.page.gohtml
│       │   ├── all-users.page.gohtml
│       │   ├── base.layout.gohtml
│       │   ├── bronze-plan.page.gohtml
│       │   ├── buy-once.page.gohtml
│       │   ├── forgot-password.page.gohtml
│       │   ├── home.page.gohtml
│       │   ├── login.page.gohtml
│       │   ├── one-user.page.gohtml
│       │   ├── receipt.page.gohtml
│       │   ├── receipt-plan.page.gohtml
│       │   ├── reset-password.page.gohtml
│       │   ├── sale.page.gohtml
│       │   ├── stripe-js.partial.gohtml
│       │   ├── terminal.page.gohtml
│       │   └── virtual-terminal-receipt.page.gohtml
│       ├── tmp
│       │   └── build-errors.log
│       └── ws-handlers.go
├── database.yml
├── dist
│   └── invoice
├── docker-compose.yaml
├── Dockerfile
├── Dockerfile_api
├── README.md
├── .env.test
├── go.mod
├── go.sum
├── internal
│   ├── cards
│   │   └── cards.go
│   ├── driver
│   │   └── driver.go
│   ├── encryption
│   │   └── encryption.go
│   ├── models
│   │   ├── models.go
│   │   └── tokens.go
│   ├── urlsigner
│   │   └── signer.go
│   └── validator
│       └── validator.go
├── invoices [error opening dir]
├── Makefile
├── migrations
│   └── schema.sql
├── pdf-templates
│   └── invoice.pdf
├── static
│   └── widget.png
└── tmp [error opening dir]
```

A brief description of the layout:

* `.gitignore` is scecifically for the golang code and also does not include /tmp and /invoices folder belongings.
* `Makefile` is not used to build the project. **Makefile is made for the testing purposes**.
* `README.md` is a detailed description of the project.
* `cmd` contains main packages, with subdirectories api/, micro/, web/.
* `migrations` contains script for sql schema and basic migrations regarding to the project table details.
* `api` contains main backend project functionality.
* `micro`package contains independent service for genereting PDF files and sending them via email.
* `web` contains FE part of the application and the main way to directly connect to the BE services
* `internal` holds all the models and model specific functions insede the app.

## Notes

* Makefile is not currently used to start a project.
* Project uses cosmtrek/air live reloader for golang apps
* Dockerfile is used to run the project on its specific ports:
  * BackEnd :4000 
  * FrontEnd :4001
  * MariaDB :3306
  * InvoiceService :5000
  

## Instructions

* To run the project locally you need to pull the directory code, fill the .env file and run the command `docker-compose up`

