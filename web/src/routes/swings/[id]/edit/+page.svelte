<script lang="ts">
	import { goto } from '$app/navigation';
	import { getUserIdFromToken } from '$lib/utils/jwt';
	import { onMount } from 'svelte';
	import { page } from '$app/stores';

	let name = '';
	let address = '';
	let lat = '';
	let lng = '';
	let city = '';
	let seatCount: number;
	let maxWeightKg: number;
	let isAccessible: boolean = false;
	let tags: string[] = [];
	let materials: number[] = [];
	let tagInput = '';
	let message = '';
	let err = '';
	let materialsOpen = false;
	let loading = true;

	let materialsList: { ID: number; Name: string }[] = [];
	let userID: number | null = null;
	let token: string | null = null;

	const swingId = $page.params.id;

	onMount(async () => {
		token = localStorage.getItem('token');
		userID = getUserIdFromToken(token || '');

		try {
			// Load materials
			const materialsResponse = await fetch('/api/materials');
			materialsList = await materialsResponse.json();

			// Load existing swing data
			const swingResponse = await fetch(`/api/swings/${swingId}`);
			if (swingResponse.ok) {
				const swing = await swingResponse.json();

				// Populate form with existing data
				name = swing.Name || '';
				address = swing.Address || '';
				lat = swing.Lat?.toString() || '';
				lng = swing.Lng?.toString() || '';
				city = swing.City || '';
				seatCount = swing.SeatCount || 0;
				maxWeightKg = swing.MaxWeightKg || 0;
				isAccessible = swing.IsAccessible || false;
				tags = swing.Tags?.map((tag: any) => tag.Name) || [];
				materials = swing.Materials?.map((material: any) => material.ID) || [];
			} else {
				err = 'Failed to load swing data';
			}
		} catch (error) {
			console.error('Failed to load data:', error);
			err = 'Failed to load data';
		} finally {
			loading = false;
		}
	});

	async function handleSubmit(e: Event) {
		e.preventDefault();
		err = '';

		try {
			if (!token) {
				alert('login first');
				return;
			}
			if (!userID) {
				alert('invalid token, login again');
				return;
			}

			const formData = new FormData();
			formData.append('UserID', userID.toString());
			formData.append('Name', name);
			formData.append('Address', address);
			formData.append('Lat', lat);
			formData.append('Lng', lng);
			formData.append('City', city);
			if (seatCount) formData.append('SeatCount', seatCount.toString());
			if (maxWeightKg) formData.append('MaxWeightKg', maxWeightKg.toString());
			formData.append('IsAccessible', isAccessible ? 'true' : 'false');
			tags.forEach((tag) => formData.append('Tags[]', tag));
			materials.forEach((id) => formData.append('MaterialIDs[]', id.toString()));

			const resp = await fetch(`/api/auth/swings/${swingId}`, {
				method: 'PUT',
				headers: {
					Authorization: `Bearer ${token}`
				},
				body: formData
			});

			if (resp.ok) {
				const data = await resp.json();
				console.log('Swing updated successfully', data);
				message = 'Swing Updated Successfully, Redirecting...';
				setTimeout(() => goto(`/swings/${swingId}`), 2000);
			} else {
				const data = await resp.json();
				err = data.error || 'failed to update';
			}
		} catch (error) {
			console.error('failed:', error);
			err = 'Error, plz try again!';
		}
	}
</script>

