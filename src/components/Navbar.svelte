<script lang="ts">
  import { Navbar, NavBrand, NavLi, NavUl, DarkMode } from "flowbite-svelte";
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
      active={$page === "home"}
      on:click={() => page.set("home")}
    >
      Home
    </NavLi>
    <NavLi
      class="cursor-pointer"
      active={$page === "fibonacci"}
      on:click={() => page.set("fibonacci")}
    >
      Fibonacci
    </NavLi>
    <NavLi
      class="cursor-pointer"
      active={$page === "grayscale"}
      on:click={() => page.set("grayscale")}
    >
      GrayScale
    </NavLi>
    <NavLi class="cursor-pointer">
      <DarkMode
        btnClass="text-gray-500 dark:text-gray-400 hover:bg-gray-100 dark:hover:bg-gray-700 focus:outline-none focus:ring-4 focus:ring-gray-200 dark:focus:ring-gray-700 rounded-lg text-sm"
      />
    </NavLi>
  </NavUl>
</Navbar>
