<script>
	import AddressList from './AddressList.svelte';
	import PasswordChange from './PasswordChange.svelte';
	import AccountInfos from './AccountInfos.svelte';
	import { browser } from '$app/environment';
	import { storeAccountInfos, storeAccountState, storeAddressList } from '../../../stores';
	import ProductList from './ProductList.svelte';
	import { onMount } from 'svelte';
	import { PUBLIC_FRONTEND_ENDPOINT_HOST } from '$env/static/public';

	export let data;
	let addressList = data?.addressLoadForm?.message?.data;
	$storeAddressList = addressList;

	//HACK: i used browser state hack but it's not a good practice if you use ssr
	//make page design if you use csr make component redirect design. Hibrit
	//model is so complex and unnecessary
	let select = '';
	let selectedMenu = [];

	const subMenuIds = ['addressesId', 'passwordChangeId', 'accountInfosId', 'productListId'];

	let accountInfos = {};

	if (browser) {
		window.addEventListener('popstate', function (event) {
			if (event) {
				selectedMenu.splice(selectedMenu.length - 1, 1);
				// console.log('selectedMenu :', selectedMenu);
				if (selectedMenu.length >= 1) {
					select = selectedMenu[selectedMenu.length - 1];
					highlightSubMenu(select);
				}
				// enableSubMenu(selectedMenu[selectedMenu.length - 1]);
				// console.log('popstate');
			}
		});
	}

	$: {
		$storeAccountInfos = data?.accountInfos;
		accountInfos = $storeAccountInfos?.msg;
	}

	onMount(() => {
		enableSubMenu('addressesId');
	});

	const enableSubMenu = (id) => {
		selectedMenu.push(id);
		select = selectedMenu[selectedMenu.length - 1];
		// console.log('selectedMenu :', selectedMenu);

		if (browser) {
			window.history.pushState('', 'account', PUBLIC_FRONTEND_ENDPOINT_HOST + '/account');
		}
		highlightSubMenu(id);
	};
	const highlightSubMenu = (id) => {
		if (browser) {
			subMenuIds.forEach((el) => {
				let item = document.getElementById(el);
				if (el != id) {
					if (item) {
						item.style.fontWeight = 400;
					}
				} else if (el == id) {
					if (item) {
						item.style.fontWeight = 500;
					}
				}
			});
		}
	};
</script>

<div style="font-size: .85rem" class="container-xxl mt-4">
	<!-- <p>{JSON.stringify(data)}</p> -->
	<!-- <h4>{JSON.stringify(data?.addressLoadForm?.message?.data)}</h4> -->
	<div class="row">
		<div class="col-2">
			<ul class="list-group list-group-flush">
				<!-- svelte-ignore a11y-click-events-have-key-events -->
				<!-- svelte-ignore a11y-no-noninteractive-element-interactions -->
				<li
					id="addressesId"
					on:click={() => {
						enableSubMenu('addressesId');
					}}
					class="list-group-item"
				>
					Addresses
				</li>
				<!-- svelte-ignore a11y-click-events-have-key-events -->
				<!-- svelte-ignore a11y-no-noninteractive-element-interactions -->
				<li
					id="passwordChangeId"
					on:click={() => {
						enableSubMenu('passwordChangeId');
					}}
					class="list-group-item"
				>
					Password Change
				</li>
				<!-- svelte-ignore a11y-click-events-have-key-events -->
				<!-- svelte-ignore a11y-no-noninteractive-element-interactions -->
				<li
					id="accountInfosId"
					on:click={() => {
						enableSubMenu('accountInfosId');
					}}
					class="list-group-item"
				>
					Account Infos
				</li>
				<!-- svelte-ignore a11y-click-events-have-key-events -->
				<!-- svelte-ignore a11y-no-noninteractive-element-interactions
				-->
				{#if accountInfos?.isSeller}
					<li
						id="productListId"
						on:click={() => {
							enableSubMenu('productListId');
						}}
						class="list-group-item"
					>
						Product List
					</li>
				{/if}
			</ul>
		</div>

		<!-- <p>{select}</p> -->

		{#if select == 'addressesId'}
			<div style="border: 1px solid;" class="col-10">
				<AddressList />
			</div>
		{:else if select == 'passwordChangeId'}
			<PasswordChange />
		{:else if select == 'accountInfosId'}
			<div style="border: 1px solid;" class="col-10">
				<AccountInfos />
			</div>
		{:else if select == 'productListId'}
			<div style="border: 1px solid;" class="col-10">
				<ProductList />
			</div>
		{/if}
	</div>
</div>
