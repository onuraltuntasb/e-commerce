import { writable } from 'svelte/store';

export const storeSelectedAddress = writable({});
export const storeAddressList = writable([]);

export const storeSelectedProduct = writable({});
export const storeProductList = writable([]);

export const storeAccountInfos = writable({});

export const storeSelectedMainCategory = writable("");
export const storeCategories = writable({});

export const storeParsedCategories = writable([])

export const storeSelectedCategory = writable("");
export const storePreElId = writable("");

export const storeAccountState = writable("");

export const storeSelectedSearchProps = writable({});

export const storeSearchablePropertyListParsed = writable([]);

export const storeCard = writable([]);



