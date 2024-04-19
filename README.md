# Bakersfield Technology Go Rewrite (WIP)

I built [bakersfieldtechnology.com](https://bakersfieldtechnology.com) using Next.js's static export feature so I could host it for free on Netlify. Using a static site prevented me from being able to embed a contact form.

I rebuilt the site–not because I'm seeing contracting work in Bakersfield, but instead so I can get experience with Go, templ and Vite. I also believe that Go apps are lightweight enough that I can probably run a dozen apps on a $5/month DigitalOcean server.

## Developing

You'll need the following on your system to build this app:

- Go ([Installation](https://go.dev/doc/install))
- Templ CLI ([Installation](https://templ.guide/quick-start/installation))
- Air ([Installation](https://github.com/cosmtrek/air))
- Node/NPM ([Installation](https://nodejs.org/en/download))

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

The Templ watcher currently has issues, but once they are resolved the water can be run with the following command:

```sh
templ generate --watch --proxy="http://localhost:3005" --cmd="go run ."
```

## Build tags in Neovim

I'm not sure exactly how to make `gopls` work with Go build tags in Neovim, but maybe when I understand a little bit more about Neovim [this Discourse dicussion](https://neovim.discourse.group/t/gopls-settings-buildflags/790/10) will make more sense.

## Building

To build the project for the local machine, run the following command:

```sh
npm run build && templ generate && go build -o bakersfieldtechnology.com
```

## Deploying

To build the project for the server, run the following command:

```sh
npm run build && templ generate && GOOS=linux GOARCH=amd64 go build -o bakersfieldtechnology.com
```

## Router

I'm using [Echo](https://echo.labstack.com/) for routing on this project because I want to try something I haven't used, Echo's error handling looks simpler than `net/http`-compatible libraries, it has comprehensive docs, and Primeagen uses it in his Go repos 🤷‍♂️

For posterity, I'm including a list of all the routers I evaluated for this project below.

- [Echo](https://echo.labstack.com/) Zero dynamic allocation router. Best docs. Not `net/http` compatible, but has simpler error handling story.
- [Chi](https://go-chi.io/#/). ~1000 lines of code. Compatible with `net/http`. No external dependencies.
- [Gorilla Mux](https://github.com/gorilla/mux). Very popular. Compatible with `net/http`.
- [Gin](https://gin-gonic.com/). Claims to be the fastest router. Zero allocations. Docs leave a lot to be desired.

## Vite

I am following the [Vite Backend Integration guide](https://vitejs.dev/guide/backend-integration.html) to try to use Vite for compiling JS/TS and Scss. However, the project uses Tailwind for the majority of its styling.

## Build tag gopls support

Support for mutually exclusive build tags is not great. Read more [here](https://github.com/golang/go/issues/29202).
