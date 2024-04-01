<script>
	import { superForm } from 'sveltekit-superforms/client';
	import { loginSchema } from '$lib/schemas.js';
	import { loginInitialData } from '$lib/initialDatas';
	import { browser } from '$app/environment';
	import Swal from 'sweetalert2';
	import { redirectWithPageRefresh } from '$lib/utils';

	const { form, errors, allErrors, enhance, delayed, tainted, message } = superForm(
		loginInitialData,
		{
			clearOnSubmit: 'none',
			validators: loginSchema,
			validationMethod: 'oninput',
			defaultValidator: 'keep',
			customValidity: false,
			delayMs: 500,
			timeoutMs: 8000,
			multipleSubmits: 'prevent',
			taintedMessage: null
		}
	);

	let isDisable = true;

	$: if ($tainted === undefined || $allErrors.length > 0) {
		isDisable = true;
	} else if ($tainted !== undefined) {
		if (Object.entries($tainted).length !== Object.entries(loginInitialData).length) {
			isDisable = true;
		} else {
			isDisable = false;
		}
	}

	$: if ($message?.alert) {
		//NOTE: carefull when you assing one object to another
		$tainted = JSON.parse(JSON.stringify(loginInitialData));
	}

	let alertType = 'success';
	let alertMsg = 'msg';
	$: if ($message?.alert) {
		if ($message?.alert == 'fail') {
			if (browser) {
				Swal.fire({
					position: 'bottom-end',
					title: 'Login failed',
					color: 'red',
					showConfirmButton: false,
					timer: 1500
				});
			}
			alertType = 'danger';
		} else if ($message?.alert == 'success') {
			if (browser) {
				Swal.fire({
					position: 'bottom-end',
					title: 'Login successful',
					color: 'green',
					showConfirmButton: false,
					timer: 1500
				});
			}
			redirectWithPageRefresh('/');
			alertType = 'success';
		}

		if ($message?.msg) {
			alertMsg = $message?.msg?.message;
		}
	}

	let passwordShow = false;
</script>

<div style="max-width: 600px;" class="container">
	<div class="card mt-5">
		<h1 style="text-align: center;">Login</h1>

		<div class="card-body">
			<form method="post" use:enhance>
				<div class="mb-3">
					<label for="emailInput" class="form-label">Email</label>
					<input
						type="email"
						class={$errors.email ? 'form-control is-invalid' : 'form-control'}
						id="emailInput"
						name="email"
						bind:value={$form.email}
						aria-invalid={$errors.email ? 'true' : undefined}
					/>

					{#if $errors.email}<span class="invalid-feedback">{$errors.email}</span>{/if}
				</div>
				<div class="mb-3">
					<div class="input-group">
						<input
							type={passwordShow ? 'text' : 'password'}
							class={$errors.password ? 'form-control is-invalid' : 'form-control'}
							id="passwordInput"
							name="password"
							value={$form.password}
							on:input={(e) => ($form.password = e.target.value)}
							aria-invalid={$errors.password ? 'true' : undefined}
						/>
						<div class="input-group-append">
							<a
								href="#"
								on:click={() => (passwordShow = !passwordShow)}
								class="input-group-text"
								id="passwordInputAppend">{passwordShow ? 'hide' : 'show'}</a
							>
						</div>

						{#if $errors.password}<span class="invalid-feedback">{$errors.password}</span>{/if}
					</div>
				</div>

				<!-- TODO: clean error when type something  -->
				{#if $message && JSON.stringify($tainted) == JSON.stringify(loginInitialData)}
					<div class={'alert alert-' + alertType} role="alert">
						{alertMsg}
					</div>
				{/if}

				<!-- TODO set disable as delayed -->

				<div class="mb-3" style="overflow: hidden;">
					<button class="btn btn-primary" style="float:right;" disabled={$delayed || isDisable}>
						{#if $delayed}
							Loading..
							<div style="width: 16px;height:16px;" class="spinner-border" role="status">
								<span class="sr-only"></span>
							</div>
						{:else}
							Submit
						{/if}
					</button>
				</div>

				<!-- <p>isDisable : {JSON.stringify(isDisable)}</p>
				<p>tainted : {JSON.stringify($tainted)}</p>
				<p>allErrors : {JSON.stringify($allErrors)}</p>
				<p>form : {JSON.stringify($form)}</p>
				<p>message.alert : {JSON.stringify($message)}</p> -->
			</form>
		</div>
	</div>
</div>
