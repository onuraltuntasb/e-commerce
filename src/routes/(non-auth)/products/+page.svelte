<script>
	import {
		storeCard,
		storeCategories,
		storeParsedCategories,
		storeSearchablePropertyListParsed,
		storeSelectedCategory,
		storeSelectedMainCategory,
		storeSelectedSearchProps
	} from '../../../stores';
	import { PUBLIC_FRONTEND_ENDPOINT_HOST } from '$env/static/public';
	import Swal from 'sweetalert2';
	import { redirectWithPageRefresh } from '$lib/utils';
	import GenericSearch from './GenericSearch.svelte';
	import { onDestroy, onMount, tick } from 'svelte';
	import { browser } from '$app/environment';
	import TreeViewTry from './TreeViewTry.svelte';
	import { parse } from 'svelte/compiler';

	import store2 from 'store2';

	let categories = [];
	let parsedCategories = {};
	let categoryArr = [];
	let searchablePropertyList = [];
	let searchablePropertyListParsed = [];
	let searchableSendData = [];
	let searchableData = [];
	let productList = [];
	let splittedProductList = [];
	let firstMount = false;
	let categoryChanged = false;
	let parsedChanged = false;

	onMount(() => {
		if (browser) {
			if (store2.get('categories')) {
				categories = store2.get('categories');

				if (store2.get('mainCategory')) {
					let cat = store2.get('mainCategory');
					$storeSelectedMainCategory = cat;
				} else {
					$storeSelectedMainCategory = 'Fashion';
					store2.set('mainCategory', 'Fashion');
				}

				if (store2.get('category')) {
					let cat = store2.get('category');
					$storeSelectedCategory = cat;
					// console.log('cat1 : ', cat);
				} else {
					store2.set('category', 'Fashion');
					$storeSelectedCategory = 'Fashion';
				}
			}
			firstMount = true;
		}
	});

	$: {
		if ($storeSelectedCategory) {
			if (browser) {
				if (store2.get('categories')) {
					categories = store2.get('categories');
					// console.log('abc1 : ', categories);
					//TODO: FİRST MAKE ALL LEAF'S EXPAND FALSE
					parseCategories();
					// console.log('abc : ', parsedCategories);
					let cat = [...categories];
					let catTemp = cat?.map((el) => {
						if (el.name == $storeSelectedCategory) {
							return (el = parsedCategories);
						} else {
							return el;
						}
					});
					// console.log('cat2 :', catTemp);
					// TODO:localstorage not updating!!

					setTimeout(() => {
						store2.set('categories', catTemp);
					}, 1000);
					categories = [...cat];

					// console.log('categories :', categories);
					// console.log('category changed!');
				}
			}

			categoryChanged = true;
			searchableData = [];
			searchHandler(true);
		}
	}

	$: {
		if ($storeSearchablePropertyListParsed) {
			parsedChanged = true;
			if (browser) {
				if (store2.get('searchablePropertyList')) {
					searchableData = [];
					parseSearchableToSendData();
					getProductList();
				}
			}
		}
	}

	const parseCategories = () => {
		// console.log('$categ ', categories);
		// console.log('$storeSelectedMainCategory ', $storeSelectedMainCategory);
		// console.log('$storeSelectedCategory ', $storeSelectedCategory);
		if (categories && categories.length > 0) {
			categories.forEach((el) => {
				// console.log('el :', el);
				// console.log('$storeSelectedMainCategory ', $storeSelectedMainCategory);

				if (el.name == $storeSelectedMainCategory) {
					//HACK: another code enters between recursive functions and
					//mutate el. so run recursives on macro event loop sequentialy
					traverseTree(el);
					// console.log('categoryArr :', categoryArr);
					traverseTreeMake(el);

					// console.log('parseCategoriesEl :', el);

					parsedCategories = el;
				}
			});
		}
		// console.log('parsedCategories :', parsedCategories);
	};

	const traverseTree = (node) => {
		if (node.children.length === 0) {
			return;
		}
		categoryArr.push(node.name);
		if (node.name == $storeSelectedCategory) {
			return;
		}
		node.children.forEach((child) => traverseTree(child));
	};

	const traverseTreeMake = (node) => {
		if (node.children.length === 0) {
			return;
		}
		categoryArr.forEach((el) => {
			if (node.name === el) {
				node.expanded = true;
			}
		});
		node.children.forEach((child) => traverseTreeMake(child));
	};

	const searchHandler = () => {
		if ($storeSelectedCategory) {
			if (browser) {
				if (!firstMount) {
					searchableSendData = [];
				}
			}
			//TODO: GEÇ OLDU AMA ZATEN ZOR BI LOGICMIŞ CHAININGLERI HAFIFE ALMA
			getProductList();
		}
	};

	const getProductList = async () => {
		searchablePropertyListParsed = [];
		searchablePropertyList = [];
		productList = [];

		const res = await fetch(
			`${PUBLIC_FRONTEND_ENDPOINT_HOST}/api/products-get-category/${$storeSelectedCategory}`,
			{
				method: 'POST',
				body: JSON.stringify({ ...searchableSendData }),
				headers: {
					'Content-Type': 'application/json'
				}
			}
		);

		const body = await res.json();
		if (body?.status == 401) {
			redirectWithPageRefresh('/login');
		}
		if (body?.alert == 'success') {
			productList = body?.msg?.productList;

			// console.log('product list be :', productList);

			splitArrayIntoChunks(3, [...productList]);

			if (productList?.length == 0) {
				// console.log('remove!!');
				if (browser) {
					localStorage.removeItem('searchablePropertyList');
				}
			}

			// console.log('productList', productList);
			searchablePropertyList = body?.msg?.searchablePropertyList;
			// console.log('searchablePropertyList', searchablePropertyList);

			if (searchablePropertyList && searchablePropertyList.length > 0) {
				for (let [key, value] of Object.entries(searchablePropertyList[0])) {
					let tempObj = {};
					value = value.map((el) => (el = { name: el, checked: false }));
					tempObj[key] = value;
					searchablePropertyListParsed.push(tempObj);
					// console.log(`${key}: ${JSON.stringify(value)}`);
				}
			}

			// console.log('searchablePropertyList :', searchablePropertyList);
			let temp = store2.get('searchablePropertyList');
			// console.log('temp :', temp);

			if (browser) {
				if (firstMount) {
					if (
						store2.get('searchablePropertyList') == null ||
						store2.get('searchablePropertyList')?.length == 0
					) {
						parseSearchableToSendData();
						store2.set('searchablePropertyList', searchablePropertyListParsed);
					} else {
						parseSearchableToSendData();
					}
				} else {
					// console.log('not first mount!');
					if (categoryChanged) {
						// console.log('category changed!!', categoryChanged);
						parseSearchableToSendData();
						store2.set('searchablePropertyList', searchablePropertyListParsed);
					}
				}
			}
		} else if (body?.alert == 'fail') {
			Swal.fire({
				position: 'bottom-end',
				title: 'Product list fetched fail!',
				color: 'red',
				showConfirmButton: false,
				timer: 1500
			});
		}

		firstMount = false;
		categoryChanged = false;
		parsedChanged = false;
	};

	const splitArrayIntoChunks = (chunkSize, pL) => {
		if (pL && pL?.length > 0) {
			splittedProductList = pL.reduce(
				(productList, item, idx) =>
					(productList[(idx / chunkSize) | 0] ??= []).push(item) && productList,
				[]
			);
			// console.log('splittedProductList :', splittedProductList);
		}
	};

	const parseSearchableToSendData = () => {
		searchableSendData = [];
		if (browser) {
			if (store2.get('searchablePropertyList')) {
				let tempArr = [];
				let tempObj = {};
				searchableData = store2.get('searchablePropertyList');

				if (searchableData) {
					searchableData?.forEach((el) => {
						// console.log('el', el);
						el[Object.keys(el)].forEach((el1) => {
							// console.log('el1', el1);
							if (el1.checked == true) {
								tempArr.push(el1.name);
							}
						});
						tempObj[Object.keys(el)] = tempArr;
						tempArr = [];
						// console.log('tempObj', tempObj);
						searchableSendData.push(tempObj);
						tempObj = {};
					});
				}
				// console.log('parseSearchableToSendData :', searchableSendData);
			}
		}
	};

	const handleAddToCard = (item) => {
		if (browser) {
			if (store2.get('card')) {
				let card = store2.get('card');
				let exist = false;
				card.forEach((el) => {
					el.forEach((el1) => {
						if (item.id == el1.id) {
							exist = true;
							let tempEl = el;
							tempEl.push(item);
							el = [...tempEl];
							tempEl = [];
						}
					});
				});
				if (!exist) {
					let card1 = [];
					card1.push(item);
					card.push(card1);
				}
				store2.set('card', card);
				$storeCard = card;
			} else {
				let card2 = [];
				card2.push([item]);
				store2.set('card', card2);
				$storeCard = card;
			}
		}
	};

	const handleProductSelect = (item) => {
		console.log('selected porduct : ', item);

		redirectWithPageRefresh('/product/' + '?p=' + item?.id);

		
	};
