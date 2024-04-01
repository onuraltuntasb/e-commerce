<script>
	import AddEditAddress from './AddEditAddress.svelte';
	import { browser } from '$app/environment';
	import { storeAddressList, storeSelectedAddress } from '../../../stores';
	import Swal from 'sweetalert2';
	import { redirectWithPageRefresh } from '$lib/utils';

	let addressList;
	$: {
		addressList = $storeAddressList;
	}

	let option = '';
	let selectedAddressId = -1;

	const handleMakeAddressPrimary = async (addressId) => {
		const res = await fetch(`api/auth/make-address-primary/${addressId}`);
		const body = await res.json();
		if (body?.msg == 'status unauthorized') {
			redirectWithPageRefresh('/login');
		}
		if (body?.alert == 'success') {
			if (addressList) {
				addressList = addressList.map((el) => {
					if (el.id == addressId) {
						// console.log('el id : ', el.id);
						// console.log('address id :', addressId);
						return {
							...el,
							isPrimary: true
						};
					} else {
						return {
							...el,
							isPrimary: false
						};
					}
				});
			}
		}
	};

	const handleDeleteAddress = async (addressId) => {
		// console.log('addressId :', addressId);
		const res = await fetch(`api/auth/address-delete/${addressId}`);
		const body = await res.json();
		if (body?.msg == 'status unauthorized') {
			redirectWithPageRefresh('/login');
		}
		// console.log('body :', body);
		if (body?.alert == 'success') {
			if (addressList) {
				addressList = addressList.filter((item) => item.id != addressId);
				$storeAddressList = addressList;
				if (browser) {
					const myModal = bootstrap.Modal.getOrCreateInstance(
						document.getElementById('deleteAddressModal')
					);
					myModal.hide();
					Swal.fire({
						position: 'bottom-end',
						title: 'Address deleted successfully!',
						color: 'green',
						showConfirmButton: false,
						timer: 1500
					});
				}
			}
		} else if (body?.alert == 'fail') {
			Swal.fire({
				position: 'bottom-end',
				title: 'Address deletion failed!',
				color: 'red',
				showConfirmButton: false,
				timer: 1500
			});
		}
	};

	const handleAddressEdit = (address) => {
		// console.log('handleAddressEdit : ', address);
		option = 'edit';
		$storeSelectedAddress = address;
	};
</script>

<div>
	<button
		on:click={() => {
			option = 'add';
		}}
		class="btn btn-danger btn-sm mt-2"
		data-bs-toggle="modal"
		data-bs-target="#addEditAddressModal">Add Address</button
	>

	<div
		class="modal fade"
		id="addEditAddressModal"
		tabindex="-1"
		aria-labelledby="addEditAddressModalLabel"
		aria-hidden="true"
	>
		<div class="modal-dialog">
			<div class="modal-content">
				<div class="modal-header">
					<h1 class="modal-title fs-5" id="addEditAddressModalLabel">{`${option} address`}</h1>
					<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"
					></button>
				</div>
				<div class="modal-body">
					{#if option == 'add' || option == 'edit'}
						<AddEditAddress option={option == 'add' ? 'add' : 'edit'} />
					{/if}
				</div>
			</div>
		</div>
	</div>

	<div>
		{#if addressList && addressList.length > 0}
			{#each addressList.sort((a, b) => a.id - b.id) as item}
				<div>
					<h5>{item.header}</h5>

					<div class="ps-3">
						<h6>{item.name}</h6>
						<div>{item.id}</div>
						<div>{item.description}</div>
						<div>{item.city}</div>
						<div>{item.district}</div>
						<div>{item.neighborhood}</div>
						<div>{item.phone}</div>
						<div>{item.billType ? 'personal' : 'corporate'}</div>
						<div>{item.ssn}</div>
						{#if item.isPrimary}
							<div style="font-style: italic;font-weight: 500;">Primary Address</div>
						{:else}
							<button class="btn btn-info btn-sm" on:click={() => handleMakeAddressPrimary(item.id)}
								>Make primary address</button
							>
						{/if}

						<button
							class="btn btn-info btn-sm"
							data-bs-toggle="modal"
							data-bs-target="#addEditAddressModal"
							on:click={() => {
								handleAddressEdit(item);
							}}>Edit</button
						>

						<button
							class="btn btn-info btn-sm"
							data-bs-toggle="modal"
							data-bs-target="#deleteAddressModal"
							on:click={() => (selectedAddressId = item.id)}>Delete</button
						>
					</div>
				</div>
			{/each}
		{:else}
			There is no address
		{/if}
	</div>

	<div
		class="modal fade"
		id="deleteAddressModal"
		tabindex="-1"
		aria-labelledby="deleteAddressModalLabel"
		aria-hidden="true"
	>
		<div class="modal-dialog">
			<div class="modal-content">
				<div class="modal-header">
					<h1 class="modal-title fs-5" id="deleteAddressModalLabel">Delete Address</h1>
					<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"
					></button>
				</div>
				<div class="modal-body">Are you sure you want to delete this delivery address?</div>
				<div class="modal-footer">
					<button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
					<button
						type="button"
						class="btn
										btn-primary"
						on:click={() => {
							handleDeleteAddress(selectedAddressId);
						}}>Delete</button
					>
				</div>
			</div>
		</div>
	</div>
</div>
