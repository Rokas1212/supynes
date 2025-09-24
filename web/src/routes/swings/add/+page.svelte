<script lang="ts">
	let name = '';
	let address = '';
	let lat = '';
	let lng = '';
	let city = '';
	let seatCount: number;
	let maxWeightKg: number;
	let isAccessible: boolean;

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
					MaxWeightKg: maxWeightKg ? maxWeightKg : null
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
<!-- TODO: add redirect to swings or the specific swing-->
<!-- TODO: style form -->
<form on:submit={handleSubmit} class="space-y-6">
	<div>
		<label for="name">Name</label>
		<input id="name" name="name" type="text" required bind:value={name} />
	</div>
	<div>
		<label for="lat">Lat: </label>
		<input id="lat" name="lat" type="number" step="any" required bind:value={lat} />
	</div>
	<div>
		<label for="lng">Lng: </label>
		<input id="lng" name="lng" type="number" step="any" required bind:value={lng} />
	</div>
	<div>
		<label for="address">Full Street: </label>
		<input id="address" name="address" type="text" required bind:value={address} />
	</div>
	<div>
		<label for="city">City: </label>
		<input id="city" name="city" type="text" required bind:value={city} />
	</div>
	<div>
		<label for="seatCount">Seat Count: </label>
		<input id="seatCount" name="seatCount" type="number" bind:value={seatCount} />
	</div>
	<div>
		<label for="">Max Weight: </label>
		<input id="" name="" type="number" bind:value={maxWeightKg} />
	</div>
	<div>
		<label for="isAccessible">Is Accessible?</label>
		<input id="isAccessible" name="isAccessible" type="checkbox" bind:checked={isAccessible} />
	</div>
	<button class="hover:bg-gray-50" type="submit"> Add </button>
</form>
