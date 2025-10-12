<script lang="ts">
	import { onMount } from 'svelte';
	export let title = 'User Profile - Supynės';
	onMount(() => {
		document.title = title;
	});
</script>
