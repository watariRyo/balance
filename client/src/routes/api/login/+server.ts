import { error, json, type RequestHandler } from '@sveltejs/kit';
import { apiClient, ApiError, createCustomApiError } from '../apiClient';
import { BASE_URL } from '$env/static/private';
import { HttpStatusCodes400 } from '../httpStatusCode';

export const POST: RequestHandler = async ({ request }) => {
	const data = await request.json();
	if (!data.userID) {
		const message = 'Username is required!';
		return json({
			ok: false,
			grpcResponse: createCustomApiError(message, HttpStatusCodes400.BAD_REQUEST)
		});
	}
	if (!data.password) {
		const message = 'Password is required!';
		return json({
			ok: false,
			grpcResponse: createCustomApiError(message, HttpStatusCodes400.BAD_REQUEST)
		});
	}

	const response = await apiClient(`${BASE_URL}/login`, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json'
		},
		body: JSON.stringify({
			data
		})
	});

	if (response instanceof ApiError) {
		return json({
			ok: false,
			grpcResponse: response
		});
	} else {
		return json({
			ok: true,
			grpcResponse: response
		});
	}
};
