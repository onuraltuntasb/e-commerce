import { PUBLIC_BACKEND_ENDPOINT_HOST } from "$env/static/public";

export async function GET({ fetch, url }) {

    let urlArr = url.href.split('/');
    let districtId = urlArr[urlArr.length - 1];

    const req = fetch(PUBLIC_BACKEND_ENDPOINT_HOST + '/auth/neighborhood/' + districtId);

    const stat = (await req).status;
    const m = await req.then((r) => r.json());
    let resData = {};

    if (stat >= 200 && stat <= 299) {
        resData = { status: stat, msg: m, alert: 'success' };
    }
    if (stat >= 400 && stat <= 599) {
        resData = { status: stat, msg: m, alert: 'success' };
    }

    return new Response(JSON.stringify(resData), {
        headers: {
            'Content-Type': 'application/json'
        }
    });
}   