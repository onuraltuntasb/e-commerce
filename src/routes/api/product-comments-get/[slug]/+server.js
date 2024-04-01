import { PUBLIC_BACKEND_ENDPOINT_HOST } from "$env/static/public";

export async function GET({ fetch, url }) {

    let urlArr = url.href.split('/');

    let productId = urlArr[urlArr.length - 1];

    const req = fetch(PUBLIC_BACKEND_ENDPOINT_HOST + '/product-comments-get/' + productId);

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