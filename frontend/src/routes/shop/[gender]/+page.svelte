<script lang="ts">
	import Nav from '../../../components/Nav.svelte';
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();
	let current = $state('XX');
</script>

<Nav />
<!-- <h1>{data.slug}'s clothing</h1> -->

<div class="current"><img src={current} alt="beep bop"/></div>

<div class="clothes">
	{#each data.items as item}
		<!-- svelte-ignore a11y_no_static_element_interactions -->
		<div class="item" onmouseenter={() => (current = item.image)}>
			<img src={item.image} alt={item.name} loading="lazy" />
			<h1>{item.name}</h1>
		</div>
	{/each}
</div>

<style>
	.clothes {
		display: flex;
		flex-wrap: wrap;
		justify-content: space-between;
		margin: 2em 1em;
	}

	.item {
		width: 200px;
		margin: 0.25em;
		background-color: lightgray;
		border: 1px solid white;
	}

	.item:hover {
		border: 1px solid grey;
		cursor: pointer;
	}

	img {
		width: 100px;
	}

	h1 {
		font-size: 1em;
		margin: 0;
		padding: 0;
	}

	.current {
		margin: 1em auto;
		width: 400px;
		text-align: center;
	}
</style>
