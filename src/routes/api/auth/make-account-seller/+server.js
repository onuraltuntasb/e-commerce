import { PUBLIC_BACKEND_ENDPOINT_HOST } from '$env/static/public';
import { getUserIdFromAccessToken } from '$lib/utils.js';

export async function GET({ fetch, cookies }) {

    let userId = getUserIdFromAccessToken(cookies)

    console.log("make-account-seller  user id : ", userId)

    const req = fetch(
        `${PUBLIC_BACKEND_ENDPOINT_HOST}/auth/make-account-seller/${userId}`,
        {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            }
        }
    );

    const stat = (await req).status;
    const body = await req.then((r) => r.json());

    console.log("stat :", stat);

    let resData = {};

    if (stat >= 200 && stat <= 299) {
        resData = { status: stat, msg: body, alert: 'success' };
    }
    if (stat >= 400 && stat <= 599) {
        resData = { status: stat, msg: body, alert: 'fail' };
    }

    return new Response(JSON.stringify(resData), {
        headers: {
            'Content-Type': 'application/json'
        }
    });
}
