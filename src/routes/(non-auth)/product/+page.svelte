<script>
	import { browser } from '$app/environment';
	import { json } from '@sveltejs/kit';
	import { onMount, tick } from 'svelte';
	import store2 from 'store2';
	import AccountInfos from '../../(auth)/account/AccountInfos.svelte';

	//TODO: get product and product's comment and user email

	let product = {};
	let productComments = [];
	let productSelectables = [];
	let selectedOptions = [];
	let isDifferent = false;

	onMount(() => {
		if (browser) {
			const url = new URL(window.location.href);
			const productId = url.searchParams.get('p');
			// console.log('productId :', productId);
			getProduct(productId);
			getProductComments(productId);
			if (store2.get('pageUrl')) {
				if (store2.get('pageUrl') != window.location.href) {
					isDifferent = true;
				}
			}
		}
	});

	const getProduct = async (productId) => {
		const res = await fetch(`api/product-get/${productId}`);
		const body = await res.json();
		if (body?.msg == 'status unauthorized') {
			redirectWithPageRefresh('/login');
		}
		if (body?.alert == 'success') {
			// console.log('body.msg : ', body?.msg);
			product = body?.msg;
			getSelectables(product);
		} else if (body?.alert == 'fail') {
			Swal.fire({
				position: 'bottom-end',
				title: 'Product fetch failed!',
				color: 'red',
				showConfirmButton: false,
				timer: 1500
			});
		}
	};

	const getProductComments = async (productId) => {
		const res = await fetch(`api/product-comments-get/${productId}`);
		const body = await res.json();
		if (body?.msg == 'status unauthorized') {
			redirectWithPageRefresh('/login');
		}
		if (body?.alert == 'success') {
			// console.log('body.msg : ', body?.msg);
			productComments = body?.msg;
		} else if (body?.alert == 'fail') {
			Swal.fire({
				position: 'bottom-end',
				title: 'Product comments fetch failed!',
				color: 'red',
				showConfirmButton: false,
				timer: 1500
			});
		}
	};

	// {#if product?.product?.product?.searchable}
	// 				{#each Object.values(product?.product?.product?.searchable) as item}
	// 					{#if item instanceof Array && typeof item?.[0] === 'object' && item?.[0] !== null}
	// 						<p>{JSON.stringify(item)}</p>
	// 					{/if}
	// 				{/each}
	// 			{/if}

	let defaultProductSelectables = [];

	const getSelectables = async () => {
		if (product) {
			// console.log('product : ', product);
			if (product?.product?.product?.searchable) {
				for (let [key, value] of Object.entries(product?.product?.product?.searchable)) {
					if (value instanceof Array) {
						if (typeof value[0] === 'object' && value[0] !== null) {
							// console.log(key, value[0]);
							let obj = {};
							obj[key] = value;

							productSelectables = [...productSelectables, obj];

							console.log('value :', value);

							let obj1 = {};
							let val = [...value];
							obj1[key] = val[0][Object.keys(val[0])[0]];
							selectedOptions = [...selectedOptions, obj1];

							let obj2 = {};
							let val1 = [...value];
							obj2[key] = [val1[0]];

							defaultProductSelectables = [...defaultProductSelectables, obj2];
						}
					}
				}
			}

			if (isDifferent) {
				store2.set('selectedOptions', selectedOptions);
			} else {
				if (!store2.get('selectedOptions')) {
					store2.set('selectedOptions', selectedOptions);
					store2.set('pageUrl', window.location.href);
				}
			}
		}
		await tick();
		processDefaultSelectables();

		// console.log('defaultProductSelectables :', defaultProductSelectables);
	};

	const processDefaultSelectables = () => {
		// console.log('defaultProductSelectables :', defaultProductSelectables);
		//process default selectables

		if (store2.get('selectedOptions')) {
			let sOptions = store2.get('selectedOptions');

			console.log('sOptions:', sOptions);

			sOptions.forEach((el) => {
				console.log('el:', el);

				let selectedOption = el[Object.keys(el)[0]];
				console.log('selectedOption :', selectedOption + 'Id');
				let els = document.getElementById(Object.keys(el)[0] + 'Id').children;
				Array.from(els).forEach((el1) => {
					// console.log('el1 :', el1);
					if (el1.id == selectedOption + 'Id') {
						el1.style.borderColor = 'red';
					} else {
						el1.style.borderColor = 'black';
					}
				});
			});
		} else {
			defaultProductSelectables?.forEach((el) => {
				// console.log('el :', el);
				// console.log('el name :', Object.keys(el)[0] + 'Id');
				// console.log('bak :', Object.values(el)[0][0]);

				let selectedOption = Object.values(el)[0][0][Object.keys(Object.values(el)[0][0])[0]];
				// console.log('selectedOption :', selectedOption + 'Id');
				let els = document.getElementById(Object.keys(el)[0] + 'Id').children;
				Array.from(els).forEach((el1) => {
					// console.log('el1 :', el1);
					if (el1.id == selectedOption + 'Id') {
						el1.style.borderColor = 'red';
					} else {
						el1.style.borderColor = 'black';
					}
				});
				// console.log('els :', els);
			});
		}
	};

	//TODO:localstorage selectedOptions write

	const handleSelectOption = (optionName, selectedOption, parentId, targetId) => {
		// console.log('handleSelectOptionItem :', optionName);
		// console.log('handleSelectOptionselectedOption :', selectedOption);

		if (selectedOptions?.length == 0) {
			let obj = {};
			obj[optionName] = selectedOption;
			selectedOptions = [...selectedOptions, obj];
		} else {
			let exist = false;
			selectedOptions?.forEach((el) => {
				if (Object.keys(el)[0] == optionName) {
					// console.log('Object.keys(el) :', Object.keys(el)[0]);

					el[Object.keys(el)[0]] = selectedOption;
					exist = true;
				}
			});
			if (!exist) {
				let obj = {};
				obj[optionName] = selectedOption;
				selectedOptions = [...selectedOptions, obj];
			}
		}

		selectedOptions = [...selectedOptions];

		if (browser) {
			let els = document.getElementById(parentId).children;
			Array.from(els).forEach((el) => {
				if (el.id == targetId) {
					// console.log('el :', el);
					el.style.borderColor = 'red';
				} else {
					el.style.borderColor = 'black';
				}
			});
			// console.log(els);
		}

		// console.log('selectedOptions :', JSON.stringify(selectedOptions));

		store2.set('selectedOptions', selectedOptions);
	};
