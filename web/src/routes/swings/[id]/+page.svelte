<script lang="ts">
	import type { Swing } from '$lib/types/swing';
	import type { Review } from '$lib/types/review';
	import { getUserIdFromToken, getUserRoleFromToken } from '$lib/utils/jwt';
	import { Trash2 } from 'lucide-svelte';

	export let data: { swing: Swing; photos: string[]; reviews: Review[] };

	let swing: Swing;
	let reviews: Review[];
	let images: string[] = [];
	let currentImage = 0;
	swing = data.swing;
	images = data.photos;
	reviews = data.reviews || [];
	let showForm = false;
	const token = localStorage.getItem('token');

	let newReview = {
		Title: '',
		Rating: 1,
		Body: '',
		UserID: '',
		SwingID: swing.ID,
		VisitedOn: new Date().toISOString().split('T')[0]
	};

	async function handleSubmit() {
		showForm = false;

		try {
			let token = localStorage.getItem('token');
			if (!token) {
				alert('login first');
				return;
			}

			newReview.UserID = JSON.parse(atob(token.split('.')[1])).userID;
			newReview.VisitedOn = newReview.VisitedOn + 'T00:00:00Z';
			const resp = await fetch('/api/auth/reviews/add', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
					Authorization: `Bearer ${token}`
				},
				body: JSON.stringify(newReview)
			});

			if (resp.ok) {
				const response = await fetch(`/api/reviews/${swing.ID}`);
				const reviewData = await response.json();
				reviews = (reviewData.reviews ?? []).map((r: any) => ({
					ID: r.ID,
					SwingID: r.SwingID,
					UserID: r.UserID,
					Rating: r.Rating,
					Title: r.Title,
					Body: r.Body,
					VisitedOn: r.VisitedOn,
					UpdatedAt: r.UpdatedAt ?? ''
				}));
			} else {
				const errorData = await resp.json();
				alert(errorData.error || 'Failed to add review');
			}
		} catch (error) {
			console.error('Failed to submit review:', error);
		}

		newReview = {
			Title: '',
			Rating: 1,
			Body: '',
			UserID: '',
			SwingID: swing.ID,
			VisitedOn: new Date().toISOString().split('T')[0]
		};
	}

	function handleDeleteReview(ID: number): any {
		if (!confirm('Are you sure you want to delete this review?')) {
			return;
		}

		const token = localStorage.getItem('token');
		if (!token) {
			alert('You must be logged in to delete a review.');
			return;
		}

		fetch(`/api/auth/reviews/${ID}`, {
			method: 'DELETE',
			headers: {
				'Content-Type': 'application/json',
				Authorization: `Bearer ${token}`
			}
		})
			.then(async (resp) => {
				if (resp.ok) {
					const response = await fetch(`/api/reviews/${swing.ID}`);
					const reviewData = await response.json();
					reviews = (reviewData.reviews ?? []).map((r: any) => ({
						ID: r.ID,
						SwingID: r.SwingID,
						UserID: r.UserID,
						Rating: r.Rating,
						Title: r.Title,
						Body: r.Body,
						VisitedOn: r.VisitedOn,
						UpdatedAt: r.UpdatedAt ?? ''
					}));
					alert('Review deleted successfully.');
				} else {
					const data = await resp.json();
					alert(data.error || 'Failed to delete review.');
				}
			})
			.catch((error) => {
				console.error('Error deleting review:', error);
			});
	}

	function handleDeleteSwing(ID: number): any {
		if (!confirm('Are you sure you want to delete this swing?')) {
			return;
		}

		const token = localStorage.getItem('token');
		if (!token) {
			alert('You must be logged in to delete a swing.');
			return;
		}

		fetch(`/api/auth/swings/${ID}`, {
			method: 'DELETE',
			headers: {
				'Content-Type': 'application/json',
				Authorization: `Bearer ${token}`
			}
		})
			.then(async (resp) => {
				if (resp.ok) {
					alert('Swing deleted successfully.');
					window.location.href = '/swings';
				} else {
					const data = await resp.json();
					alert(data.error || 'Failed to delete swing.');
				}
			})
			.catch((error) => {
				console.error('Error deleting swing:', error);
			});
	}
</script>

