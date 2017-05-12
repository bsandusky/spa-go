# spa-go

SPA-like application written in Go, using built-in templating and no(t very much) Javascript. Something of an experiment based on some previous other work I've done on the same, though in a private repo.

## Goals
- Leverage Go's built-in template compilation to parse & cache static assets on init()
- Take advantage of Go server performance for routing and rendering events
- Severely limit the amount of JS on the front-end specifically to only necessary page-specific interactions
- Build an SPA-like application that can dynamically change based on the URL without an entire JS front-end app with its own runtime
- Explore ways to make template <> BE communications more RESTful, at least in spirit
- Design a lightweight but feasible front-end router system
- Write modular, testable code
