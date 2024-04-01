import { PUBLIC_BACKEND_ENDPOINT_HOST, } from "$env/static/public";
import { addAddressInitialData } from "$lib/initialDatas";
import { addressSchema } from "$lib/schemas";
import { getUserIdFromAccessToken } from "$lib/utils.js";
import { error, fail, } from "@sveltejs/kit";
import { superValidate } from "sveltekit-superforms/server";

export const load = async ({ cookies, fetch: sFetch }) => {

	if (!addAddressInitialData) {
		throw error(404, 'add address initial data is not found!');
	}

	const addressLoadForm = await superValidate(addAddressInitialData, addressSchema);

	const userId = getUserIdFromAccessToken(cookies);

	const req = sFetch(PUBLIC_BACKEND_ENDPOINT_HOST + '/auth/addresses-get/' + userId, {
		method: 'GET',
		headers: {
			'Content-Type': 'application/json'
		}
	});

	const stat = (await req).status;
	// console.log("load addresses stat : ", stat);
	const body = await req.then(r => r.json());
	// console.log("load addresses body : ", body);

	if (stat >= 200 && stat <= 299) {
		addressLoadForm.message = { status: stat, data: body == null ? [] : body, alert: 'success' };
	} else if (stat >= 400 && stat <= 599) {
		addressLoadForm.message = { status: stat, data: body, alert: 'fail' };
	}

	// console.log("load address form :", addressLoadForm);

	let accountInfos = {};
	const getAccountInfo = async () => {
		const req = sFetch(PUBLIC_BACKEND_ENDPOINT_HOST + '/auth/account-info-get/' + userId);

		const stat = (await req).status;
		const body = await req.then((r) => r.json());
		// console.log("body## : ", body)

		if (stat >= 200 && stat <= 299) {
			return accountInfos = { status: stat, msg: body, alert: 'success' };
		}
		if (stat >= 400 && stat <= 599) {
			return accountInfos = { status: stat, msg: body, alert: 'fail' };
		}
	}
	accountInfos = getAccountInfo();

	return { addressLoadForm, accountInfos }

};

export const actions = {
	addressAdd: async ({ request, cookies, fetch }) => {

		const formData = await request.formData();
		const addressAddForm = await superValidate(formData, addressSchema);
		if (!addressAddForm.valid) {
			return fail(400, { addressAddForm });
		}

		if (formData.has('extra')) {
			let extra = JSON.parse(formData.get('extra'));

			const userId = getUserIdFromAccessToken(cookies)
			let sendObj = {
				userId: userId,
				header: addressAddForm.data.header,
				name: addressAddForm.data.name,
				country: addressAddForm.data.country,
				cityId: extra["cityId"],
				districtId: extra["districtId"],
				neighborhoodId: extra["neighborhoodId"],
				description: addressAddForm.data.description,
				phone: addressAddForm.data.phone,
				billType: addressAddForm.data.billType,
				ssn: addressAddForm.data.ssn,
				isPrimary: false
			}

			console.log("addressAddForm : ", sendObj);

			const req = fetch(PUBLIC_BACKEND_ENDPOINT_HOST + '/auth/address-add', {
				method: 'POST',
				body: JSON.stringify(sendObj),
				headers: {
					'Content-Type': 'application/json'
				}
			});

			const stat = (await req).status;
			// console.log("add address stat : ", stat);
			const body = await req.then(r => r.json());
			// console.log("add address body : ", body);

			if (stat >= 200 && stat <= 299) {
				addressAddForm.message = { status: stat, data: body, alert: 'success' };
			}
			if (stat >= 400 && stat <= 599) {
				addressAddForm.message = { status: stat, data: body, alert: 'fail' };
			}
			return { addressAddForm };
		} else {
			throw error(400, "there is no extra part on form data!");
		}
	},

	addressEdit: async ({ request, cookies, fetch }) => {

		const formData = await request.formData();
		const addressEditForm = await superValidate(formData, addressSchema);
		if (!addressEditForm.valid) {
			return fail(400, { addressEditForm });
		}

		if (formData.has('extra')) {
			let extra = JSON.parse(formData.get('extra'));
			const userId = getUserIdFromAccessToken(cookies)

			let sendObj = {
				id: extra["addressId"],
				userId: userId,
				header: addressEditForm.data.header,
				name: addressEditForm.data.name,
				country: addressEditForm.data.country,
				cityId: extra["cityId"],
				districtId: extra["districtId"],
				neighborhoodId: extra["neighborhoodId"],
				description: addressEditForm.data.description,
				phone: addressEditForm.data.phone,
				billType: addressEditForm.data.billType,
				ssn: addressEditForm.data.ssn,
				isPrimary: false
			}

			// console.log("address edit send obj : ", sendObj);


			const req = fetch(PUBLIC_BACKEND_ENDPOINT_HOST + '/auth/address-update', {
				method: 'PUT',
				body: JSON.stringify(sendObj),
				headers: {
					'Content-Type': 'application/json'
				}
			});

			const stat = (await req).status;
			// console.log("address edit stat : ", stat);
			const body = await req.then(r => r.json());
			// console.log("address edit body : ", body);

			if (stat >= 200 && stat <= 299) {
				addressEditForm.message = { status: stat, data: body, alert: 'success' };
			}
			if (stat >= 400 && stat <= 599) {
				addressEditForm.message = { status: stat, data: body, alert: 'fail' };
			}
			return { addressEditForm };
		} else {
			throw error(400, "there is no extra data on form!");
		}
	}
};

