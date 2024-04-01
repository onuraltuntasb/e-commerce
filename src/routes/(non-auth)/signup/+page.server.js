import { PUBLIC_BACKEND_ENDPOINT_HOST } from '$env/static/public';
import { signupInitialData } from '$lib/initialDatas.js';
import { signupSchema } from '$lib/schemas.js';
import { parseJwtCookies } from '$lib/utils';
import { error, fail } from '@sveltejs/kit';
import { superValidate } from 'sveltekit-superforms/server';


export const load = async () => {
	if (!signupInitialData) {
		throw error(404, 'Not found');
	}

	const form = await superValidate(signupInitialData, signupSchema);


	return { form };
};

export const actions = {
	default: async ({ request, cookies, fetch }) => {
		const form = await superValidate(request, signupSchema);
		if (!form.valid) {
			return fail(400, { form });
		}

		//TODO: server-side form validation is necessary
		await new Promise(resolve => setTimeout(resolve, 3000)); // 3 sec

		const req = fetch(PUBLIC_BACKEND_ENDPOINT_HOST + '/signup', {
			method: 'POST',
			body: JSON.stringify(form.data),
			headers: {
				'Content-Type': 'application/json'
			}
		});


		const results = (await req).headers.getSetCookie();

		parseJwtCookies(results, cookies);
		//TODO setCookies in one line later just like set headers

		const stat = (await req).status;
		const body = await req.then(r => r.json());

		if (stat >= 200 && stat <= 299) {
			form.message = { status: stat, msg: body, alert: 'success' };
		}
		if (stat >= 400 && stat <= 599) {
			form.message = { status: stat, msg: body, alert: 'fail' };
		}

		return { form };

	}
};
