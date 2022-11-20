import "./app.postcss";
import App from "./App.svelte";

document.addEventListener("DOMContentLoaded", (event) => {
  const wasmUrl = new URL("/main.wasm", import.meta.url).href;
  // @ts-ignore
  const go = new Go();
  WebAssembly.instantiateStreaming(fetch(wasmUrl), go.importObject).then(
    (result) => {
      go.run(result.instance);
    }
  );
});

const app = new App({
  target: document.getElementById("app"),
});

export default app;
