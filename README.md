# Disdy DOM 
## - What is ?
> Disdy DOM makes life easier when working with the DOM and other things.Disdy DOM is in a very early stage, its use is not recommended if you are not crazy XD

### How to compile ?
- using `go`:
```bash
GOARCH=wasm GOOS=js go build -o <name of output file>.wasm <main name file>.go
```
- using `tinygo`
```bash
tinygo build -o <name of output file>.wasm -target wasm <main name file>.go
```

#### Go or Tinygo ???

lite binary: `Tinygo`

full and oficial: `Go`
`They are both very good`

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)