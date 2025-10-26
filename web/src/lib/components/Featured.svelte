<script lang="ts">
	import { onMount } from 'svelte';
	import SwingCard from './SwingCard.svelte';

	type Tag = {
		Name: string;
	};
	type Swing = {
		ID: number;
		Name: string;
		Address: string;
		City: string;
		Tags: Tag[];
	};
	let swings: Swing[] = [];

	interface FeaturedSwing {
		id: string;
		name: string;
		swing: string;
		location: string;
		rating: number;
		reviewCount: number;
		image: string;
		tags: string[];
	}

	let featuredSwings: FeaturedSwing[] = [];

	onMount(async () => {
		try {
			const resp = await fetch('/api/swings', {
				method: 'GET'
			});

			if (resp.ok) {
				const data = await resp.json();
				swings = data;
				console.log('Swings', data);
				featuredSwings = swings.slice(0, 6).map((swing) => ({
					id: swing.ID.toString(),
					name: swing.Name,
					swing: swing.Address,
					location: swing.City,
					rating: 4.5,
					reviewCount: 20,
					image: 'https://placehold.co/600x400',
					tags: swing.Tags.map((tag) => tag.Name)
				}));
			}
		} catch (error) {
			console.error('Failed to fetch:', error);
		}
	});
</script>

<section class="px-4 py-16">
	<div class="mx-auto max-w-7xl">
		<div class="grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3">
			{#each featuredSwings as swing}
				<SwingCard {...swing} />
			{/each}
		</div>
	</div>
</section>
