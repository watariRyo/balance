import { json, type RequestHandler } from '@sveltejs/kit';
import { apiClient, ApiError } from '../apiClient';
import { BASE_URL } from '$env/static/private';

export const GET: RequestHandler = async ({ cookies }) => {
	const accessToken = cookies.get('access_token');
	const sessionID = cookies.get('session_id');

	const response = await apiClient(`${BASE_URL}/api/get/user`, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
			Authorization: accessToken || ''
		},
		body: JSON.stringify({
			session_id: sessionID
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

export const POST: RequestHandler = async ({ request }) => {
	const data = await request.json();

	const response = await apiClient(`${BASE_URL}/api/register/user`, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json'
		},
		body: JSON.stringify({
			user_id: data.userID,
			password: data.password
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
