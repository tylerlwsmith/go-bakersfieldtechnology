# Bakersfield Technology Go Rewrite (WIP)

I built [bakersfieldtechnology.com](https://bakersfieldtechnology.com) using Next.js's static export feature so I could host it for free on Netlify. Using a static site prevented me from being able to embed a contact form.

I'm currently rebuilding the site–not because I'm seeing contracting work in Bakersfield, but instead so I can get experience with Go, Templ and Vite. I also believe that Go apps are lightweight enough that I can probably run a dozen apps on a $5/month DigitalOcean server.

## Routers considered

I am currently evaluating the following routers:

- [Gorilla Mux](https://github.com/gorilla/mux)
- [Chi](https://go-chi.io/#/)
- [Gin](https://gin-gonic.com/)
- [Echo](https://echo.labstack.com/)

## Vite

I am following the [Vite Backend Integration guide](https://vitejs.dev/guide/backend-integration.html) to try to use Vite for compiling JS/TS and Scss.