</script>

{#if Object.values(parsedCategories).length > 0}
	<div class="container-fluid mt-4">
		<div class="row">
			<div style="border: 1px solid;" class="col-3">
				<h6>categories</h6>
				<div class="container">
					<div class="row">
						<div class="col-9">
							<TreeViewTry jsonTree={parsedCategories} />
						</div>
						<!-- <div class="col-3">
							<button
								on:click={() => {
									// searchHandler();
								}}
								class="btn btn-danger btn-sm mt-3">search</button
							>
						</div> -->

						{#if searchablePropertyList && searchablePropertyListParsed.length > 0}
							<div class="mt-2">
								<GenericSearch />
							</div>
						{/if}
					</div>
				</div>
			</div>
			<!-- backend pagination -->
			<div class="col-9">
				<div class="container">
					{#if splittedProductList && splittedProductList?.length > 0}
						{#each splittedProductList as group}
							<div class="row mt-2">
								{#each group as product}
									<div class="col-md-4">
										<!-- svelte-ignore a11y-click-events-have-key-events -->
										<!-- svelte-ignore a11y-no-static-element-interactions -->
										<div
											on:click={() => {
												handleProductSelect(product);
											}}
											class="card"
											style="width: 18rem;"
										>
											<div class="card-body">
												<h5 class="card-title">
													{product?.product?.product?.searchable?.name}
												</h5>
												<h6 class="card-title">
													{product?.product?.product?.searchable?.price}
												</h6>

												<button
													on:click={() => {
														handleAddToCard(product);
													}}
													class="btn btn-primary">+</button
												>
											</div>
										</div>
									</div>
								{/each}
							</div>
						{/each}
					{:else}
						there is no product with these filters
					{/if}
				</div>
			</div>
		</div>
	</div>
{:else}
	<h5 style="text-align: center;">Select Category</h5>
{/if}
