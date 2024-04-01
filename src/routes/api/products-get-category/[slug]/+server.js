import { PUBLIC_BACKEND_ENDPOINT_HOST } from "$env/static/public";
import { getProductSearchableList, getUnionedAllObjList } from "$lib/utils";

export async function POST({ fetch, url, request }) {

    const content = await request.json();

    // console.log("content : ", JSON.stringify(content));


    //TODO: 1.  post content olarak searchProperty göndermek gerek
    //TODO: 2.  ama önce backend uygun bi mockup data yapıp o dataya göre search
    //logic yazılmalı
    //TODO: 3.  backend de herbir attribute için postgres query nasıl yazılır
    //veya generic nasıl yazılır denemesi yapılıp stabilleştirilmeli 
    //TODO: sonra mockup search logic'e göre query sonuçları nasıl
    //birleştirilebilir o düşünülmeli belki concatination ile tek bir query de
    //tüm sonuçlar alınabilir hatta daha güzel temiz olur

    let urlArr = url.href.split('/');
    let categoryName = urlArr[urlArr.length - 1];

    const req = fetch(PUBLIC_BACKEND_ENDPOINT_HOST + '/products-get/category/' + categoryName,
        {
            method: 'POST',
            body: JSON.stringify(content),
            headers: {
                'Content-Type': 'application/json'
            }
        }
    );

    const stat = (await req).status;
    console.log("stat : ", stat)
    var m = await req.then((r) => r.json());
    let resData = {};


    console.log("product get m:", (m));

    let parsedProducts = [];
    parsedProducts = m ? m?.map(el => el?.product) : [];

    let rProductSearchableList = getProductSearchableList((parsedProducts));

    let res = getUnionedAllObjList(rProductSearchableList);
    // console.log(" \n res :", (res))

    let mData = {};
    mData = { productList: m ? m : [], searchablePropertyList: res }


    if (stat >= 200 && stat <= 299) {
        resData = { status: stat, msg: mData, alert: 'success' };
    }
    if (stat >= 400 && stat <= 599) {
        resData = { status: stat, msg: mData, alert: 'fail' };
    }

    return new Response(JSON.stringify(resData), {
        headers: {
            'Content-Type': 'application/json'
        }
    });
}   