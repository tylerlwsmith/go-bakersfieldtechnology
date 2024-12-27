# Bakersfield Technology Go Rewrite

> This project has been superseded by a new Astro-powered version of the site. You can view the new repo [here](https://github.com/tylerlwsmith/bakersfield-technology-astro).

I originally built the 2-page [bakersfieldtechnology.com](https://bakersfieldtechnology.com) website in 2021 using Next.js. Years later when I wanted to learn [Templ](https://templ.guide/) and the [Echo framework](https://echo.labstack.com/), rebuilding the website in Go seemed like a great learning opportunity.

On the frontend, the site uses Tailwind & Scss for styles, TypeScript for interactivity, and Vite for bundling. When running a production build, the assets are compiled into the binary using `//go:embed` tags. The work I did building the Vite integration on this app also spun out into its own [blog post](https://dev.to/tylerlwsmith/build-a-vite-5-backend-integration-with-flask-jch).

You can view the repo for the original Next.js-powered site [here](https://github.com/tylerlwsmith/bakersfieldtechnology.com).

## Developing

You'll need the following on your system to build this app:

- Go ([Installation guide](https://go.dev/doc/install))
- Templ CLI ([Installation guide](https://templ.guide/quick-start/installation))
- Air ([Installation guide](https://github.com/cosmtrek/air))
- Node/NPM ([Installation guide](https://nodejs.org/en/download))

After cloning the repo, run the following commands:

```sh
go mod tidy
npm install
```

To build the assets during development, run the following command:

```sh
npm run dev
```

You can start the Go app with live reloading using the following command:

```sh
air
```

The Templ watcher has issues at the time of development, but once they are resolved the watcher can be run with the following command:

```sh
templ generate --watch --proxy="http://localhost:3005" --cmd="go run ."
```

## Building locally

To build the project for the local machine, run the following command:

```sh
npm run build && templ generate && go build -o bakersfieldtechnology.com
```

## Building for production

To build the project for the production x86 Linux server, run the following command:

```sh
npm run build && templ generate && GOOS=linux GOARCH=amd64 go build -o bakersfieldtechnology.com
```

You'll find a unit file and Caddyfile for the server in the `deployment/` directory of this project. The unit file will require some modification to tailor it to the particular server.

## Router

I'm using [Echo](https://echo.labstack.com/) for routing on this project because I want to try something I haven't used, Echo's error handling looks simpler than `net/http`-compatible libraries, it has comprehensive docs, and Primeagen uses it in his Go repos 🤷‍♂️

For posterity, I'm including a list of all the routers I evaluated for this project below.

- [Echo](https://echo.labstack.com/) Zero dynamic allocation router. Best docs. Not `net/http` compatible, but has simpler error handling story.
- [Chi](https://go-chi.io/#/). ~1000 lines of code. Compatible with `net/http`. No external dependencies.
- [Gorilla Mux](https://github.com/gorilla/mux). Very popular. Compatible with `net/http`.
- [Gin](https://gin-gonic.com/). Claims to be the fastest router. Zero allocations. Docs leave a lot to be desired.

## Build tag gopls support

The `gopls` support for mutually exclusive build tags isn't great. Read more [here](https://github.com/golang/go/issues/29202).

## Build tags in Neovim

I'm not sure exactly how to make `gopls` work with Go build tags in Neovim, but maybe when I understand a little bit more about Neovim [this Discourse dicussion](https://neovim.discourse.group/t/gopls-settings-buildflags/790/10) will make more sense.
