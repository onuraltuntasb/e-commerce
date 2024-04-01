import { PUBLIC_BACKEND_ENDPOINT_HOST } from "$env/static/public";

export async function GET({ fetch }) {

    const req = fetch(PUBLIC_BACKEND_ENDPOINT_HOST + '/auth/city');

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

// /make-address-primary/ids?user-id={userId}&address-id={addressId}