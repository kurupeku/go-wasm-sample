<script lang="ts">
  import { Button, Fileupload, Label, Img } from "flowbite-svelte";
  import { toBase64, grayScale } from "../composables/gray";

  let disabled = true;
  let file: File;
  let base64 = "";
  let loading = false;

  async function handleFile(e: Event) {
    const input = e.target as HTMLInputElement;
    file = input.files[0];
    disabled = !file;
    if (file) base64 = await toBase64(file);
  }

  async function handleSubmit() {
    if (!file) {
      base64 = "";
      return;
    }

    loading = true;
    base64 = await grayScale(base64);
    loading = false;
  }
</script>

<div class="mb-8">
  <Label class="pb-2">Choose Image (Only <code>.png</code>)</Label>
  <Fileupload on:change={handleFile} accept=".png" />
  <Button {disabled} on:click={handleSubmit}>
    {#if loading}
      Converting ...
    {:else}
      Convert
    {/if}
  </Button>
</div>
<div class="py-10">
  {#if base64}
    <Img src={base64} alt="result" size="max-w-xs" />
  {/if}
</div>
