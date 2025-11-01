<script lang="ts">
	import { onMount } from 'svelte';
	import { getUserIdFromToken } from '$lib/utils/jwt';
	import { goto } from '$app/navigation';
	import { getFavorites } from '$lib/utils/favorites';

	type Swing = {
		ID: number;
		Name: string;
		Description: string;
		Address: string;
		City: string;
		Lat: Float64Array;
		Lng: Float64Array;
	};
	let swings: Swing[] = [];
	let favorites: number[] = [];
	let token: string | null = null;

	onMount(async () => {
		try {
			token = localStorage.getItem('token');
			const resp = await fetch('/api/swings', {
				method: 'GET'
			});

			if (resp.ok) {
				const data = await resp.json();
				swings = data;
				console.log('Swings', data);
			}

			favorites = await getFavorites(getUserIdFromToken(token || ''));
		} catch (error) {
			console.error('Failed to fetch:', error);
		}
	});

	async function getAverageRating(swingID: number): Promise<number | null> {
		try {
			const resp = await fetch(`/api/average-ratings/${swingID}`, {
				method: 'GET'
			});

			if (resp.ok) {
				const data = await resp.json();
				return data.averageRating;
			} else {
				console.error('Failed to fetch average rating');
				return null;
			}
		} catch (error) {
			console.error('Error fetching average rating:', error);
			return null;
		}
	}

	async function handleFavorite(id: number) {
		if (!token) {
			goto('/users/login');
			return;
		}
		const userID = getUserIdFromToken(token);
		if (!userID) {
			goto('/users/login');
			return;
		}
		console.log(userID);

		try {
			const resp = await fetch('/api/auth/favorite', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
					Authorization: `Bearer ${token}`
				},
				body: JSON.stringify({
					SwingID: id
				})
			});

			if (resp.ok) {
				favorites = await getFavorites(userID);
			} else {
				alert('ups');
			}
		} catch (error) {
			console.error('Failed to add favorite:', error);
		}
	}

	let hoveredRow: number | null = null;

	function handleRowClick(event: MouseEvent, id: number) {
		const target = event.target as HTMLElement;
		if (target.closest('.favorite-cell')) return;
		goto(`/swings/${id}`);
	}
</script>

<section class="px-4 py-16">
	<div class="mx-auto max-w-7xl">
		<div class="mb-12 text-center">
			<h2 class="mb-4 text-4xl font-extrabold text-gray-900">Swings</h2>
		</div>

		<div class="flex justify-center">
			<div class="overflow-x-auto rounded-lg shadow-lg">
				<table class="w-full min-w-full divide-y divide-gray-200">
					<thead class="bg-gray-50">
						<tr>
							<th
								scope="col"
								class="px-6 py-4 text-left text-xs font-medium uppercase tracking-wider text-gray-500"
							>
								Name
							</th>
							<th
								scope="col"
								class="px-6 py-4 text-left text-xs font-medium uppercase tracking-wider text-gray-500"
							>
								Address
							</th>
							<th
								scope="col"
								class="px-6 py-4 text-left text-xs font-medium uppercase tracking-wider text-gray-500"
							>
								City
							</th>
							<th
								scope="col"
								class="px-6 py-4 text-left text-xs font-medium uppercase tracking-wider text-gray-500"
							>
								Coordinates
							</th>
							<th
								scope="col"
								class="px-6 py-4 text-left text-xs font-medium uppercase tracking-wider text-gray-500"
							>
								Description
							</th>
							<th
								scope="col"
								class="px-6 py-4 text-left text-xs font-medium uppercase tracking-wider text-gray-500"
								>Average Rating
							</th>
							<th
								scope="col"
								class="px-6 py-4 text-center text-xs font-medium uppercase tracking-wider text-gray-500"
							>
								Favorite
							</th>
						</tr>
					</thead>
					<tbody class="divide-y divide-gray-200 bg-white">
						{#each swings as swing}
							<tr
								class="transition-colors hover:bg-gray-50 {hoveredRow === swing.ID
									? 'bg-gray-50'
									: ''}"
								on:mouseenter={() => (hoveredRow = swing.ID)}
								on:mouseleave={() => (hoveredRow = null)}
								on:click={(event) => handleRowClick(event, swing.ID)}
								style="cursor: pointer;"
							>
								<td class="whitespace-nowrap px-6 py-4 text-sm font-medium text-gray-900"
									>{swing.Name}</td
								>
								<td class="whitespace-nowrap px-6 py-4 text-sm text-gray-500">{swing.Address}</td>
								<td class="whitespace-nowrap px-6 py-4 text-sm text-gray-500">{swing.City}</td>
								<td class="whitespace-nowrap px-6 py-4 text-sm text-gray-500"
									>Lat: {swing.Lat}°, Lng: {swing.Lng}°</td
								>
								<td class="px-6 py-4 text-sm text-gray-500">{swing.Description}</td>
								<td class="whitespace-nowrap px-6 py-4 text-sm text-gray-500">
									{#await getAverageRating(swing.ID) then averageRating}
										{#if averageRating !== null}
											{averageRating.toFixed(2)}
										{:else}
											N/A
										{/if}
									{/await}
								</td>
								<td
									class="favorite-cell whitespace-nowrap px-6 py-4 text-center text-xl"
									on:click|stopPropagation={() => handleFavorite(swing.ID)}
									style="cursor: pointer;"
								>
									{#if favorites.includes(swing.ID)}
										<span class="text-yellow-400">★</span>
									{:else}
										<span class="text-gray-400 hover:text-yellow-400">☆</span>
									{/if}
								</td>
							</tr>
						{/each}
					</tbody>
				</table>
			</div>
		</div>
	</div>
</section>

<style>
	@import './styles.css';
</style>
