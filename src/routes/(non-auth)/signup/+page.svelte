<script>
	import { superForm } from 'sveltekit-superforms/client';
	import { signupSchema } from '$lib/schemas.js';
	import { signupInitialData } from '$lib/initialDatas';
	import { browser } from '$app/environment';
	import Swal from 'sweetalert2';

	const { form, errors, allErrors, enhance, delayed, tainted, message } = superForm(
		signupInitialData,
		{
			clearOnSubmit: 'none',
			validators: signupSchema,
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
		if (Object.entries($tainted).length !== Object.entries(signupInitialData).length) {
			isDisable = true;
		} else {
			isDisable = false;
		}
	}

	$: if ($message?.alert) {
		$tainted = { ...signupInitialData };
	}

	let alertType = 'success';
	let alertMsg = 'msg';
	$: if ($message?.alert) {
		if ($message?.alert == 'fail') {
			if (browser) {
				Swal.fire({
					position: 'bottom-end',
					title: 'Sign up failed',
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
					title: 'Sign up successful',
					color: 'green',
					showConfirmButton: false,
					timer: 1500
				});
				redirectWithPageRefresh('/');
			}
			alertType = 'success';
		}

		if ($message?.msg) {
			alertMsg = $message?.msg?.message;
		}
	}

	let passwordShow = false;
	let passwordConfirmShow = false;
</script>

<div style="max-width: 600px;" class="container">
	<div class="card mt-5">
		<h1 style="text-align: center;">Signup</h1>

		<div class="card-body">
			<form method="post" use:enhance>
				<div class="mb-3">
					<label for="firstNameInput" class="form-label">First Name</label>
					<input
						type="text"
						class={$errors.firstName ? 'form-control is-invalid' : 'form-control'}
						id="firstNameInput"
						name="firstName"
						bind:value={$form.firstName}
						aria-invalid={$errors.firstName ? 'true' : undefined}
					/>

					{#if $errors.firstName}<span class="invalid-feedback">{$errors.firstName}</span>{/if}
				</div>

				<div class="mb-3">
					<label for="lastNameInput" class="form-label">Last Name</label>
					<input
						type="text"
						class={$errors.lastName ? 'form-control is-invalid' : 'form-control'}
						id="lastNameInput"
						name="lastName"
						bind:value={$form.lastName}
						aria-invalid={$errors.lastName ? 'true' : undefined}
					/>

					{#if $errors.lastName}<span class="invalid-feedback">{$errors.lastName}</span>{/if}
				</div>

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
					<label for="passwordInput" class="form-label">Password</label>

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
							<button
								type="button"
								on:click={() => (passwordShow = !passwordShow)}
								class="input-group-text imitate-a-tag"
								id="passwordInputAppend">{passwordShow ? 'hide' : 'show'}</button
							>
						</div>

						{#if $errors.passwordConfirm}<span class="invalid-feedback"
								>{$errors.passwordConfirm}</span
							>{/if}
					</div>
				</div>
				<div class="mb-3">
					<label for="passwordConfirmInput" class="form-label">Confirm Password</label>

					<div class="input-group">
						<input
							type={passwordConfirmShow ? 'text' : 'password'}
							class={$errors.passwordConfirm ? 'form-control is-invalid' : 'form-control'}
							id="passwordConfirmInput"
							name="passwordConfirm"
							value={$form.passwordConfirm}
							on:input={(e) => ($form.passwordConfirm = e.target.value)}
							aria-invalid={$errors.passwordConfirm ? 'true' : undefined}
						/>
						<div class="input-group-append">
							<button
								type="button"
								on:click={() => (passwordConfirmShow = !passwordConfirmShow)}
								class="input-group-text imitate-a-tag"
								id="passwordConfirmInputAppend">{passwordConfirmShow ? 'hide' : 'show'}</button
							>
						</div>

						{#if $errors.passwordConfirm}<span class="invalid-feedback"
								>{$errors.passwordConfirm}</span
							>{/if}
					</div>
				</div>

				{#if $message && JSON.stringify($tainted) == JSON.stringify(signupInitialData)}
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

				<!-- TODO handle internal error in hooks.server -->
				<!-- <p>message : {JSON.stringify($message)}</p> -->
			</form>
		</div>
	</div>
</div>
