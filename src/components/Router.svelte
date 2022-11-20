<script lang="ts">
  import { Breadcrumb, BreadcrumbItem } from "flowbite-svelte";
  import { page } from "../store/page";
  import Home from "./Home.svelte";
  import Fibonacci from "./Fibonacci.svelte";
  import GrayScale from "./GrayScale.svelte";
</script>

<main>
  <div class="md:container md:mx-auto pt-4">
    <Breadcrumb class="mb-8" aria-label="Default breadcrumb example">
      <BreadcrumbItem>
        <h1>{$page}</h1>
      </BreadcrumbItem>
    </Breadcrumb>
    {#if $page === "fibonacci"}
      <Fibonacci />
    {:else if $page === "grayscale"}
      <GrayScale />
    {:else}
      <Home />
    {/if}
  </div>
</main>

<style>
  h1 {
    text-transform: capitalize;
  }
</style>
