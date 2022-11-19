<script type="ts">
  import {
    Table,
    TableBody,
    TableBodyCell,
    TableBodyRow,
  } from "flowbite-svelte";

  export let title = "";
  export let result = 0;
  export let time = 0;
</script>

<h2 class="text-xl mb-4">{title}</h2>
<Table>
  <TableBody class="divide-y">
    <TableBodyRow>
      <TableBodyCell>Result</TableBodyCell>
      <TableBodyCell
        tdClass="px-6 py-4 whitespace-nowrap font-medium text-right"
      >
        {result}
      </TableBodyCell>
    </TableBodyRow>
    <TableBodyRow>
      <TableBodyCell>Benchmark</TableBodyCell>
      <TableBodyCell
        tdClass="px-6 py-4 whitespace-nowrap font-medium text-right"
      >
        {time} ms
      </TableBodyCell>
    </TableBodyRow>
  </TableBody>
</Table>
