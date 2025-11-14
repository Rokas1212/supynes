<script lang="ts">
	import { onMount } from 'svelte';
	import { getUserIdFromToken, isAdmin } from '$lib/utils/jwt';
	import { goto } from '$app/navigation';
	import { getFavorites } from '$lib/utils/favorites';
	import { getUserRoleFromRoleNumber } from '$lib/utils/jwt';

	export let title = 'User Profile - Supynės';

	type UserProfile = {
		userID: number;
		displayName: string;
		email: string;
		createdAt: string;
		role: string;
	};

	type Swing = {
		ID: number;
		Name: string;
		Description: string;
		IsAccessible: boolean;
		City: string;
		Address: string;
	};

	type Material = {
		ID: number;
		Name: string;
	};

	let editing = false;
	let newMaterialName = '';
	let profile: UserProfile | null = null;
	let swings: Swing[] = [];
	let favorites: number[];
	let err = '';
	let materials: Material[] = [];

	onMount(async () => {
		document.title = title;
		favorites = await getFavorites(getUserIdFromToken(localStorage.getItem('token') || ''));
		if (isAdmin(localStorage.getItem('token') || '')) {
			await getMaterials();
		}

		const resp = await fetch('/api/auth/profile', {
			method: 'GET',
			headers: {
				'Content-Type': 'application/json',
				Authorization: `Bearer ${localStorage.getItem('token')}`
			}
		});

		if (resp.ok) {
			const data = await resp.json();
			profile = {
				userID: data.userID,
				displayName: data.displayName,
				email: data.email,
				createdAt: new Date(data.createdAt).toLocaleDateString(),
				role: data.role
			};
			swings = data.swings.map((swing: Swing) => ({
				ID: swing.ID,
				Name: swing.Name,
				Description: swing.Description,
				IsAccessible: swing.IsAccessible,
				City: swing.City,
				Address: swing.Address
			}));
		} else {
			err = 'Failed to load profile data, ' + `${resp.status} - ${resp.statusText}`;
		}
	});

	async function getMaterials() {
		try {
			const response = await fetch('/api/materials');
			materials = await response.json();
		} catch (error) {
			console.error('Failed to load materials:', error);
		}
	}

	async function handleUpdateMaterial(materialID: number, materialName: string) {
		const token = localStorage.getItem('token');
		if (!token) {
			alert('You must be logged in to update a material.');
			return;
		}
		if (materialName.trim() === '') {
			alert('Material name cannot be empty.');
			return;
		}
		try {
			const resp = await fetch(`/api/auth/materials/update/${materialID}`, {
				method: 'PUT',
				headers: {
					'Content-Type': 'application/json',
					Authorization: `Bearer ${token}`
				},
				body: JSON.stringify({ Name: materialName.trim() })
			});
			if (resp.ok) {
				alert('Material updated successfully.');
				editing = false;
				await getMaterials();
			} else {
				const data = await resp.json();
				alert(data.error || 'Failed to update material.');
			}
		} catch (error) {
			alert('Network error. Please try again.');
		}
	}

	async function handleAddMaterial() {
		const token = localStorage.getItem('token');
		if (!token) {
			alert('You must be logged in to add a material.');
			return;
		}
		if (newMaterialName.trim() === '') {
			alert('Material name cannot be empty.');
			return;
		}
		try {
			const resp = await fetch('/api/auth/materials/add', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
					Authorization: `Bearer ${token}`
				},
				body: JSON.stringify({ Name: newMaterialName.trim() })
			});
			if (resp.ok) {
				alert('Material added successfully.');
				newMaterialName = '';
				await getMaterials();
			} else {
				const data = await resp.json();
				alert(data.error || 'Failed to add material.');
			}
		} catch (error) {
			alert('Network error. Please try again.');
		}
	}

	async function handleMaterialDelete(materialID: number) {
		const token = localStorage.getItem('token');
		if (!token) {
			alert('You must be logged in to delete a material.');
			return;
		}
		if (!confirm('Are you sure you want to delete this material?')) {
			return;
		}
		try {
			const resp = await fetch(`/api/auth/materials/remove/${materialID}`, {
				method: 'DELETE',
				headers: {
					'Content-Type': 'application/json',
					Authorization: `Bearer ${token}`
				}
			});
			if (resp.ok) {
				alert('Material deleted successfully.');
				materials = materials.filter((material) => material.ID !== materialID);
			} else {
				const data = await resp.json();
				alert(data.error || 'Failed to delete material.');
			}
		} catch (error) {
			alert('Network error. Please try again.');
		}
	}

	async function handleUserDelete() {
		const token = localStorage.getItem('token');
		if (!token) {
			alert('You must be logged in to delete your user.');
			return;
		}
		if (!confirm('Are you sure you want to delete your account? This action cannot be undone.')) {
			return;
		}
		try {
			const userID = getUserIdFromToken(token);
			if (!userID) {
				alert('Invalid token. Please log in again.');
				return;
			}
			const resp = await fetch('/api/auth/users/delete', {
				method: 'DELETE',
				headers: {
					'Content-Type': 'application/json',
					Authorization: `Bearer ${token}`
				}
			});
			if (resp.ok) {
				localStorage.removeItem('token');
				alert('Your account has been deleted.');
				goto('/');
			} else {
				const data = await resp.json();
				alert(data.error || 'Failed to delete user.');
				localStorage.removeItem('token');
				goto('/users/login');
			}
		} catch (error) {
			alert('Network error. Please try again.');
		}
	}

	async function handleSwingDelete(swingID: number) {
		const token = localStorage.getItem('token');
		if (!token) {
			alert('You must be logged in to delete a swing.');
			return;
		}
		if (!confirm('Are you sure you want to delete this swing? This action cannot be undone.')) {
			return;
		}
		try {
			const resp = await fetch(`/api/auth/swings/${swingID}`, {
				method: 'DELETE',
				headers: {
					'Content-Type': 'application/json',
					Authorization: `Bearer ${token}`
				}
			});
			if (resp.ok) {
				alert('Swing deleted successfully.');
				swings = swings.filter((swing) => swing.ID !== swingID);
			} else {
				const data = await resp.json();
				alert(data.error || 'Failed to delete swing.');
			}
		} catch (error) {
			alert('Network error. Please try again.');
		}
	}

	async function getSwingNameByID(swingID: number) {
		try {
			const resp = await fetch(`/api/swing-name/${swingID}`, {
				method: 'GET'
			});
			if (resp.ok) {
				const data = await resp.json();
				return data.name;
			} else {
				const data = await resp.json();
				console.error(data.error || 'Failed to fetch swing name.');
			}
		} catch (error) {
			console.error('Network error. Please try again.', error);
		}
	}
