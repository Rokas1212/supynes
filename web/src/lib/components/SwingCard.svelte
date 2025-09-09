<script>
	import { Star, MapPin, Users } from 'lucide-svelte';

	export let id;
	export let name;
	export let playground;
	export let location;
	export let rating;
	export let reviewCount;
	export let image;
	/** @type {string[]} */
	export let tags = [];
</script>

<div
	class="hover:shadow-hover bg-card text-card-foreground group cursor-pointer overflow-hidden rounded-2xl border transition-all duration-300 hover:-translate-y-1"
>
	<div class="aspect-[4/3] overflow-hidden">
		<img
			src={image}
			alt={`${name} at ${playground}. location: ${location}, id: ${id}`}
			class="h-full w-full object-cover transition-transform duration-300 group-hover:scale-105"
		/>
	</div>
	<div class="p-4">
		<div class="mb-2 flex items-start justify-between">
			<h3 class="text-foreground line-clamp-1 text-lg font-semibold">{name}</h3>
			<div class="bg-playground-yellow/10 flex items-center gap-1 rounded-full px-2 py-1">
				<Star class="fill-playground-yellow text-playground-yellow h-4 w-4" />
				<span class="text-foreground text-sm font-medium">{rating}</span>
			</div>
		</div>

		<div class="mb-3 space-y-2">
			<div class="text-muted-foreground flex items-center gap-1">
				<MapPin class="h-4 w-4" />
				<span class="text-sm">{playground}</span>
			</div>
			<div class="text-muted-foreground flex items-center gap-1">
				<Users class="h-4 w-4" />
				<span class="text-sm">{reviewCount} reviews</span>
			</div>
		</div>

		<div class="flex flex-wrap gap-1">
			{#each tags.slice(0, 3) as tag}
				<span class="bg-muted text-muted-foreground rounded-full px-2 py-1 text-xs">{tag}</span>
			{/each}
		</div>
	</div>
</div>
