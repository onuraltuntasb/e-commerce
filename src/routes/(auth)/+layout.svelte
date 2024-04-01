<script>
	import '../../style/app.css';
	import { afterNavigate, goto } from '$app/navigation';
	import { browser } from '$app/environment';
	import { storeCategories, storeSelectedCategory, storeSelectedMainCategory } from '../../stores';
	import { PUBLIC_FRONTEND_ENDPOINT_HOST } from '$env/static/public';
	import { redirectWithPageRefresh } from '$lib/utils';
	// import { onMount } from 'svelte';

	// let currentTheme;

	// onMount(() => {
	// 	const savedTheme = document.documentElement.getAttribute('data-bs-theme');
	// 	if (savedTheme) {
	// 		currentTheme = savedTheme;
	// 		return;
	// 	}

	// 	const preference_is_dark = window.matchMedia('(prefers-color-scheme: dark)').matches;

	// 	const theme = preference_is_dark ? 'dark' : 'light';
	// 	setTheme(theme); // TODO
	// });

	// function setTheme(theme) {
	// 	const oneYear = 60 * 60 * 24 * 365;
	// 	document.cookie = `theme=${theme}; max-age=${oneYear}; path=/`;
	// 	document.documentElement.setAttribute('data-bs-theme', theme);
	// 	currentTheme = theme;
	// }

	// function toggleTheme() {
	// 	const theme = currentTheme === 'light' ? 'dark' : 'light';
	// 	setTheme(theme);
	// }

	export let data;

	let alertMsg = data?.data?.alertMsg;
	let alertClassType = '';

	if (data?.alert === 'success') {
		alertClassType = 'success';
	} else if (data?.alert === 'fail') {
		alertClassType = 'danger';
	} else if (data?.alert === 'warning') {
		alertClassType = 'warning';
	}

	let visible = true;
	setTimeout(() => (visible = false), 2000);

	let email = data?.data?.email;

	afterNavigate(() => {
		email = data?.data?.email;
	});

	let categories = data?.data?.categories?.mainCategories;
	let mainCategories = [];
	const processCategoriesData = (categories) => {
		if (categories) {
			if (browser) {
				localStorage.setItem('categories', JSON.stringify(categories));
			}
			categories.forEach((el) => {
				// console.log('el', el);
				mainCategories.push(el);
			});
		}
		//TODO: else handle
	};
	processCategoriesData(categories);

	const handleLogout = async () => {
		const res = await fetch('api/auth/logout');
		const body = await res.json();
		if (body?.alert == 'success') {
			if (browser) {
				Swal.fire({
					position: 'bottom-end',
					title: 'Log out successful!',
					color: 'green',
					showConfirmButton: false,
					timer: 1500
				});
				window.location.href = PUBLIC_FRONTEND_ENDPOINT_HOST + '/login';
			}
		} else if (body?.alert == 'fail') {
			if (browser) {
				Swal.fire({
					position: 'bottom-end',
					title: 'Log out fail!',
					color: 'red',
					showConfirmButton: false,
					timer: 1500
				});
			}
		}
	};
</script>

<div>
	<nav class="navbar navbar-expand-lg bg-body-tertiary">
		<div class="container-fluid">
			<a class="navbar-brand" href="#!">e-commerce</a>

			<div class="navbar">
				<!-- no auth -->

				<!-- with auth -->
				<div class="d-flex">
					{#if email}
						<div class="dropdown">
							<button
								class="btn btn-secondary dropdown-toggle me-2"
								type="button"
								data-bs-toggle="dropdown"
								data-bs-placement="left"
							>
								{email.substring(0, 1)}
							</button>
							<ul class="dropdown-menu">
								<li>
									<button
										on:click={() => {
											if (browser) {
												window.location.href = PUBLIC_FRONTEND_ENDPOINT_HOST + '/account';
											}
										}}
										class="dropdown-item">account</button
									>
								</li>

								<li>
									<button on:click={() => handleLogout()} class="dropdown-item">logout</button>
								</li>
							</ul>
						</div>

						<button type="button" class="btn btn-primary me-2">Notifications</button>
						<button type="button" class="btn btn-primary me-2">Card</button>
					{:else if email === undefined || !email}
						<div class="d-flex me-2">
							<a href="/signup" class="me-2">Signup</a>
							<a href="/login" class="me-2">Login</a>
						</div>
					{/if}
				</div>
				<!-- <button class="btn btn-primary" aria-label="toggle theme" on:click={toggleTheme}
					>Toggle Theme</button
				> -->
			</div>
		</div>
	</nav>
	<div style="background-color: blanchedalmond;">
		<div class="container-xxl">
			{#if mainCategories}
				{#each mainCategories as item}
					<button
						type="button"
						class="btn btn-link btn-sm imitate-a-tag me-2"
						on:click={() => {
							$storeSelectedMainCategory = item?.text;
							if (browser) {
								redirectWithPageRefresh('/product?category=' + item?.text);

								if ($storeSelectedCategory != item?.text) {
									if (localStorage.getItem('searchablePropertyList')) {
										localStorage.removeItem('searchablePropertyList');
									}
								}
								$storeSelectedCategory = item?.text;
							}
						}}
					>
						{item.text}
					</button>
				{/each}
			{/if}
		</div>
	</div>

	{#if data?.data?.alertMsg && visible}
		<div class={'alert alert-' + alertClassType} style="text-align: center;" role="alert">
			{alertMsg}
		</div>
	{/if}

	<slot />
</div>
