<script>
	import { JSONEditor } from 'svelte-jsoneditor';
	import { browser } from '$app/environment';
	import Swal from 'sweetalert2';
	import { storeProductList, storeSelectedProduct } from '../../../stores';
	import { redirectWithPageRefresh } from '$lib/utils';

	export let option;

	$: content = {
		text: option == 'add' ? '{}' : JSON.stringify($storeSelectedProduct?.product) // can be used to pass a stringified JSON document instead
	}; //[$storeSelectedProduct]

	function handleChange(updatedContent, previousContent, { contentErrors, patchResult }) {
		// content is an object { json: unknown } | { text: string }
		// console.log('onChange: ', { updatedContent, previousContent, contentErrors, patchResult });
		content = updatedContent;
		// console.log('content:', content);
	}

	const handleSubmit = async (product) => {
		if (option == 'add') {
			handleAddSubmit();
		} else if (option == 'edit') {
			handleEditSubmit(product);
		}
	};

	const handleEditSubmit = async (product) => {
		// console.log('handleEditSubmit product :  ', product);

		let userId = product?.userId;
		let id = product?.id;

		const res = await fetch(`api/auth/product-edit/${id}`, {
			method: 'PUT',
			body: JSON.stringify(content?.text),
			headers: {
				'Content-Type': 'application/json'
			}
		});
		const body = await res.json();
		if (body?.msg == 'status unauthorized') {
			redirectWithPageRefresh('/login');
		}
		if (body?.alert == 'success') {
			let productList = $storeProductList ? $storeProductList : [];
			let productListTemp = productList.filter((el) => el.id != id);
			// console.log('parse : ', JSON.parse(content?.text));
			productListTemp.push({ id: id, userId: userId, product: JSON.parse(content?.text) });
			$storeProductList = productListTemp;
			// console.log('productListTemp : ', productListTemp);

			if (browser) {
				const myModal = bootstrap.Modal.getOrCreateInstance(
					document.getElementById('addEditProductModal')
				);
				myModal.hide();
				Swal.fire({
					position: 'bottom-end',
					title: 'Product updated successfully!',
					color: 'green',
					showConfirmButton: false,
					timer: 1500
				});
			}
		} else if (body?.alert == 'fail') {
			if (browser) {
				const myModal = bootstrap.Modal.getOrCreateInstance(
					document.getElementById('addEditProductModal')
				);
				myModal.hide();
				Swal.fire({
					position: 'bottom-end',
					title: 'Product update failed!',
					color: 'green',
					showConfirmButton: false,
					timer: 1500
				});
			}
		}
	};

	const handleAddSubmit = async () => {
		const res = await fetch(`api/auth/product-add`, {
			method: 'POST',
			body: JSON.stringify(content?.text),
			headers: {
				'Content-Type': 'application/json'
			}
		});
		const body = await res.json();
		if (body?.msg == 'status unauthorized') {
			redirectWithPageRefresh('/login');
		}
		if (body?.alert == 'success') {
			if (browser) {
				let productId = body?.msg?.productId;
				let userId = body?.msg?.userId;
				console.log('text :', content?.text);

				let productList = $storeProductList ? $storeProductList : [];
				console.log('productList', productList);
				productList?.push({ id: productId, userId, product: JSON.parse(content?.text) });
				$storeProductList = productList;

				const myModal = bootstrap.Modal.getOrCreateInstance(
					document.getElementById('addEditProductModal')
				);
				myModal.hide();
				Swal.fire({
					position: 'bottom-end',
					title: 'Product added successfully!',
					color: 'green',
					showConfirmButton: false,
					timer: 1500
				});
			}
		} else if (body?.alert == 'fail') {
			if (browser) {
				const myModal = bootstrap.Modal.getOrCreateInstance(
					document.getElementById('addEditProductModal')
				);
				myModal.hide();
				Swal.fire({
					position: 'bottom-end',
					title: 'Product addition failed!',
					color: 'red',
					showConfirmButton: false,
					timer: 1500
				});
			}
		}
	};
</script>

<div>
	<JSONEditor mode={'text'} mainMenuBar={false} {content} onChange={handleChange} />

	<!-- <p>{JSON.stringify($storeSelectedProduct)}</p> -->

	<button
		on:click={() => {
			handleSubmit($storeSelectedProduct);
		}}
		class="btn btn-primary ms-2"
		style="float:right;">{option == 'add' ? 'Save' : 'Update'}</button
	>

	<button
		on:click={() => {
			$storeSelectedProduct = { product: {} };
		}}
		type="button"
		class="btn btn-secondary"
		style="float:right;"
		data-bs-dismiss="modal">Close</button
	>
</div>