<div class="mx-auto max-w-7xl px-8 py-8">
	<div class="mb-6 flex items-center justify-between">
		<h1 class="mb-8 text-4xl text-gray-800">{swing.Name}</h1>
		<div class="flex items-center gap-4">
			{#if token && getUserIdFromToken(token) === swing.UserID}
				<a
					href={`/swings/${swing.ID}/edit`}
					class="cursor-pointer rounded-lg bg-gray-500 px-4 py-2 text-white hover:bg-gray-600"
					>Edit Swing</a
				>
			{/if}
			{#if token && (getUserIdFromToken(token) === swing.UserID || getUserRoleFromToken(token) === 'admin')}
				<button
					on:click={() => handleDeleteSwing(swing.ID)}
					class="flex cursor-pointer items-center rounded-lg bg-red-500 px-4 py-2 text-white hover:bg-red-600"
					><Trash2 class="mr-2" size="16" /> Delete Swing</button
				>
			{/if}
		</div>
	</div>

	<!-- image -->
	<div class="relative mb-8">
		{#if images.length > 0}
			{#each images as image, i}
				<div class="transition-opacity duration-300" class:hidden={currentImage !== i}>
					<img
						src={image}
						alt={swing.Name}
						class="max-h-96 w-full rounded-lg object-cover shadow-md"
					/>
				</div>
			{/each}
		{:else}
			<div class="flex h-96 items-center justify-center rounded-lg bg-gray-200">
				<img
					src="https://placehold.co/800x400/e2e8f0/64748b?text=No+Images+Available"
					alt="Placeholder"
					class="max-h-96 w-full rounded-lg object-cover shadow-md"
				/>
			</div>
		{/if}

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

	<!-- Reviews -->
	<div class="mt-8 rounded-lg bg-gray-100 p-6 shadow-md">
		<div class="mb-4 flex items-center justify-between">
			<h3 class="text-xl text-gray-700">Reviews</h3>
			<button
				class="cursor-pointer rounded-lg bg-gray-500 px-4 py-2 text-white hover:bg-gray-600"
				on:click={() => (showForm = true)}
			>
				Add Review
			</button>
		</div>
		{#if reviews.length === 0}
			<p class="text-gray-600">No reviews available.</p>
		{/if}

		{#if showForm}
			<div class="fixed inset-0 z-50 flex items-center justify-center bg-black/50">
				<div class="w-full max-w-md rounded-lg bg-white p-6">
					<h2 class="mb-4 text-xl">Add Review</h2>
					<form on:submit|preventDefault={handleSubmit} class="space-y-4">
						<div>
							<label for="title" class="block text-sm text-gray-600">Title</label>
							<input
								id="title"
								type="text"
								bind:value={newReview.Title}
								class="w-full rounded border p-2"
								required
							/>
						</div>
						<div>
							<label for="rating" class="block text-sm text-gray-600">Rating</label>
							<input
								id="rating"
								type="number"
								bind:value={newReview.Rating}
								min="1"
								max="5"
								class="w-full rounded border p-2"
								required
							/>
						</div>
						<div>
							<label for="review" class="block text-sm text-gray-600">Review</label>
							<textarea
								id="review"
								bind:value={newReview.Body}
								class="w-full rounded border p-2"
								required
							></textarea>
						</div>
						<div>
							<label for="visitedOn" class="block text-sm text-gray-600">Visited On</label>
							<input
								id="visitedOn"
								type="date"
								bind:value={newReview.VisitedOn}
								class="w-full rounded border p-2"
								required
							/>
						</div>
						<div class="flex justify-end gap-2">
							<button
								type="button"
								class="rounded bg-gray-300 px-4 py-2"
								on:click={() => (showForm = false)}>Cancel</button
							>
							<button type="submit" class="rounded bg-blue-500 px-4 py-2 text-white">Submit</button>
						</div>
					</form>
				</div>
			</div>
		{/if}
		<ul class="space-y-4">
			{#each reviews as review}
				<li class="rounded-lg bg-white p-4 shadow-sm">
					<div class="mb-2 flex items-center justify-between">
						<h4 class="text-lg font-semibold text-gray-800">
							{review.Title}
						</h4>
						<div class="text-sm text-gray-500">
							User: {review.UserID}
							{#if review.UserID === getUserIdFromToken(token || '') || getUserRoleFromToken(token || '') === 'admin'}
								<button
									class="cursor-pointer text-red-500 hover:underline"
									on:click={() => handleDeleteReview(review.ID)}><Trash2 class="h-5 w-5" /></button
								>
							{/if}
						</div>
					</div>
					<p class="text-black-500">
						Rating: <span class="te</svg>xt-yellow-500"
							>{'★'.repeat(review.Rating)}{'☆'.repeat(5 - review.Rating)}</span
						>
					</p>
					<p class="mt-2 text-gray-600">{review.Body}</p>
					<p class="mt-2 text-sm text-gray-500">
						Visited On: {new Date(review.VisitedOn).toLocaleDateString()}
					</p>
				</li>
			{/each}
		</ul>
	</div>
</div>
