<script lang="ts">
	import '../app.css';
	import { Navbar } from '$lib';
	import Footer from '$lib/components/Footer.svelte';

	let { children } = $props();
</script>

<Navbar title="Supynės" />

<div class="content-wrapper">
	{@render children?.()}
</div>

<Footer></Footer>

<style>
	.content-wrapper {
		background-color: rgb(197, 197, 187);
		background-size: cover;
		background-position: center;
		background-repeat: no-repeat;
		min-height: calc(100vh - 64px - 48px);
		width: 100%;
		padding: 20px 0;
	}
</style>
