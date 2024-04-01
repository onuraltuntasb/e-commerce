
import { browser } from "$app/environment";
import { PUBLIC_FRONTEND_ENDPOINT_HOST } from "$env/static/public";
import { error, } from "@sveltejs/kit";


export const parseJwtCookies = (results, cookies) => {
    if (results) {
        for (let i = 0; i < results.length; i++) {
            var el = results[i];

            var cookieName = el.split(';')[0].split('=')[0];
            var cookieValue = el.split(';')[0].split('=')[1];
            var cookiePath = el.split(';')[1].split('=')[1];
            var cookieExpires = el.split(';')[3].split('=')[1];
            var cookieMaxAge = el.split(';')[4].split('=')[1];
            var cookieHttpOnly = el.split(';')[5] == 'HttpOnly' ? true : false;
            var cookieSecure = el.split(';')[6] == 'Secure' ? true : false;
            var cookieSameSite = el.split(';')[7].split('=')[1];

            cookies.set(cookieName, cookieValue, {
                path: cookiePath,
                expire: new Date(cookieExpires),
                maxAge: parseInt(cookieMaxAge),
                httpOnly: cookieHttpOnly,
                secure: cookieSecure,
                sameSite: cookieSameSite
            });
        }
    }
}

export const getUserIdFromAccessToken = (cookies) => {
    var accessTokenCookie = cookies.get('e-commerce-access-token');

    if (accessTokenCookie) {
        const claims = atob(accessTokenCookie.split('.')[1]);
        const userId = JSON.parse(claims)?.sub;
        return parseInt(userId);
    }
    else {
        throw error(401, "there is no access token!")
    }
}

export const authCheck = (cookies) => {
    var refreshTokenCookie = cookies.get('e-commerce-refresh-token');
    if (refreshTokenCookie == undefined) {
        return "goLogin";
    }

    return "";
}

export const redirectWithPageRefresh = (location) => {
    if (browser) {
        window.location.href = PUBLIC_FRONTEND_ENDPOINT_HOST + location;
    }
}


export const getProductSearchableList = (productListJSON) => {


    let productSearchableList = [];


    productListJSON?.forEach(item => {

        if (item?.product?.searchable) {
            let searchableJSON = item?.product?.searchable;
            let sMap = {};

            for (let [key, value] of Object.entries(searchableJSON)) {
                // console.log(`${key}: ${JSON.stringify(value)}`);

                if (typeof value === "string" && value !== null) {
                    if (sMap[key]) {
                        sMap[key].push(value)
                    } else {
                        let ar = [];
                        ar.push(value);
                        sMap[key] = ar;
                    }
                }


                else if (Array.isArray(value)) {
                    // console.log("value : ", value);
                    value.forEach(el => {
                        if ((typeof el === "object") && (el !== null)) {
                            if (sMap[key]) {
                                sMap[key].push(el[key])
                            } else {
                                let ar = [];
                                ar.push(el[key]);
                                sMap[key] = ar;
                            }

                        } else if (typeof el === "string" && el !== null) {
                            if (sMap[key]) {
                                sMap[key].push(el)
                            } else {
                                let ar = [];
                                ar.push(el);
                                sMap[key] = ar;

                            }
                        }
                    })
                }
                // console.log("arr values : ", JSON.stringify(sMap));

            }
            productSearchableList.push(sMap);

        }
    })

    return productSearchableList;

}


// console.log("rProductSearchableList : ", JSON.stringify(rProductSearchableList));

export const getUnion = (a, b) => {
    var union = [...new Set([...a, ...b])];
    return union;
}

