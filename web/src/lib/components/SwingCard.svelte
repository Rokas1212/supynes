<script>
	import { Star, MapPin, Users } from 'lucide-svelte';

	export let id;
	export let name;
	export let swing;
	export let location;
	export let rating;
	export let reviewCount;
	export let image;
	/** @type {string[]} */
	export let tags = [];
</script>

<div
	class="group cursor-pointer overflow-hidden rounded-2xl border duration-300 hover:-translate-y-5"
>
	<div class="aspect-[4/3] overflow-hidden">
		<img
			src={image}
			alt={`${name} at ${swing}. location: ${location}, id: ${id}`}
			class="h-full w-full object-cover duration-300 group-hover:scale-105"
		/>
	</div>
	<div class="p-3">
		<div class="mb-2 flex justify-between">
			<h3 class="line-clamp-1 text-lg font-semibold">{name}</h3>
			<div class="flex items-center gap-1 px-2 py-1">
				<Star class="h-4 w-4" />
				<span class="text-sm font-medium">{rating}</span>
			</div>
		</div>

		<div class="mb-3 space-y-2">
			<div class="flex items-center gap-1">
				<MapPin class="h-4 w-4" />
				<span class="text-sm">{swing}</span>
			</div>
			<div class="flex items-center gap-1">
				<Users class="h-4 w-4" />
				<span class="text-sm">{reviewCount} reviews</span>
			</div>
		</div>

		<div class="flex flex-wrap gap-1">
			{#each tags.slice(0, 3) as tag}
				<span class="px-1 py-1 text-xs">{tag}</span>
			{/each}
		</div>
	</div>
</div>
