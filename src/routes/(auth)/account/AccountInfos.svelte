<script>
	import Swal from 'sweetalert2';
	import { storeAccountInfos } from '../../../stores';

	let accountInfos = {};

	$: {
		accountInfos = $storeAccountInfos?.msg;
	}

	const makeSellerHandler = async () => {
		const res = await fetch(`api/auth/make-account-seller`);
		const body = await res.json();
		if (body?.msg == 'status unauthorized') {
			redirectWithPageRefresh('/login');
		}
		if (body?.alert == 'success') {
			accountInfos.isSeller = true;
			$storeAccountInfos.msg = accountInfos;
			Swal.fire({
				position: 'bottom-end',
				title: 'Make seller success!',
				color: 'green',
				showConfirmButton: false,
				timer: 1500
			});
		} else if (body?.alert == 'fail') {
			Swal.fire({
				position: 'bottom-end',
				title: 'Account info fetched fail!',
				color: 'red',
				showConfirmButton: false,
				timer: 1500
			});
		}
	};
</script>

<div>
	<h5>Account Infos</h5>
	<!-- <p>{JSON.stringify(accountInfos)}</p>
	<p>{JSON.stringify($storeAccountInfos)}</p> -->
	<div>
		<div>firstName : {accountInfos?.firstName}</div>
		<div>lastName : {accountInfos?.lastName}</div>
		<div>email : {accountInfos?.email}</div>
		<div>
			isSeller : {accountInfos?.isSeller}
			<button
				on:click={() => {
					makeSellerHandler();
				}}>make</button
			>
		</div>
	</div>
</div>
