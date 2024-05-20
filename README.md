# beecs-template

A template repository showing how to extend the [beecs](https://github.com/mlange-42/beecs) honeybee model while using its CLI app [beecs-cli](https://github.com/mlange-42/beecs-cli).

## Usage

Try the example:

```
go run . -d _examples/base --observers --systems
```

## Getting started

Take a look at file [`main.go`](https://github.com/mlange-42/beecs-template/blob/main/main.go)
to get an idea how it works.

## Dependencies on Linux systems

As [beecs-cli](https://github.com/mlange-42/beecs-cli) provides live plotting using OpenGL, the dependencies of [go-gl/gl](https://github.com/go-gl/gl) and [go-gl/glfw](https://github.com/go-gl/glfw) apply. For Ubuntu/Debian-based systems, these are:

- `libgl1-mesa-dev`
- `xorg-dev`
