import { PUBLIC_BACKEND_ENDPOINT_HOST } from "$env/static/public";
import { getUserIdFromAccessToken } from "$lib/utils";

export async function PUT({ fetch, cookies, url, request }) {

    const content = await request.json();

    let userId = getUserIdFromAccessToken(cookies)
    let urlArr = url.href.split('/');
    let id = urlArr[urlArr.length - 1];

    const req = fetch(`${PUBLIC_BACKEND_ENDPOINT_HOST}/auth/product-update/${id}/${userId}`,
        {
            method: 'PUT',
            body: content,
            headers: {
                'Content-Type': 'application/json'
            }
        });

    const stat = (await req).status;
    const m = await req.then((r) => r.json());
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