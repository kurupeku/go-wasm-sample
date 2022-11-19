<script lang="ts">
  import {
    Button,
    ButtonGroup,
    NumberInput,
    Label,
    Spinner,
    Toggle,
  } from "flowbite-svelte";
  import { fibonacci, fibonacciMemorized } from "../composables/calc";
  import Result from "./FibonacciResult.svelte";

  interface BenchResult {
    result: number;
    time: number;
  }

  let value = 0;
  let memorized = false;
  let goLoading = false;
  let goResult = 0;
  let goTime = 0;
  let jsLoading = false;
  let jsResult = 0;
  let jsTime = 0;

  $: value > maxValue && (value = maxValue);
  $: maxValue = memorized ? 80 : 40;

  function benchmark(fn: (v: number) => number): BenchResult {
    const s = new Date().getTime();
    const result = fn(value);
    const time = new Date().getTime() - s;

    return { result, time };
  }

  function goExec(): Promise<BenchResult> {
    return new Promise((resolve) => {
      resolve(
        benchmark(memorized ? window.fibonacciMemorized : window.fibonacci)
      );
    });
  }

  function jsExec(): Promise<BenchResult> {
    return new Promise((resolve) => {
      resolve(benchmark(memorized ? fibonacciMemorized : fibonacci));
    });
  }

  function handleSubmit() {
    goLoading = true;
    goExec().then(({ result, time }) => {
      goResult = result;
      goTime = time;
      goLoading = false;
    });
    jsLoading = true;
    jsExec().then(({ result, time }) => {
      jsResult = result;
      jsTime = time;
      jsLoading = false;
    });
  }

  function handleMax() {
    value = maxValue;
  }
</script>

<div class="mb-8">
  <Label class="mb-2">Index Number (Max: {maxValue})</Label>
  <ButtonGroup size="sm">
    <NumberInput bind:value max={maxValue} />
    <Button on:click={handleMax}>Max</Button>
    <Button on:click={handleSubmit}>Submit</Button>
  </ButtonGroup>
  <Toggle class="mt-4" bind:checked={memorized}>Use Cache</Toggle>
</div>
<div class="flex">
  <div class="flex-auto w-1/2 px-4">
    {#if goLoading}
      <Spinner />
    {:else}
      <Result title="WASM (Golang)" result={goResult} time={goTime} />
    {/if}
  </div>
  <div class="flex-auto w-1/2 px-4">
    {#if jsLoading}
      <Spinner />
    {:else}
      <Result title="JavaScript" result={jsResult} time={jsTime} />
    {/if}
  </div>
</div>
