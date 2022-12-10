## WebAssembly using Go in the Browser

This is a proof of concept where I'm testing golang in the browser.

## Requirements

To build your Go code you could you tinygo with the following command:

```bash
tinygo build -o main.wasm -target=wasm main.go
```

You can either use Docker to generate this build:

```bash
docker run --rm -v $(pwd):/src tinygo/tinygo:latest cp $(tinygo env TINYGOROOT)/targets/wasm_exec.js wasm_exec_tiny.js
docker run --rm -v $(pwd):/src tinygo/tinygo:latest tinygo build -o /src/build/go.wasm -target=wasm /src/cmd/main.go
```

Once it's generated you can import using the Webassembly API.
