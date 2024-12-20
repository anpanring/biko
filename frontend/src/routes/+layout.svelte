<script>
	import { page } from '$app/stores';
	import Nav from '../components/Nav.svelte';
	import '../globals.css';
</script>

<head>
	<title>Biko Bostadinov</title>
</head>

<div id="visualGrid" class="visualGrid"></div>

<Nav />

<main>
	{#key $page.url.pathname}
		<slot />
	{/key}
</main>

<style>
	.visualGrid {
		position: fixed;
		width: calc(100vw - 20px);
		min-height: calc(100vh - 20px);
		top: 10px;
		left: 10px;
		background-image: radial-gradient(circle at 1px 1px, #000 1px, transparent 0px);
		background-size: calc((100vw - 20px) / 32) calc((100vw - 10px) / 32);
		pointer-events: none;
		z-index: -1;
		background-repeat: repeat;
	}
</style>
