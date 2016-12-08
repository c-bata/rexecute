# rexecute

Re execute the command when specified files are updated.

## Installation

```console
$ go install github.com/c-bata/rexecute
```

## Design

```console
$ find . -name "*.tex" | rexecute make build
```

![rexecute demo](./rexecute-demo.gif)

## License

MIT License@Masashi Shibata

