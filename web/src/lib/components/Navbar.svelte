<script lang="ts">
	import { onMount } from 'svelte';

	// Navbar props
	export let title: string = 'Supynes';
	export let navLinks: Array<{ href: string; label: string }> = [
		{ href: '/swings', label: 'Swings' },
		{ href: '/about', label: 'About' }
	];
	export let showLogin: boolean = true;

	// Active link tracking
	let currentPath: string;

	if (typeof window !== 'undefined') {
		currentPath = window.location.pathname;
	}

	$: isActive = (path: string) => currentPath === path;

	let isOpen: boolean = false;

	onMount(async () => {
		if (localStorage.getItem('token')) {
			showLogin = false;
		}
	});

	async function handleUserDelete() {
		const userID = localStorage.getItem('userID');
		const token = localStorage.getItem('token');
		if (!userID || !token) {
			alert('You must be logged in to delete your user.');
			return;
		}
		if (!confirm('Are you sure you want to delete your account? This action cannot be undone.')) {
			return;
		}
		try {
			const resp = await fetch('/api/users/delete', {
				method: 'DELETE',
				headers: {
					'Content-Type': 'application/json',
					//TODO: ADD ACTUAL TOKEN eh
					Authorization: `Bearer ${token}`
				},
				body: JSON.stringify({ userID: parseInt(userID) })
			});
			if (resp.ok) {
				localStorage.removeItem('userID');
				localStorage.removeItem('token');
				alert('Your account has been deleted.');
				window.location.href = '/';
			} else {
				const data = await resp.json();
				alert(data.error || 'Failed to delete user.');
			}
		} catch (error) {
			alert('Network error. Please try again.');
		}
	}

	async function logout() {
		localStorage.clear();
		window.location.href = '/';
	}
</script>

<nav class="border-b bg-white shadow-sm">
	<div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
		<div class="flex h-16 justify-between">
			<div class="flex items-center">
				<!-- Logo -->
				<a href="/" class="flex flex-shrink-0 items-center">
					<span class="text-xl font-bold text-gray-800">{title}</span>
				</a>

				<!-- Nav Links -->
				<div class="hidden sm:ml-6 sm:flex sm:space-x-8">
					{#each navLinks as link}
						<a
							href={link.href}
							class="px-3 py-2 text-sm font-medium {isActive(link.href)
								? 'border-b-2 border-indigo-500 text-gray-900'
								: 'text-gray-500 hover:border-gray-300 hover:text-gray-700'}"
							aria-current={isActive(link.href) ? 'page' : undefined}
						>
							{link.label}
						</a>
					{/each}
				</div>
			</div>

			<!-- Right side buttons -->
			<div class="flex items-center">
				{#if showLogin}
					<a
						href="/users/login"
						class="inline-flex items-center rounded-md border border-transparent bg-indigo-600 px-4 py-2 text-sm font-medium text-white hover:bg-indigo-700"
					>
						Log in
					</a>
				{:else}
					<button
						on:click={() => logout()}
						class="inline-flex items-center rounded-md border border-transparent bg-indigo-600 px-4 py-2 text-sm font-medium text-white hover:bg-indigo-700"
					>
						Log out
					</button>
					<button
						type="button"
						on:click={() => handleUserDelete()}
						class="inline-flex items-center rounded-md border border-transparent bg-indigo-600 px-4 py-2 text-sm font-medium text-white hover:bg-indigo-700"
					>
						Delete My User
					</button>
				{/if}

				<!-- Mobile menu button -->
				<div class="ml-4 sm:hidden">
					<button
						type="button"
						class="inline-flex items-center justify-center rounded-md p-2 text-gray-400 hover:bg-gray-100 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-indigo-500"
						aria-expanded={isOpen}
						on:click={() => (isOpen = !isOpen)}
					>
						<span class="sr-only">Open main menu</span>
						{#if isOpen}
							<!-- Close icon (X) -->
							<svg
								class="h-6 w-6"
								xmlns="http://www.w3.org/2000/svg"
								fill="none"
								viewBox="0 0 24 24"
								stroke="currentColor"
							>
								<path
									stroke-linecap="round"
									stroke-linejoin="round"
									stroke-width="2"
									d="M6 18L18 6M6 6l12 12"
								/>
							</svg>
						{:else}
							<!-- Hamburger icon -->
							<svg
								class="h-6 w-6"
								xmlns="http://www.w3.org/2000/svg"
								fill="none"
								viewBox="0 0 24 24"
								stroke="currentColor"
							>
								<path
									stroke-linecap="round"
									stroke-linejoin="round"
									stroke-width="2"
									d="M4 6h16M4 12h16M4 18h16"
								/>
							</svg>
						{/if}
					</button>
				</div>
			</div>
		</div>
	</div>
	<!-- Mobile Menu -->
	{#if isOpen}
		<div class="sm:hidden">
			<div class="space-y-1 pb-3 pt-2">
				{#each navLinks as link}
					<a
						href={link.href}
						class="block border-l-4 py-2 pl-3 pr-4 {isActive(link.href)
							? 'border-indigo-500 bg-indigo-50 text-indigo-700'
							: 'border-transparent text-gray-600 hover:border-gray-300 hover:bg-gray-50 hover:text-gray-800'}"
						aria-current={isActive(link.href) ? 'page' : undefined}
					>
						{link.label}
					</a>
				{/each}
			</div>
		</div>
	{/if}
</nav>
