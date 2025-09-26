<script lang="ts">
	import { onMount } from 'svelte';

	type Swing = {
		ID: number;
		Name: string;
		Address: string;
		City: string;
		Lat: Float64Array;
		Lng: Float64Array;
	};
	let swings: Swing[] = [];
	let favorites: number[] = [];

	onMount(async () => {
		try {
			const resp = await fetch('/api/swings', {
				method: 'GET'
			});

			if (resp.ok) {
				const data = await resp.json();
				swings = data;
				console.log('Swings', data);
			}
		} catch (error) {
			console.error('Failed to fetch:', error);
		}

		getFavorites();
	});

	async function getFavorites() {
		let userID = localStorage.getItem('userID');

		if (!userID) {
			alert('login first');
			return;
		}

		try {
			const resp = await fetch(`/api/favorites?userID=${userID}`, {
				method: 'GET',
				headers: {
					'Content-Type': 'application/json'
				}
			});

			if (resp.ok) {
				const data = await resp.json();

				favorites = data.swings;
			}
		} catch (error) {
			console.error('Failed to fetch favorites:', error);
		}
	}

	async function handleFavorite(id: number) {
		let userID = localStorage.getItem('userID');

		if (!userID) {
			alert('login first');
			return;
		}

		try {
			const resp = await fetch('/api/favorite', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					UserID: parseInt(userID),
					SwingID: id
				})
			});

			if (resp.ok) {
				getFavorites();
			} else {
				alert('ups');
			}
		} catch (error) {
			console.error('Failed to add favorite:', error);
		}
	}
</script>

<section class="px-4 py-16">
	<div class="mx-auto max-w-7xl">
		<div class="mb-12 text-center">
			<h2 class="mb-4 text-3xl font-bold">Bum Swings</h2>
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
						<th scope="col" class="px-6 py-3"> Add to favorite </th>
					</tr>
				</thead>
				<tbody>
					{#each swings as swing}
						<tr class="border-t border-black font-medium dark:text-black">
							<td class="px-6 py-4">{swing.ID}</td>
							<td class="px-6 py-4">{swing.Name}</td>
							<td class="px-6 py-4">{swing.Address}</td>
							<td class="px-6 py-4">{swing.City}</td>
							<td class="px-6 py-4">{swing.Lat}, {swing.Lng}</td>
							<td on:click={() => handleFavorite(swing.ID)} class="px-6 py-4 text-center">
								{#if favorites.includes(swing.ID)}★{:else}☆{/if}
							</td>
						</tr>
					{/each}
				</tbody>
			</table>
		</div>
	</div>
</section>
