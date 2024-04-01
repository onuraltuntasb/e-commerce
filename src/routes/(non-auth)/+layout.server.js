import { PUBLIC_BACKEND_ENDPOINT_HOST } from "$env/static/public";

export const prerender = false;

export const load = async ({ cookies, fetch: sFetch }) => {

	var accessTokenCookie = cookies.get('e-commerce-access-token');

	const req = sFetch(PUBLIC_BACKEND_ENDPOINT_HOST + '/categories', {
		headers: {
			'Content-Type': 'application/json'
		}
	});

	const stat = (await req).status;
	const body = await req.then(r => r.json());



	if (accessTokenCookie) {
		const claims = atob(accessTokenCookie.split('.')[1]);
		const email = JSON.parse(claims)?.name.split(' ')[2];

		return { status: 200, data: { categories: body, email: email, alertMsg: '' }, alert: 'success' };
	} else {
		if (stat >= 200 && stat <= 299) {
			return { status: 200, data: { categories: body, alertMsg: '' }, alert: 'success' };
			//send alert to client before redirect
		}
		if (stat >= 400 && stat <= 599) {
			return { status: 200, data: { categories: body, alertMsg: '' }, alert: 'fail' };
		}
	}

};