export const getUnionedAllObjList = (productSearchableList) => {


    let unionedAllObjList = [];

    if (productSearchableList?.length == 0) {
        return [];
    }

    const half = Math.ceil(productSearchableList.length / 2);

    const firstHalf = productSearchableList.slice(0, half)
    const secondHalf = productSearchableList.slice(half)

    // console.log(" \n productSearchableList", JSON.stringify(productSearchableList));
    // console.log(" \n firstHalf", JSON.stringify(firstHalf));
    // console.log(" \n secondHalf", JSON.stringify(secondHalf));
    //TODO: 0 or 1 product condition

    let iteratorHalf = [];
    let otherHalf = [];

    if (firstHalf.length >= secondHalf.length) {
        iteratorHalf = firstHalf;
        otherHalf = secondHalf;
    } else {
        iteratorHalf = secondHalf;
        otherHalf = firstHalf;
    }

    // console.log(" \n iteratorHalf :", JSON.stringify(iteratorHalf))

    for (let i = 0; i < iteratorHalf.length; i++) {
        const iteratorHalfEl = iteratorHalf[i];
        const otherHalfEl = otherHalf[i] ? otherHalf[i] : {};

        //  console.log(" \n iteratorHalfEl :", JSON.stringify(iteratorHalfEl))
        // console.log(" \n otherHalfEl :", JSON.stringify(otherHalfEl))

        let iteratorHalfObj = {};
        let otherHalfObj = {};

        // console.log(" \n Object.entries(iteratorHalfEl) :", JSON.stringify(Object.entries(iteratorHalfEl)))
        // console.log(" \n Object.entries(otherHalfEl) :", JSON.stringify(Object.entries(otherHalfEl)))

        if (Object.entries(iteratorHalfEl).length >= Object.entries(otherHalfEl).length) {
            iteratorHalfObj = iteratorHalfEl;
            otherHalfObj = otherHalfEl;
        } else {
            iteratorHalfObj = otherHalfEl;
            otherHalfObj = iteratorHalfEl;
        }

        // console.log(" \n iteratorHalfObj :", JSON.stringify(iteratorHalfObj))


        let unionedRowObj = {};

        let newIteratorHalfObj = {};
        let tempIteratorHalfObj = { ...iteratorHalfObj };

        // console.log(" \n iteratorHalfObj :", JSON.stringify(iteratorHalfObj))
        // console.log(" \n otherHalfObj :", JSON.stringify(otherHalfObj))

        for (let i = 0; i < Object.entries(otherHalfObj).length; i++) {
            const otherHalfObjKey = (Object.keys(otherHalfObj)[i])

            if (iteratorHalfObj[otherHalfObjKey]) {
                let newIteratorHalfObjInner = {};
                newIteratorHalfObjInner[otherHalfObjKey] = iteratorHalfObj[otherHalfObjKey];

                newIteratorHalfObj = { ...newIteratorHalfObj, ...newIteratorHalfObjInner };

                delete tempIteratorHalfObj[otherHalfObjKey];

            }
        }

        newIteratorHalfObj = { ...newIteratorHalfObj, ...tempIteratorHalfObj };
        iteratorHalfObj = { ...newIteratorHalfObj };

        // console.log(" \n iteratorHalfObj :", JSON.stringify(iteratorHalfObj))
        // console.log(" \n otherHalfObj :", JSON.stringify(otherHalfObj))

        for (let i = 0; i < Object.entries(iteratorHalfObj).length; i++) {

            // console.log(" \n Object.entries(iteratorHalfObj)[i] :", JSON.stringify(Object.values(iteratorHalfObj)[i]))

            const iteratorHalfObjKey = (Object.keys(iteratorHalfObj)[i])
            const iteratorHalfObjValue = (Object.values(iteratorHalfObj)[i])

            const otherHalfObjKey = (Object.keys(otherHalfObj)[i])
            const otherHalfObjValue = (Object.values(otherHalfObj)[i])

            if (iteratorHalfObjKey == otherHalfObjKey) {
                // console.log("gir");
                let unionedValue = getUnion(iteratorHalfObjValue, otherHalfObjValue);
                let unionedObj = {};
                unionedObj[iteratorHalfObjKey] = unionedValue;

                // console.log(" \n unionedObj :", JSON.stringify(unionedObj))
                unionedRowObj = { ...unionedRowObj, ...unionedObj }


            } else {
                let unionedObj = {};
                unionedObj[iteratorHalfObjKey] = iteratorHalfObjValue;
                unionedRowObj = { ...unionedRowObj, ...unionedObj }
                unionedObj = {};

                if (otherHalfObjKey && otherHalfObjValue) {
                    unionedObj[otherHalfObjKey] = otherHalfObjValue;
                    unionedRowObj = { ...unionedRowObj, ...unionedObj }
                }

            }


            // console.log(" \n unionedAllObj :", JSON.stringify(unionedAllObj))
            // console.log(" \n iteratorHalfObjKey :", JSON.stringify(iteratorHalfObjKey))
            // console.log(" \n iteratorHalfObjValue :", JSON.stringify(iteratorHalfObjValue))
            // console.log(" \n otherHalfObjKey :", JSON.stringify(otherHalfObjKey))
            // console.log(" \n otherHalfObjValue :", JSON.stringify(otherHalfObjValue))
        }
        unionedAllObjList.push(unionedRowObj);
        // console.log(" \n unionedRowObj :", JSON.stringify(unionedRowObj))

    }
    // console.log(" \n unionedAllObjList :", (unionedAllObjList))

    if (unionedAllObjList?.length == 1) {
        return unionedAllObjList;
    }
    return getUnionedAllObjList(unionedAllObjList);
}
