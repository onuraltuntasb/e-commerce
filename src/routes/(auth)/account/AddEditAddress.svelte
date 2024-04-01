<script>
	import { addAddressInitialData } from '$lib/initialDatas';
	import { addressSchema } from '$lib/schemas';
	import { onMount } from 'svelte';
	import { superForm } from 'sveltekit-superforms/client';
	import { storeAddressList, storeSelectedAddress } from '../../../stores';
	import Swal from 'sweetalert2';
	import { browser } from '$app/environment';
	import { PUBLIC_FRONTEND_ENDPOINT_HOST } from '$env/static/public';

	let addressList = [];
	let addressId = -1;
	export let option;

	onMount(() => {
		addressList = $storeAddressList;
		addressId = $storeSelectedAddress.id;
		handleGetCities();
		// console.log('option :', option);
		// console.log('addressId : ', addressId);
	});

	$: cityId = -1;
	$: districtId = -1;
	$: neighborhoodId = -1;

	let isDisable = true;

	let cityList = [];
	let districtList = [];
	let neighborhoodList = [];

	//HACK: IDK WHY EDIT ADDRESS INITIAL CAN'T SET IN ONMOUNT FROM STORE
	const { form, errors, allErrors, enhance, tainted, message } = superForm($storeSelectedAddress, {
		clearOnSubmit: 'none',
		validators: addressSchema,
		validationMethod: 'oninput',
		defaultValidator: 'keep',
		customValidity: false,
		delayMs: 500,
		timeoutMs: 8000,
		multipleSubmits: 'prevent',
		taintedMessage: null,
		onSubmit({ formData }) {
			formData.set('extra', JSON.stringify({ cityId, districtId, neighborhoodId, addressId }));
		}
	});

	$: (function validationCheck() {
		if ($tainted === undefined || $allErrors.length > 0) {
			isDisable = true;
		} else if ($tainted !== undefined) {
			if (Object.entries($tainted).length !== Object.entries(addAddressInitialData).length) {
				isDisable = true;
			} else {
				isDisable = false;
			}
		} //[$tainted,$allErrors]

		if ($message?.alert || option == 'edit') {
			$tainted = addAddressInitialData;
		} //[$message,tainted]
	})();

	let body = '';
	let alertStatus = '';

	$: if ($message) {
		body = $message?.data;
		// console.log('body :', body);
		alertStatus = $message?.alert;
	} //[$message]

	$: if (alertStatus == 'success') {
		// console.log('alertStatus :', alertStatus);
		if (option == 'add') {
			// console.log('option :', option);
			let addressListTemp = [...addressList, { id: body?.addressId, ...$form }];
			$storeAddressList = addressListTemp;
			// console.log('addressList: ', addressListTemp);
			alertStatus = '';
		} else if (option == 'edit') {
			// console.log('option :', option);
			let addressListTemp = addressList.filter((el) => el.id != addressId);
			let tempAdd = addressList.filter((el) => el.id == addressId);
			// console.log('tempAdd', tempAdd);
			addressListTemp.push({ id: addressId, ...$form, isPrimary: tempAdd[0]?.isPrimary });
			$storeAddressList = addressListTemp;
			alertStatus = '';
			// console.log('addressList: ', addressListTemp);
		}
		Swal.fire({
			position: 'bottom-end',
			title: `Address ${option} success`,
			color: 'green',
			showConfirmButton: false,
			timer: 1500
		});
		// const myModal = bootstrap.Modal.getOrCreateInstance(
		// 	document.getElementById('addEditAddressModal')
		// );
		// myModal.hide();
	} else if (alertStatus == 'fail') {
		Swal.fire({
			position: 'bottom-end',
			title: `Address ${option} fail`,
			text: JSON.stringify(body),
			color: 'red',
			showConfirmButton: false
		});
		// const myModal = bootstrap.Modal.getOrCreateInstance(
		// 	document.getElementById('addEditAddressModal')
		// );
		// myModal.hide();
	} //[$form]

	const handleGetCities = async () => {
		const res = await fetch('api/city');
		const body = await res.json();
		if (body?.msg == 'status unauthorized') {
			if (browser) {
				window.location.href = PUBLIC_FRONTEND_ENDPOINT_HOST + '/login';
			}
		}

		// console.log('option : ', option);

		cityList = body.msg;
		if (option == 'add') {
			$form.city = body.msg[0].name;
		}
		cityId = body.msg[0].id;
		// console.log('cityId :', cityId);
		handleGetDistricts();
	};

	//HACK: THERE IS SOME ASYNC ISSUE WHEN GETTING FORM VALUE AND CITY LIST
	//FILTER IS NOT BEST PRACTICE BUT SSR DON'T LET ONCHANGE ON OPTION ELEMENT
	const handleGetDistricts = async () => {
		//HACK: SETTIMEOUT GOES MACRO TASK QUEUE SO IT'S LOWER PRIORITY IT'S
		//GONNA WAIT MICRO TASKS LIKE ASYNC FETCH TO FINISH $form.city idk why
		//its async
		setTimeout(async () => {
			let city = [...cityList].filter((el) => el.name == $form.city)[0];
			// console.log("city : ",city);
			cityId = city?.id;
			const res = await fetch('api/district/' + cityId);
			const body = await res.json();
			if (body?.msg == 'status unauthorized') {
				if (browser) {
					window.location.href = PUBLIC_FRONTEND_ENDPOINT_HOST + '/login';
				}
			}
			districtList = body.msg;
			if (option == 'add') {
				$form.district = body.msg[0].name;
			}
			districtId = body.msg[0].id;
			handleGetNeighborhoods();
		}, 0);
	};

	const handleGetNeighborhoods = async () => {
		setTimeout(async () => {
			let district = [...districtList].filter((el) => el.name == $form.district)[0];
			// console.log("district : ",district);
			districtId = district?.id;

			const res = await fetch('api/neighborhood/' + districtId);
			const body = await res.json();
			neighborhoodList = body.msg;
			if (option == 'add') {
				$form.neighborhood = body.msg[0].name;
			}
			$form.postalCode = body.msg[0].postalCode;
			neighborhoodId = body.msg[0].id;
		}, 0);
	};

	const handleNeighborhoodChange = async () => {
		setTimeout(async () => {
			let neighborhood = [...neighborhoodList].filter((el) => el.name == $form.neighborhood)[0];
			// console.log(neighborhood);
			neighborhoodId = neighborhood?.id;
			$form.postalCode = neighborhood?.postalCode;
		}, 0);
	};
