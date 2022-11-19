<script lang="ts">
  import { Breadcrumb, BreadcrumbItem } from "flowbite-svelte";
  import { page } from "../store/page";
  import Fibonacci from "./Fibonacci.svelte";
</script>

<main>
  <div class="md:container md:mx-auto">
    <Breadcrumb class="mb-8" aria-label="Default breadcrumb example">
      <BreadcrumbItem>
        <h1>{$page}</h1>
      </BreadcrumbItem>
    </Breadcrumb>
    {#if $page === "fibonacci"}
      <Fibonacci />
    {/if}
  </div>
</main>

<style>
  h1 {
    text-transform: capitalize;
  }
</style>
