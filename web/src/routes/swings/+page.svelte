<script lang="ts">
	import { onMount } from 'svelte';
	import { getUserIdFromToken } from '$lib/utils/jwt';
	import { goto } from '$app/navigation';

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
		} catch (error) {
			console.error('Failed to fetch:', error);
		}

		getFavorites();
	});

	async function getFavorites() {
		if (!token) {
			return;
		}
		const userID = getUserIdFromToken(token);

		if (!userID) {
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
				getFavorites();
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
			<h2 class="mb-4 text-3xl font-bold">Bum Swings</h2>
			<p class="mx-auto max-w-2xl text-lg">Super Grounds!</p>
		</div>

		<div class="flex justify-center">
			<div class="overflow-x-auto">
				<table class="border-collapse border border-black">
					<thead>
						<tr class="border border-black">
							<th scope="col" class="px-6 py-3"> ID </th>
							<th scope="col" class="px-6 py-3"> Name </th>
							<th scope="col" class="px-6 py-3"> Address </th>
							<th scope="col" class="px-6 py-3"> City </th>
							<th scope="col" class="px-6 py-3"> Coordinates </th>
							<th scope="col" class="px-6 py-3"> Description </th>
							<th scope="col" class="px-6 py-3"> Add to favorite </th>
						</tr>
					</thead>
					<tbody>
						{#each swings as swing}
							<tr
								class="border-t border-black font-medium dark:text-black {hoveredRow === swing.ID
									? 'bg-gray-200'
									: ''}"
								on:mouseenter={() => (hoveredRow = swing.ID)}
								on:mouseleave={() => (hoveredRow = null)}
								on:click={(event) => handleRowClick(event, swing.ID)}
								style="cursor: pointer;"
							>
								<td class="px-6 py-4">{swing.ID}</td>
								<td class="px-6 py-4">{swing.Name}</td>
								<td class="px-6 py-4">{swing.Address}</td>
								<td class="px-6 py-4">{swing.City}</td>
								<td class="px-6 py-4">{swing.Lat}, {swing.Lng}</td>
								<td class="px-6 py-4">{swing.Description}</td>
								<td
									class="favorite-cell px-6 py-4 text-center"
									on:click|stopPropagation={() => handleFavorite(swing.ID)}
									style="cursor: pointer;"
								>
									{#if favorites.includes(swing.ID)}★{:else}☆{/if}
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
