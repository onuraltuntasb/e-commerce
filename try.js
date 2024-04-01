const productJSON = {
    "product": {
        "searchable": {
            "sex": [
                {
                    "sex": "male",
                    "stock": 5
                }
            ],
            "name": "Camper Runner K21 Siyah Erkek Sneaker 000000000101907225",
            "brand": "Camper",
            "color": [
                {
                    "color": "black",
                    "stock": 0
                },
                {
                    "color": "white",
                    "stock": 10
                }
            ],
            "price": "5.499",
            "shoeSize": [
                {
                    "size": "40",
                    "stock": 5
                },
                {
                    "size": "41",
                    "stock": 5
                }
            ],
            "categories": [
                "Man Sneaker",
                "Man Casual Shoe",
                "Man Shoe",
                "Shoe & Purse",
                "Fashion"
            ]
        }
    }
}

if (productJSON?.product?.searchable) {
    let searchableJSON = productJSON?.product?.searchable;
    let arrSearchableJSON = Object.values(searchableJSON);

    //arrSearchableJSON.forEach(el => {
    //console.log(el);
    //});

    let sMap = {};




    for (let [key, value] of Object.entries(searchableJSON)) {
        console.log(`${key}: ${JSON.stringify(value)}`);
        // if(isArray
        // sMap[key] = 
    }

}




