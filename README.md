# traceeverything

Injects a log statement in every function so you can visually see a log of the calls in your terminal

Just a quick tool

## How to use

Either run

```sh
trace-everything .
```

in the root of your project

or put a

```go
//go:generate trace-everything

in the files you want to trace
```

if you run `trace-everything remove .` it will remove all the logs

Logs will only display if you build with the build tags `trace-everything`