</script>

<div class="mx-auto max-w-3xl p-6">
	{#if err}
		<p class="mb-4 flex justify-center text-red-500">{err}</p>
	{/if}
	{#if profile}
		<h2 class="mb-6 text-3xl font-bold text-gray-800">{profile.displayName}'s Profile</h2>
		<div class="mb-8 rounded-lg bg-white p-6 shadow-md">
			<p class="mb-3"><span class="font-semibold text-gray-700">Email:</span> {profile.email}</p>
			<p class="mb-3">
				<span class="font-semibold text-gray-700">Member Since:</span>
				{profile.createdAt}
			</p>
			<p class="mb-3">
				<span class="font-semibold text-gray-700">Role:</span>
				{getUserRoleFromRoleNumber(profile.role)}
			</p>
			<button
				type="button"
				on:click={() => handleUserDelete()}
				class="inline-flex items-center rounded-md border border-transparent bg-red-600 px-4 py-2 text-sm font-medium text-white hover:bg-red-700"
			>
				Delete My Account
			</button>
		</div>
	{/if}
	<div>
		<h3 class="mb-4 text-2xl font-bold text-gray-800">Your Swings</h3>
		{#if swings.length > 0}
			<ul class="space-y-4">
				{#each swings as swing}
					<li>
						<a
							href={`/swings/${swing.ID}`}
							class="block rounded-lg bg-white p-4 shadow-sm transition-shadow hover:shadow-md"
						>
							<div class="mb-2 flex items-center justify-between">
								<h4 class="text-lg font-semibold text-gray-800">{swing.Name}</h4>
								<button
									type="button"
									class="inline-flex items-center rounded-md border border-transparent bg-red-600 px-4 py-2 text-sm font-medium text-white hover:bg-red-700"
									on:click|preventDefault|stopPropagation={() => handleSwingDelete(swing.ID)}
								>
									Delete
								</button>
							</div>
							<p class="text-gray-600">{swing.Description}</p>
							<p class="mt-2 text-sm text-gray-500">
								{#if swing.IsAccessible}
									<span
										class="inline-block rounded-full bg-green-100 px-2 py-1 text-xs font-semibold text-green-800"
										>Accessible</span
									>
								{:else}
									<span
										class="inline-block rounded-full bg-red-100 px-2 py-1 text-xs font-semibold text-red-800"
										>Not Accessible</span
									>
								{/if}
							</p>
						</a>
					</li>
				{/each}
			</ul>
		{:else}
			<p class="italic text-gray-600">No swings found.</p>
		{/if}
	</div>
	<div class="mb-8 mt-8">
		<h3 class="mb-4 text-2xl font-bold text-gray-800">Your Favorites</h3>
		{#if favorites != null && favorites.length > 0}
			<ul class="space-y-4">
				{#each favorites as swingID}
					<li>
						<a
							href={`/swings/${swingID}`}
							class="block rounded-lg bg-white p-4 shadow-sm transition-shadow hover:shadow-md"
						>
							<div class="mb-2 flex items-center justify-between">
								{#await getSwingNameByID(swingID)}
									<h4 class="text-lg font-semibold text-gray-800">Loading...</h4>
								{:then swingName}
									<h4 class="text-lg font-semibold text-gray-800">{swingName}</h4>
								{:catch}
									<h4 class="text-lg font-semibold text-red-800">Error loading swing name</h4>
								{/await}
							</div>
						</a>
					</li>
				{/each}
			</ul>
		{:else}
			<p class="italic text-gray-600">No favorites found.</p>
		{/if}
	</div>
	{#if isAdmin(localStorage.getItem('token') || '')}
		<div class="mt-8">
			<h3 class="mb-4 text-2xl font-bold text-gray-800">Existing Materials</h3>
			{#if materials.length > 0}
				<ul class="pb-10">
					{#each materials as material}
						<li class="pb-2 text-gray-700">
							{material.Name}
							<button
								class="inline-flex items-center rounded-md border border-transparent bg-red-600 px-1 py-1 text-sm font-medium text-white hover:bg-red-700"
								on:click={() => handleMaterialDelete(material.ID)}>Delete</button
							>
							{#if editing}
								<input
									type="text"
									bind:value={material.Name}
									class="ml-2 rounded border px-2 py-1 text-gray-700 focus:border-blue-500 focus:outline-none"
								/>
								<button
									class="inline-flex items-center rounded-md border border-transparent bg-green-600 px-1 py-1 text-sm font-medium text-white hover:bg-green-700"
									on:click={() => handleUpdateMaterial(material.ID, material.Name)}>Save</button
								>
							{:else}
								<button
									class="inline-flex items-center rounded-md border border-transparent bg-blue-600 px-1 py-1 text-sm font-medium text-white hover:bg-blue-700"
									on:click={() => {
										editing = true;
									}}>Edit</button
								>
							{/if}
						</li>
					{/each}
				</ul>
			{:else}
				<p class="italic text-gray-600">No materials found.</p>
			{/if}
		</div>

		<div class="mt-8 pb-20">
			<h3 class="mb-4 text-2xl font-bold text-gray-800">Add New Material</h3>
			<input
				type="text"
				bind:value={newMaterialName}
				placeholder="Material name"
				class="rounded border px-3 py-2 text-gray-700 focus:border-blue-500 focus:outline-none"
			/>
			<button
				class="inline-flex cursor-pointer items-center rounded-md border border-transparent bg-gray-600 px-4 py-2 text-sm font-medium text-white hover:bg-indigo-700"
				on:click={handleAddMaterial}
			>
				Add
			</button>
		</div>
	{/if}
</div>
