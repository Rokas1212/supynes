import type { PageLoad } from './$types';
import { error } from '@sveltejs/kit';

export const load: PageLoad = async ({ params, fetch }) => {
	const api_url = `/api/swings/${params.id}`;
	const photo_url = `/api/photos/${params.id}`;
	const review_url = `/api/reviews/${params.id}`;

	try {
		const [swingResp, photoResp, reviewResp] = await Promise.all([
			fetch(api_url),
			fetch(photo_url),
			fetch(review_url)
		]);

		if (!swingResp.ok || !photoResp.ok || !reviewResp.ok) {
			console.error(`API Error: ${swingResp.status} - ${swingResp.statusText}`);
			throw error(swingResp.status, 'Failed to load swing data.');
		}

		const swing = await swingResp.json();
		const photoData = await photoResp.json();
		const photos = photoData.photos || [];
		const reviewData = await reviewResp.json();
		console.log('Review data fetched:', reviewData);
		const reviews = (reviewData.reviews ?? []).map((r: any) => ({
			ID: r.ID,
			SwingID: r.SwingID,
			UserID: r.UserID,
			Rating: r.Rating,
			Title: r.Title,
			Body: r.Body,
			VisitedOn: r.VisitedOn,
			UpdatedAt: r.UpdatedAt ?? ''
		}));

		return {
			swing,
			photos,
			reviews
		};
	} catch (e) {
		console.error('Uncaught error during load:', e);
		throw error(500, 'Internal Server Error during data fetching.');
	}
};
