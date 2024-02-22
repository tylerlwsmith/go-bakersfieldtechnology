# Bakersfield Technology Go Rewrite (WIP)

I built [bakersfieldtechnology.com](https://bakersfieldtechnology.com) using Next.js's static export feature so I could host it for free on Netlify. Using a static site prevented me from being able to embed a contact form.

I'm currently rebuilding the site–not because I'm seeing contracting work in Bakersfield, but instead so I can get experience with Go, templ and Vite. I also believe that Go apps are lightweight enough that I can probably run a dozen apps on a $5/month DigitalOcean server.

## Developing

This repo is currently using [Air](https://github.com/cosmtrek/air) for live reloading. You can start this app with the following command:

```sh
air
```

The templ watcher currently has issues, but once they are resolved the water can be run with the following command:

```sh
templ generate --watch --proxy="http://localhost:3005" --cmd="go run ."
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
