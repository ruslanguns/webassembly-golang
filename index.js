import "./assets/js/wasm_exec_tiny.js"

const go = new Go();

WebAssembly.instantiateStreaming(fetch("build/go.wasm"), go.importObject)
  .then((result) => {
    go.run(result.instance);
  })