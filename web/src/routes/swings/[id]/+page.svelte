<script lang="ts">
	import type { Swing } from '$lib/types/swing';

	export let data: { swing: Swing; photos: string[] };

	let swing: Swing;
	let images: string[] = [];
	let currentImage = 0;
	swing = data.swing;
	images = data.photos;
</script>

<div class="mx-auto max-w-7xl px-8 py-8">
	<h1 class="mb-8 text-4xl text-gray-800">{swing.Name}</h1>

	<!-- image -->
	<div class="relative mb-8">
		{#each images as image, i}
			<div class="transition-opacity duration-300" class:hidden={currentImage !== i}>
				<img
					src={image.replace('/app/media/', '/media/')}
					alt={swing.Name}
					class="max-h-96 w-full rounded-lg object-cover shadow-md"
				/>
			</div>
		{/each}

		<!-- Navigation buttons -->
		<button
			class="absolute left-4 top-1/2 -translate-y-1/2 rounded-full bg-black/50 p-2 text-white"
			on:click={() => (currentImage = (currentImage - 1 + images.length) % images.length)}
		>
			←
		</button>
		<button
			class="absolute right-4 top-1/2 -translate-y-1/2 rounded-full bg-black/50 p-2 text-white"
			on:click={() => (currentImage = (currentImage + 1) % images.length)}
		>
			→
		</button>
	</div>

	<div class="grid grid-cols-1 gap-8 md:grid-cols-2 lg:grid-cols-3">
		<div class="rounded-lg bg-gray-100 p-6 shadow-md">
			<h3 class="mb-4 text-xl text-gray-700">Details</h3>
			<p class="text-gray-600">{swing.Description || 'No description available'}</p>
		</div>

		<div class="rounded-lg bg-gray-100 p-6 shadow-md">
			<h3 class="mb-4 text-xl text-gray-700">Specifications</h3>
			<ul class="space-y-2">
				<li class="text-gray-600">Seat Count: {swing.SeatCount}</li>
				<li class="text-gray-600">Max Weight: {swing.MaxWeightKg}kg</li>
				<li class="text-gray-600">Location: {swing.City}</li>
				<li class="text-gray-600">Accessible: {swing.IsAccessible ? 'Yes' : 'No'}</li>
			</ul>
		</div>

		<div class="rounded-lg bg-gray-100 p-6 shadow-md">
			<h3 class="mb-4 text-xl text-gray-700">Coordinates</h3>
			<ul class="space-y-2">
				<li class="text-gray-600">Lng: {swing.Lng}</li>
				<li class="text-gray-600">Lat: {swing.Lat}</li>
			</ul>
		</div>
	</div>
</div>
