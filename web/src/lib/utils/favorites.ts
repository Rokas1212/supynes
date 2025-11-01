export async function getFavorites(userID: number | null): Promise<number[]> {
	if (!userID) {
		return [];
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

			return data.swings;
		}
		return [];
	} catch (error) {
		console.error('Failed to fetch favorites:', error);
		return [];
	}
}
