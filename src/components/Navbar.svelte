<script lang="ts">
  import { Navbar, NavBrand, NavLi, NavUl } from "flowbite-svelte";
  import { page } from "../store/page";
</script>

<Navbar>
  <NavBrand>
    <span
      class="self-center whitespace-nowrap text-xl font-semibold dark:text-white"
    >
      Go WASM Sample
    </span>
  </NavBrand>
  <NavUl>
    <NavLi
      class="cursor-pointer"
      active={$page === "index"}
      on:click={() => page.set("index")}>Home</NavLi
    >
    <NavLi
      class="cursor-pointer"
      active={$page === "fibonacci"}
      on:click={() => page.set("fibonacci")}>Fibonacci</NavLi
    >
    <NavLi
      class="cursor-pointer"
      active={$page === "scraping"}
      on:click={() => page.set("scraping")}>Scraping</NavLi
    >
  </NavUl>
</Navbar>
