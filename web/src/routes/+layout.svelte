<script lang="ts">
	import '../app.css';
	import { Navbar } from '$lib';
	import Footer from '$lib/components/Footer.svelte';

	let { children } = $props();
</script>

<Navbar title="Supynės" />

{@render children?.()}

<Footer></Footer>
