<script lang="ts">
	let name = '';
	let address = '';
	let lat = '';
	let lng = '';
	let city = '';
	let err = '';
	async function handleSubmit(e: Event) {
		e.preventDefault();
		err = '';

		try {
			const resp = await fetch('/api/playground', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					UserID: 1,
					Name: name,
					Address: address,
					Lat: parseFloat(lat),
					Lng: parseFloat(lng),
					City: city
				})
			});

			if (resp.ok) {
				const data = await resp.json();
				console.log(data);
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
<!-- TODO: add redirect to playgrounds or the specific playground-->
<!-- TODO: allow to create swings together with playground -->
<!-- TODO: style form -->
<form on:submit={handleSubmit} class="space-y-6">
	<div>
		<label for="name">Name</label>
		<input id="name" name="name" required bind:value={name} />
	</div>
	<div>
		<label for="address">Address</label>
		<input id="address" name="address" required bind:value={address} />
	</div>
	<div>
		<label for="lat">Lat</label>
		<input id="lat" name="lat" required bind:value={lat} />
	</div>
	<div>
		<label for="lng">Lng</label>
		<input id="lng" name="lng" required bind:value={lng} />
	</div>
	<div>
		<label for="city">City</label>
		<input id="city" name="city" required bind:value={city} />
	</div>
	<button type="submit"> Add </button>
</form>
