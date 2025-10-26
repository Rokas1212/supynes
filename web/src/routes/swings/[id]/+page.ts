import type { PageLoad } from './$types';
import { error } from '@sveltejs/kit';

export const load: PageLoad = async ({ params, fetch }) => {
	const api_url = `/api/swings/${params.id}`;
	const photo_url = `/api/photos/${params.id}`;

	try {
		const [swingResp, photoResp] = await Promise.all([fetch(api_url), fetch(photo_url)]);

		if (!swingResp.ok || !photoResp.ok) {
			console.error(`API Error: ${swingResp.status} - ${swingResp.statusText}`);
			throw error(swingResp.status, 'Failed to load swing data.');
		}

		const swing = await swingResp.json();
		const photoData = await photoResp.json();
		const photos = photoData.photos || [];
		console.log('Loaded swing:', swing);
		console.log('Loaded photos:', photos);

		return { swing, photos };
	} catch (e) {
		console.error('Uncaught error during load:', e);
		throw error(500, 'Internal Server Error during data fetching.');
	}
};
