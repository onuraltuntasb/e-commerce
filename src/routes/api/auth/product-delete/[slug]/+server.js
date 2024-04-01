import { PUBLIC_BACKEND_ENDPOINT_HOST } from '$env/static/public';
import { getUserIdFromAccessToken } from '$lib/utils.js';

export async function DELETE({ fetch, url, cookies }) {
    let urlArr = url.href.split('/');

    let id = urlArr[urlArr.length - 1];
    let userId = getUserIdFromAccessToken(cookies)

    const req = fetch(
        `${PUBLIC_BACKEND_ENDPOINT_HOST}/auth/product-delete/${id}/${userId}`,
        {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json'
            }
        }
    );

    const stat = (await req).status;
    const m = await req.then((r) => r.json());

    console.log("stat :", stat);

    let resData = {};

    if (stat >= 200 && stat <= 299) {
        resData = { status: stat, msg: m, alert: 'success' };
    }
    if (stat >= 400 && stat <= 599) {
        resData = { status: stat, msg: m, alert: 'fail' };
    }

    return new Response(JSON.stringify(resData), {
        headers: {
            'Content-Type': 'application/json'
        }
    });
}