{#if loading}
	<div class="flex justify-center p-8">
		<p>Loading swing data...</p>
	</div>
{:else}
	<form on:submit={handleSubmit} class="mx-auto max-w-lg space-y-4 rounded">
		<h1 class="mb-6 text-2xl font-bold">Edit Swing</h1>

		<div class="flex flex-col">
			<label for="name" class="mb-1 text-sm font-medium">Name</label>
			<input
				id="name"
				name="name"
				type="text"
				required
				bind:value={name}
				class="rounded border p-2"
			/>
		</div>

		<div class="grid grid-cols-2 gap-4">
			<div class="flex flex-col">
				<label for="lat" class="mb-1 text-sm font-medium">Latitude</label>
				<input
					id="lat"
					name="lat"
					type="number"
					step="any"
					required
					bind:value={lat}
					class="rounded border p-2"
				/>
			</div>
			<div class="flex flex-col">
				<label for="lng" class="mb-1 text-sm font-medium">Longitude</label>
				<input
					id="lng"
					name="lng"
					type="number"
					step="any"
					required
					bind:value={lng}
					class="rounded border p-2"
				/>
			</div>
		</div>

		<div class="flex flex-col">
			<label for="address" class="mb-1 text-sm font-medium">Full Street</label>
			<input
				id="address"
				name="address"
				type="text"
				required
				bind:value={address}
				class="rounded border p-2"
			/>
		</div>

		<div class="flex flex-col">
			<label for="city" class="mb-1 text-sm font-medium">City</label>
			<input
				id="city"
				name="city"
				type="text"
				required
				bind:value={city}
				class="rounded border p-2"
			/>
		</div>

		<div class="grid grid-cols-2 gap-4">
			<div class="flex flex-col">
				<label for="seatCount" class="mb-1 text-sm font-medium">Seat Count</label>
				<input
					id="seatCount"
					name="seatCount"
					type="number"
					bind:value={seatCount}
					class="rounded border p-2"
				/>
			</div>
			<div class="flex flex-col">
				<label for="maxWeight" class="mb-1 text-sm font-medium">Max Weight (kg)</label>
				<input
					id="maxWeight"
					name="maxWeight"
					type="number"
					bind:value={maxWeightKg}
					class="rounded border p-2"
				/>
			</div>
		</div>

		<div class="flex items-center">
			<input
				id="isAccessible"
				name="isAccessible"
				type="checkbox"
				bind:checked={isAccessible}
				class="mr-2"
			/>
			<label for="isAccessible" class="text-sm font-medium">Accessible for disabled</label>
		</div>

		<div class="flex flex-col">
			<label for="tags" class="mb-1 text-sm font-medium">Tags</label>
			<input
				id="tags"
				name="tags"
				type="text"
				bind:value={tagInput}
				on:keydown={(e) => {
					if (e.key === 'Enter') {
						e.preventDefault();
						if (tagInput.trim()) {
							if (tags.includes(tagInput.trim())) {
								alert('Tag already added');
								return;
							}
							tags = [...tags, tagInput.trim()];
							tagInput = '';
						}
					}
				}}
				placeholder="Click enter to submit tag"
				class="rounded border p-2"
			/>

			{#if tags.length > 0}
				<div class="mt-2 text-sm text-black">
					{#each tags as tag}
						<span class="bg-gray tags mr-2 inline-block rounded px-2 py-1">
							{tag}
							<button
								class="text-gray ml-1"
								on:click={() => {
									tags = tags.filter((t) => t !== tag);
								}}
							>
								x
							</button>
						</span>
					{/each}
				</div>
			{/if}
		</div>

		<div class="flex flex-col">
			<label for="materials" class="mb-1 text-sm font-medium">Materials</label>
			<div class="relative">
				<button
					type="button"
					class="w-full rounded border p-2 text-left"
					on:click={() => (materialsOpen = !materialsOpen)}
				>
					{materials.length
						? materialsList
								.filter((m) => materials.includes(m.ID))
								.map((m) => m.Name)
								.join(', ')
						: 'Select materials'}
				</button>
				{#if materialsOpen}
					<div
						class="absolute z-10 mt-1 max-h-48 w-full overflow-y-auto rounded border bg-white shadow-lg"
					>
						{#each materialsList as material}
							<label class="block p-2 hover:bg-gray-100">
								<input
									type="checkbox"
									value={material.ID}
									checked={materials.includes(material.ID)}
									on:change={() => {
										if (materials.includes(material.ID)) {
											materials = materials.filter((id) => id !== material.ID);
										} else {
											materials = [...materials, material.ID];
										}
									}}
								/>
								<span class="ml-2">{material.Name}</span>
							</label>
						{/each}
					</div>
				{/if}
			</div>
		</div>

		{#if err}
			<div class="flex flex-col items-center font-bold text-red-400">{err}</div>
		{/if}
		{#if message}
			<div class="flex flex-col items-center font-bold text-green-400">{message}</div>
		{/if}

		<button type="submit" class="w-full rounded bg-blue-500 p-2 text-white hover:bg-blue-600">
			Update Swing
		</button>
	</form>
{/if}