</script>

<div>{JSON.stringify(product)}</div>

<div class="container-fluid">
	<div class="row">
		<div class="col-4">
			<div style="border:1px solid; display:flex;align-items:center;justify-content:center">
				<img
					width="%100"
					height="500px"
					src="https://n11scdn1.akamaized.net/a1/602_857/05/85/97/28/IMG-9148616888945759533.jpg"
					alt="productImage"
				/>
			</div>
		</div>
		<div class="col-5">
			<h6>{product?.product?.product?.searchable?.name}</h6>
			<hr />
			<h4>{product?.product?.product?.searchable?.price}</h4>
			<hr />
			<div style="display:flex">
				<div
					style="width: 50px;height:50px;border:1px solid; font-size:12px;text-align:center;margin-right:8px"
				>
					<div>Brand</div>
					<div>{product?.product?.product?.searchable?.brand}</div>
				</div>

				<div
					style="width: 50px;height:50px;border:1px solid;
				font-size:12px;text-align:center;"
				>
					<div>Sex</div>
					{#if product?.product?.product?.searchable?.sex}
						{#each product?.product?.product?.searchable?.sex as item}
							<p>{item?.sex}</p>
						{/each}
					{/if}
				</div>
				<!-- TODO: selectable product json -->
			</div>
			<hr />

			<div>
				<h6>Product Options</h6>

				{#if productSelectables?.length > 0}
					{#each productSelectables as item}
						<p>{Object.keys(item)} :</p>
						<div id={Object.keys(item) + 'Id'} style="display:flex;text-align:center">
							<!-- svelte-ignore a11y-no-static-element-interactions -->
							{#each item[Object.keys(item)] as item1}
								<!-- svelte-ignore a11y-click-events-have-key-events -->
								<div
									id={item1[Object.keys(item)] + 'Id'}
									on:click={(e) => {
										handleSelectOption(
											Object.keys(item)[0],
											item1[Object.keys(item)],
											Object.keys(item) + 'Id',
											item1[Object.keys(item)] + 'Id'
										);
									}}
									style="border: 1px solid; width:50px; margin-right:8px"
								>
									{item1[Object.keys(item)]}
								</div>
							{/each}
						</div>

						<br />
					{/each}
				{/if}

				<div style="border:1px solid;height:40px"></div>
			</div>
		</div>
		<div class="col-3">
			right

			{#if selectedOptions?.length > 0}
				<p>{JSON.stringify(selectedOptions)}</p>
			{/if}
		</div>
	</div>
</div>
