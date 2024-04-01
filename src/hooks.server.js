import { PUBLIC_BACKEND_ENDPOINT_HOST } from '$env/static/public';

//TODO handle, handleFetch, load and action's event,request,cookies,fetch are real
//mess. to find which is which check usage before use it 
export async function handle({ event, resolve }) {

	if (event.request.url.startsWith('http://localhost:5173/')) {

		let refreshTokenCookie = event.cookies.get('e-commerce-refresh-token')
		let accessTokenCookie = event.cookies.get('e-commerce-access-token')

		// console.log("refresh token :", refreshTokenCookie);
		// console.log("access token : ", accessTokenCookie);
		// console.log("new refresh token :", event.request.headers);


		// if (refreshTokenCookie == undefined) {
		// 	console.log("refresh token undefined");
		// 	// throw redirect(303, `/login`);
		// }
		// if (accessTokenCookie == undefined) {
		// 	console.log("access token undefined");
		// }

		if (refreshTokenCookie && !accessTokenCookie) {
			// console.log("1 refresh token :", event.request.headers);
			const req = event.fetch(PUBLIC_BACKEND_ENDPOINT_HOST + '/auth/renew-access-token', {
				method: 'GET',
				headers: {
					'Content-Type': 'application/json'
				}
			});
			const results = (await req).headers.getSetCookie();


			if (results) {
				for (let i = 0; i < results.length; i++) {
					var el = results[i];

					var cookieName = el.split(';')[0].split('=')[0];
					var cookieValue = el.split(';')[0].split('=')[1];
					var cookiePath = el.split(';')[1].split('=')[1];
					var cookieExpires = el.split(';')[3].split('=')[1];
					var cookieMaxAge = el.split(';')[4].split('=')[1];
					var cookieHttpOnly = el.split(';')[5] == 'HttpOnly' ? true : false;
					var cookieSecure = el.split(';')[6] == 'Secure' ? true : false;
					var cookieSameSite = el.split(';')[7].split('=')[1];

					event.cookies.set(cookieName, cookieValue, {
						path: cookiePath,
						expire: new Date(cookieExpires),
						maxAge: parseInt(cookieMaxAge),
						httpOnly: cookieHttpOnly,
						secure: cookieSecure,
						sameSite: cookieSameSite
					});
				}
			}
		}

	}
	const response = await resolve(event);
	return response;
}

//TODO: INTERNAL ERROR 500 HANDLE IN HERE LATER
// export function handleError({ event, error }) {
// 	console.error(error.stack);

// 	console.log(event)

// 	return {
// 		message: 'everything is fine',
// 		code: 'JEREMYBEARIMY'
// 	};
// }
