<script lang="ts">
	let name = '';
	let address = '';
	let lat = '';
	let lng = '';
	let city = '';
	let seatCount: number;
	let maxWeightKg: number;
	let isAccessible: boolean;
	let tags: string[] = [];
	let tagInput = '';
	let message = 'fafa';
	let err = '';
	async function handleSubmit(e: Event) {
		e.preventDefault();
		err = '';

		try {
			let userID = localStorage.getItem('userID');

			if (!userID) {
				alert('login first');
				return;
			}

			const resp = await fetch('/api/swing', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					UserID: parseInt(userID),
					Name: name,
					Address: address,
					Lat: parseFloat(lat),
					Lng: parseFloat(lng),
					City: city,
					SeatCount: seatCount ? seatCount : null,
					IsAccessible: isAccessible ? isAccessible : false,
					MaxWeightKg: maxWeightKg ? maxWeightKg : null,
					Tags: tags
				})
			});

			if (resp.ok) {
				const data = await resp.json();
				console.log(data);
				message = 'Swing Created Successfully, Redirecting...';
				setTimeout(() => {}, 1000);
				window.location.href = '/swings';
			} else {
				const data = await resp.json();
				err = data.error || 'failed';
			}
		} catch (error) {
			console.error('failed:', error);
			err = 'Error, plz try again!';
		}
	}
</script>

<!-- TODO: add picture upload -->
<!-- TODO: add redirect to swings or the specific swing-->
<!-- TODO: style form -->
<form on:submit={handleSubmit} class="mx-auto max-w-lg space-y-4 rounded">
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

	<div class="flex flex-col items-center font-bold text-red-400">{message}</div>
	<button type="submit" class="add w-full rounded p-2 text-white hover:bg-blue-600"
		>Add Swing</button
	>
</form>

<style>
	@import '../styles.css';
</style>