</script>

<!-- <h4>{JSON.stringify($message)}</h4> -->
<!-- <h1>The sendedAddress is {JSON.stringify(sendedAddress)}</h1> -->
<!-- <h1>storeSelectedAddress : {JSON.stringify($storeSelectedAddress)}</h1> -->

<form method="POST" action={option == 'add' ? '?/addressAdd' : '?/addressEdit'} use:enhance>
	<div class="row">
		<div class="mb-3">
			<label for="headerInput" class="form-label">Address Header *</label>
			<input
				type="text"
				class={$errors.header ? 'form-control is-invalid' : 'form-control'}
				id="headerInput"
				name="header"
				bind:value={$form.header}
				aria-invalid={$errors.header ? 'true' : undefined}
			/>

			{#if $errors.header}<span class="invalid-feedback">{$errors.header}</span>{/if}
		</div>

		<div class="mb-3">
			<label for="nameInput" class="form-label">Name *</label>
			<input
				type="text"
				class={$errors.name ? 'form-control is-invalid' : 'form-control'}
				id="nameInput"
				name="name"
				bind:value={$form.name}
				aria-invalid={$errors.name ? 'true' : undefined}
			/>

			{#if $errors.name}<span class="invalid-feedback">{$errors.name}</span>{/if}
		</div>

		<div class="mb-3">
			<label for="countryInput" class="form-label">Country *</label>
			<select
				class={$errors.country ? 'form-control is-invalid' : 'form-control'}
				id="countryInput"
				name="country"
				on:change={() => {
					handleGetCities();
				}}
				bind:value={$form.country}
				aria-invalid={$errors.country ? 'true' : undefined}
			>
				<option value="" selected="selected" disabled>Choose...</option>
				{#each ['Turkey'] as item}
					<option value={item} selected={($form.country = item)}>{item}</option>
				{/each}
			</select>

			{#if $errors.country}<span class="invalid-feedback">{$errors.country}</span>{/if}
		</div>

		<div class="mb-3">
			<label for="cityInput" class="form-label">City *</label>
			<select
				class={$errors.city ? 'form-control is-invalid' : 'form-control'}
				id="cityInput"
				name="city"
				on:change={() => {
					handleGetDistricts();
				}}
				bind:value={$form.city}
				aria-invalid={$errors.city ? 'true' : undefined}
			>
				<option value="" selected="selected" disabled>Choose...</option>
				{#each cityList as item}
					<option value={item?.name}>{item?.name}</option>
				{/each}
			</select>

			{#if $errors.city}<span class="invaliæd-feedback">{$errors.city}</span>{/if}
		</div>

		<div class="mb-3">
			<label for="districtInput" class="form-label">District *</label>
			<select
				class={$errors.district ? 'form-control is-invalid' : 'form-control'}
				id="districtInput"
				name="district"
				on:change={() => {
					handleGetNeighborhoods();
				}}
				bind:value={$form.district}
				aria-invalid={$errors.district ? 'true' : undefined}
			>
				<option value="" selected="selected" disabled>Choose...</option>
				{#each districtList as item}
					<option value={item?.name}>{item?.name}</option>
				{/each}
			</select>

			{#if $errors.district}<span class="invaliæd-feedback">{$errors.district}</span>{/if}
		</div>

		<div class="mb-3">
			<label for="neighborhoodInput" class="form-label">Neighborhood *</label>
			<select
				class={$errors.neighborhood ? 'form-control is-invalid' : 'form-control'}
				id="neighborhoodInput"
				name="neighborhood"
				on:change={() => {
					handleNeighborhoodChange();
				}}
				bind:value={$form.neighborhood}
				aria-invalid={$errors.neighborhood ? 'true' : undefined}
			>
				{#each neighborhoodList as item, i}
					<option id="neighborhood-{i}" value={item?.name}>{item?.name}</option>
				{/each}
			</select>

			{#if $errors.neighborhood}<span class="invaliæd-feedback">{$errors.neighborhood}</span>{/if}
		</div>

		<div class="mb-3">
			<label for="descriptionInput" class="form-label">Description *</label>
			<textarea
				type="text"
				class={$errors.description ? 'form-control is-invalid' : 'form-control'}
				id="descriptionInput"
				name="description"
				bind:value={$form.description}
				aria-invalid={$errors.description ? 'true' : undefined}
			/>

			{#if $errors.description}<span class="invalid-feedback">{$errors.description}</span>{/if}
		</div>
	</div>
	<div class="mb-3">
		<label for="postalCodeInput" class="form-label">Post Code *</label>
		<input
			type="text"
			class={$errors.postalCode ? 'form-control is-invalid' : 'form-control'}
			id="postalCodeInput"
			name="postalCode"
			bind:value={$form.postalCode}
			aria-invalid={$errors.postalCode ? 'true' : undefined}
		/>

		{#if $errors.postalCode}<span class="invalid-feedback">{$errors.postalCode}</span>{/if}
	</div>

	<!--TODO: phone input customize -->
	<div class="mb-3">
		<label for="phoneInput" class="form-label">Phone *</label>
		<input
			type="tel"
			class={$errors.phone ? 'form-control is-invalid' : 'form-control'}
			id="phoneInput"
			name="phone"
			bind:value={$form.phone}
			aria-invalid={$errors.phone ? 'true' : undefined}
		/>
		<small class="text-muted">(5__) ___-____</small>
		{#if $errors.phone}<span class="invalid-feedback">{$errors.phone}</span>{/if}
	</div>

	<div class="mb-3">
		<p>Bill Type *</p>
		<label class="me-3">
			<input type="radio" bind:group={$form.billType} name="billType" value={true} />
			Individual
		</label>

		<label>
			<input type="radio" bind:group={$form.billType} name="billType" value={false} />
			Corporate
		</label>

		{#if $errors.billType}<span class="invalid-feedback">{$errors.billType}</span>{/if}
	</div>

	<div class="mb-3">
		<label for="ssnInput" class="form-label">SSN *</label>
		<input
			type="text"
			class={$errors.ssn ? 'form-control is-invalid' : 'form-control'}
			id="ssnInput"
			name="ssn"
			bind:value={$form.ssn}
			aria-invalid={$errors.ssn ? 'true' : undefined}
		/>

		{#if $errors.ssn}<span class="invalid-feedback">{$errors.ssn}</span>{/if}
	</div>

	<div class="mb-3">
		<button class="btn btn-primary ms-2" style="float:right;" disabled={isDisable}
			>{option == 'add' ? 'Save' : 'Update'}</button
		>
	</div>
</form>
<button
	on:click={() => {
		$storeSelectedAddress = {};
	}}
	type="button"
	class="btn btn-secondary"
	style="float:right;"
	data-bs-dismiss="modal">Close</button
>
