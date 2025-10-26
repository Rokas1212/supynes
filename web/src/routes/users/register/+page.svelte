<script lang="ts">
	import { goto } from '$app/navigation';

	let email = '';
	let password = '';
	let displayName = '';
	let err = '';

	async function handleSubmit(e: Event) {
		e.preventDefault();
		err = '';

		try {
			const resp = await fetch('/api/register', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({ email: email, password: password, displayName: displayName })
			});

			if (resp.ok) {
				const data = await resp.json();
				console.log('Registration successful', data);

				goto('/');
			} else {
				const data = await resp.json();
				err = data.error || 'registration failed';
			}
		} catch (error) {
			console.error('Failed to register:', error);
			err = 'Error, plz try again!';
		}
	}
</script>

<div class="flex min-h-full flex-col justify-center px-6 py-12 lg:px-8">
	<div class="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
		<form on:submit={handleSubmit} class="space-y-6">
			<div>
				<label for="email" class="block text-sm font-medium text-gray-900">Email address</label>
				<input
					id="email"
					name="email"
					type="email"
					required
					bind:value={email}
					autocomplete="email"
					class="mt-2 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500"
				/>
			</div>

			<div>
				<label for="displayName" class="block text-sm font-medium text-gray-900">Username</label>
				<input
					id="displayName"
					name="displayName"
					type="text"
					required
					bind:value={displayName}
					class="mt-2 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500"
				/>
			</div>

			<div>
				<label for="password" class="block text-sm font-medium text-gray-900">Password</label>
				<input
					id="password"
					name="password"
					type="password"
					required
					bind:value={password}
					autocomplete="current-password"
					class="mt-2 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500"
				/>
			</div>

			<button
				type="submit"
				class="w-full cursor-pointer rounded-md bg-blue-600 px-4 py-2 font-semibold text-white"
			>
				Register
			</button>
		</form>
	</div>
</div>
