<script lang="ts">
	import { base } from '$app/paths';
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();
	console.log(data);
	let current = $state(data.items[0].image);
</script>

<!-- <h1>{data.slug}'s clothing</h1> -->

<div class="container">
	<div class="current"><img src={current} alt="beep bop" /></div>

	<div class="clothes">
		{#each data.items as item}
			<!-- svelte-ignore a11y_no_static_element_interactions -->
			<div class="item" onmouseenter={() => (current = item.image)}>
				<a href="{base}/shop/{data.slug}/{item.id}">
					<img src={item.image} alt={item.name} loading="lazy" />
					<h1>{item.name}</h1>
				</a>
			</div>
		{/each}
	</div>
</div>

<style>
	.container {
		display: grid;
		grid-template-columns: var(--grid-cols);
		grid-template-rows: var(--grid-rows);
		height: calc(100dvh - 20px);
		--current-img-width: 300px;
	}

	.clothes {
		grid-column-start: 12;
		grid-column-end: calc(var(--grid-max) + 1);
		grid-row-start: 3;
		column-gap: 10px;
		display: flex;
		flex-wrap: wrap;
		height: 100vh;
	}

	.item {
		margin-bottom: var(--grid-module);
		border: 1px solid white;
		animation: fadeIn 0.25s linear;
		box-sizing: border-box;
		flex: 0 0 calc(var(--grid-module) * 3);
	}

	@keyframes fadeIn {
		0% {
			opacity: 0;
		}

		100% {
			opacity: 1;
		}
	}

	.item:hover {
		border: 1px solid grey;
		cursor: pointer;
	}

	img {
		width: 100%;
	}

	h1 {
		font-size: 13px;
		margin: 0;
		padding: 0;
	}

	.current {
		animation: fadeIn 0.25s linear;
		grid-column-start: 1;
		grid-column-end: 11;
		grid-row-start: 3;
	}

	.current img {
		position: fixed;
		width: calc(var(--grid-module) * 10 + 3px);
	}
</style>
