<script lang="ts">
	import '../app.css';
	import { Navbar } from '$lib';

	let { children } = $props();
</script>

<Navbar title="Supynės" />

{@render children?.()}
