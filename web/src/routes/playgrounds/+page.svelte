<script lang="ts">
	import { onMount } from 'svelte';

	type Playground = {
		ID: number;
		Name: string;
		Address: string;
		City: string;
		Lat: Float64Array;
		Lng: Float64Array;
	};
	let playgrounds: Playground[] = [];

	onMount(async () => {
		try {
			const resp = await fetch('/api/playgrounds', {
				method: 'GET'
			});

			if (resp.ok) {
				const data = await resp.json();
				playgrounds = data;
				console.log('Playgrounds', data);
			}
		} catch (error) {
			console.error('Failed to fetch:', error);
		}
	});
</script>

<section class="px-4 py-16">
	<div class="mx-auto max-w-7xl">
		<div class="mb-12 text-center">
			<h2 class="mb-4 text-3xl font-bold">Bum Playgrounds</h2>
			<p class="mx-auto max-w-2xl text-lg">Super Grounds!</p>
		</div>

		<div class="relative grid grid-cols-1 gap-6 overflow-x-auto md:grid-cols-2 lg:grid-cols-3">
			<table class="border-collapse border border-black">
				<thead>
					<tr class="border border-black">
						<th scope="col" class="px-6 py-3"> ID </th>
						<th scope="col" class="px-6 py-3"> Name </th>
						<th scope="col" class="px-6 py-3"> Address </th>
						<th scope="col" class="px-6 py-3"> City </th>
						<th scope="col" class="px-6 py-3"> Coordinates </th>
					</tr>
				</thead>
				<tbody>
					{#each playgrounds as playground}
						<tr class="border-t border-black font-medium dark:text-black">
							<td class="px-6 py-4">{playground.ID}</td>
							<td class="px-6 py-4">{playground.Name}</td>
							<td class="px-6 py-4">{playground.Address}</td>
							<td class="px-6 py-4">{playground.City}</td>
							<td class="px-6 py-4">{playground.Lat}, {playground.Lng}</td>
						</tr>
					{/each}
				</tbody>
			</table>
		</div>
	</div>
</section>
