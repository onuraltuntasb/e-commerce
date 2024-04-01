import { PUBLIC_BACKEND_ENDPOINT_HOST } from "$env/static/public";
import { getUserIdFromAccessToken } from "$lib/utils";

export async function GET({ fetch, cookies }) {

    let userId = getUserIdFromAccessToken(cookies)

    const req = fetch(PUBLIC_BACKEND_ENDPOINT_HOST + '/auth/products-get-seller/' + userId);

    const stat = (await req).status;
    // console.log("stat : ", stat)
    const body = await req.then((r) => r.json());
    let resData = {};


    // console.log("product get m:", m);



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