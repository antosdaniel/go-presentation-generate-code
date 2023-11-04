# Go Presentation: generate code

Small gRPC service presenting concepts of code generation. We are generating ORM, decorators, gRPC client and server.

[Slides](https://slides.com/antosdaniel/generating-code-in-go)

## Quick start

To start server you will need Docker installed. Once you have it, just run:
```
make dev
```

If you are using JetBrains IDE, you can execute some example calls from `dev/example gRPC calls.http`.

## Abstract

When I started my adventure with Go, I couldn't stand generated mocks. Coming from a more dynamic language, I thought that I can recapture that "magic" using Go's reflection... Here I am, 5 years later, advocating for the opposite!

One of the first "Aha!" moments was using gowrap to generate decorators. I believe that good observability in the system primarily requires good discipline. This is where generating decorators, that bring traces, logs and metrics shines. Developers can focus on generating business value, while telemetry will come from coding to the interface, with go:generate one line above. We haven't had issues with missing traces ever since.

The second big leap in productivity came from generating code based on schema. We are using GraphQL and gRPC. Especially for GraphQL, writing all required scaffolding was quite a chore. This often lead to some obscure bugs with types, if we didn't check everything during runtime. Once we changed to gqlgen - a schema-first Go generator for GraphQL - all of that changed. Not only we squashed many bugs right after migrating, because compiler showed all the issues to us, but collaboration between backend and frontend improved. When you put schema first, and it is your source of truth, you can discuss it with your coworkers sooner, and iterate on implementation, or improvement, quicker.

The last part (for now) is about database. While schema-first approach also contributed to change - after all, for SQL databases you probably already have your table definitions, and migrations live in your codebase - undetected breaking changes, and high complexity of reflection based ORMs is what pushed us to sqlboiler. This library generates ORM for you, based on a database you already have. If you never had to go through reflection-based ORM, you might be surprised, that understanding what happens under a simple "Find by ID" query can take you a day. This is something we all hit, once the magic of ORM fails us, and we have to dig into what went wrong. Generated code doesn't have that issue. Computers do not mind repeating themselves, and will not make a mistake when writing the same line of code again, and again, and again. The result is a simple "Find" function, that has 10 lines, and has SQL query you would have written yourself in database console. Oh, and breaking changes? They will be plain as day in pull request.
