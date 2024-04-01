import { PUBLIC_BACKEND_ENDPOINT_HOST } from '$env/static/public';
import { getUserIdFromAccessToken } from '$lib/utils.js';

export async function POST({ fetch: sFetch, cookies, request }) {

    const content = await request.json();

    // console.log("content## :", content);

    let userId = getUserIdFromAccessToken(cookies)

    // console.log("userId :", userId)

    //TODO: parse error "o"
    const req = sFetch(
        `${PUBLIC_BACKEND_ENDPOINT_HOST}/auth/product-add/${userId}`,
        {
            method: 'POST',
            body: content,
            headers: {
                'Content-Type': 'application/json'
            }
        }
    );

    const stat = (await req).status;
    const body = await req.then((r) => r.json());

    // console.log("stat :", stat);

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
