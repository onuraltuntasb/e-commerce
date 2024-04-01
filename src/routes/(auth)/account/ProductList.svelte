<script>
	import { onMount } from 'svelte';
	import AddEditProduct from './AddEditProduct.svelte';
	import Swal from 'sweetalert2';
	import { storeProductList, storeSelectedProduct } from '../../../stores';
	import { browser } from '$app/environment';
	import { PUBLIC_FRONTEND_ENDPOINT_HOST } from '$env/static/public';
	import { redirectWithPageRefresh } from '$lib/utils';

	let option = '';
	let productList = [];
	let selectedProductId = -1;

	//TODO:check edit delete product

	onMount(() => {
		getProductList();
		// console.log(productList);
	});

	$: {
		productList = $storeProductList;
	} //[$storeProductList]

	const getProductList = async () => {
		const res = await fetch(`${PUBLIC_FRONTEND_ENDPOINT_HOST}/api/auth/products-get-seller`);
		const body = await res.json();
		if (body?.msg == 'status unauthorized') {
			redirectWithPageRefresh('/login');
		}

		if (body?.alert == 'success') {
			$storeProductList = body?.msg;
			productList = $storeProductList;
		} else if (body?.alert == 'fail') {
			Swal.fire({
				position: 'bottom-end',
				title: 'Product list fetched fail!',
				color: 'red',
				showConfirmButton: false,
				timer: 1500
			});
		}
	};

	const handleProductEdit = (product) => {
		// console.log('handleProductEdit : ', product);
		option = 'edit';
		$storeSelectedProduct = product;
	};

	const handleProductDelete = async () => {
		// console.log('selectedProductId', selectedProductId);

		const res = await fetch(`api/auth/product-delete/${selectedProductId}`, {
			method: 'DELETE'
		});
		const body = await res.json();
		if (body?.msg == 'status unauthorized') {
			redirectWithPageRefresh('/login');
		}
		if (body?.alert == 'success') {
			let productList = $storeProductList;
			let productListTemp = productList.filter((el) => el.id != selectedProductId);
			$storeProductList = productListTemp;

			if (browser) {
				const myModal = bootstrap.Modal.getOrCreateInstance(
					document.getElementById('deleteProductModal')
				);
				myModal.hide();
				Swal.fire({
					position: 'bottom-end',
					title: 'Product delete successfully!',
					color: 'green',
					showConfirmButton: false,
					timer: 1500
				});
			}
		} else if (body?.alert == 'fail') {
			if (browser) {
				const myModal = bootstrap.Modal.getOrCreateInstance(
					document.getElementById('deleteProductModal')
				);
				myModal.hide();
				Swal.fire({
					position: 'bottom-end',
					title: 'Product delete failed!',
					color: 'green',
					showConfirmButton: false,
					timer: 1500
				});
			}
		}
	};
</script>

<div>
	{#if productList && productList.length > 0}
		{#each productList.sort((a, b) => a - b) as item}
			<div class="mb-2 p-2" style="border: solid 1px black; ">
				<div>{item?.product?.product?.searchable?.name}</div>
				<div>{item?.product?.product?.searchable?.price}</div>
				<div>{item?.product?.product?.searchable?.sex[0]?.sex}</div>

				<button
					class="btn btn-info btn-sm"
					data-bs-toggle="modal"
					data-bs-target="#addEditProductModal"
					on:click={() => {
						handleProductEdit(item);
					}}>Edit</button
				>
				<button
					class="btn btn-info btn-sm"
					data-bs-toggle="modal"
					data-bs-target="#deleteProductModal"
					on:click={() => {
						selectedProductId = item?.id;
					}}>Delete</button
				>
			</div>
		{/each}
	{:else}
		<p>there is no product</p>
	{/if}

	<!-- <h1>{JSON.stringify(productList)}</h1> -->

	<button
		on:click={() => {
			option = 'add';
		}}
		class="btn btn-danger btn-sm mt-2"
		data-bs-toggle="modal"
		data-bs-target="#addEditProductModal">Add Product</button
	>

	<div
		class="modal fade"
		id="addEditProductModal"
		tabindex="-1"
		aria-labelledby="addEditProductModalLabel"
		aria-hidden="true"
	>
		<div class="modal-dialog">
			<div class="modal-content">
				<div class="modal-header">
					<h1 class="modal-title fs-5" id="addEditProductModalLabel">{`${option} product`}</h1>
					<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"
					></button>
				</div>
				<div class="modal-body">
					{#if option == 'add' || option == 'edit'}
						<AddEditProduct option={option == 'add' ? 'add' : 'edit'} />
					{/if}
				</div>
			</div>
		</div>
	</div>

	<div
		class="modal fade"
		id="deleteProductModal"
		tabindex="-1"
		aria-labelledby="deleteProductModalLabel"
		aria-hidden="true"
	>
		<div class="modal-dialog">
			<div class="modal-content">
				<div class="modal-header">
					<h1 class="modal-title fs-5" id="deleteProductModalLabel">Delete Product</h1>
					<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"
					></button>
				</div>
				<div class="modal-body">Are you sure you want to delete this Product?</div>
				<div class="modal-footer">
					<button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
					<button
						type="button"
						class="btn
										btn-primary"
						on:click={() => {
							handleProductDelete();
						}}>Delete</button
					>
				</div>
			</div>
		</div>
	</div>
</div>
