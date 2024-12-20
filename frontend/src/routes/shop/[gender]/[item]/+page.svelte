<script lang="ts">
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();
</script>

<div class="item-page-container">
	<h1>{data.gender} &gt; {data.item}</h1>
</div>

<style>
	.item-page-container {
		margin-top: calc(var(--grid-module) * 2);
	}
</style>