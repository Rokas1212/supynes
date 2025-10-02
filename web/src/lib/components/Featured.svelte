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
				featuredSwings = swings.map((swing) => ({
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

	// const featuredSwings = [
	// 	{
	// 		id: '1',
	// 		name: 'Quadro Swings',
	// 		swing: 'Sunny Meadows Park',
	// 		location: 'Downtown District',
	// 		rating: 4.8,
	// 		reviewCount: 24,
	// 		image:
	// 			'https://static.wixstatic.com/media/373df8_a357b0d6e9d14a0386509ff8581d451b~mv2.jpg/v1/fill/w_560,h_420,al_c,q_80,usm_0.66_1.00_0.01,enc_avif,quality_auto/373df8_a357b0d6e9d14a0386509ff8581d451b~mv2.jpg',
	// 		tags: ['Family Friendly', 'Accessible', 'Shaded']
	// 	},
	// 	{
	// 		id: '2',
	// 		name: 'Adventure Tire Swing',
	// 		swing: 'Riverside Recreation Center',
	// 		location: 'Westside',
	// 		rating: 4.6,
	// 		reviewCount: 18,
	// 		image:
	// 			'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQTYnI65ak-JgE6G2hlmS7Wib6my_o_luuJ8Q&s',
	// 		tags: ['Adventure', 'Age 5+', 'Popular']
	// 	},
	// 	{
	// 		id: '3',
	// 		name: 'Baby Bucket Swings',
	// 		swing: 'Little Tots Swing',
	// 		location: 'Northside',
	// 		rating: 4.9,
	// 		reviewCount: 32,
	// 		image:
	// 			'https://d3cy9zhslanhfa.cloudfront.net/media/5566D47B-D120-443F-860DA1F6ADB1922D/A2562916-C3C4-4D36-8F14BB076DE519A3/webimage-671C73CC-5036-4C5F-B7490F93A71950D3.jpg',
	// 		tags: ['Toddler Safe', 'Fenced', 'Clean']
	// 	},
	// 	{
	// 		id: '4',
	// 		name: 'Double Swing Set',
	// 		swing: 'Central Community Park',
	// 		location: 'City Center',
	// 		rating: 4.5,
	// 		reviewCount: 15,
	// 		image:
	// 			'https://www.sodexsport.com/sodex_sport_library/Product-photos/EN/equipment/swing/bragmaia-vietnam-swing-double-swings-set-zea.jpg',
	// 		tags: ['Social', 'Well-maintained', 'Popular']
	// 	},
	// 	{
	// 		id: '5',
	// 		name: 'Handicap Accessible Swing',
	// 		swing: 'Inclusive Play Space',
	// 		location: 'East District',
	// 		rating: 5.0,
	// 		reviewCount: 28,
	// 		image: 'https://www.rehabmart.com/imagesfromrd/sportsplay_Wheelchair_Swing_Platform.jpg',
	// 		tags: ['Accessible', 'Inclusive', 'Modern']
	// 	},
	// 	{
	// 		id: '6',
	// 		name: 'Wooden Swing Set',
	// 		swing: 'Heritage Park',
	// 		location: 'Historic District',
	// 		rating: 4.4,
	// 		reviewCount: 21,
	// 		image: 'https://tara-design.com/cdn/shop/products/swing_2_large.jpg?v=1525201197',
	// 		tags: ['Traditional', 'Natural', 'Scenic']
	// 	}
	// ];
</script>

<section class="px-4 py-16">
	<div class="mx-auto max-w-7xl">
		<div class="mb-12 text-center">
			<h2 class="mb-4 text-3xl font-bold">Featured Swings</h2>
			<p class="mx-auto max-w-2xl text-lg">Super Swings !</p>
		</div>

		<div class="grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3">
			{#each featuredSwings as swing}
				<SwingCard {...swing} />
			{/each}
		</div>
	</div>
</section>
