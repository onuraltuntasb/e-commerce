
import { PUBLIC_BACKEND_ENDPOINT_HOST } from '$env/static/public';
import { parseJwtCookies } from '$lib/utils.js';

export async function GET({ fetch, cookies }) {

    console.log("hit lg get!")

    const req = fetch(PUBLIC_BACKEND_ENDPOINT_HOST + '/auth/logout', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    });

    const results = (await req).headers.getSetCookie();

    //NOTE: setCookies in one line later just like set headers
    parseJwtCookies(results, cookies);

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