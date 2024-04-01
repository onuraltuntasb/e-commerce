<script>
	import '../../style/app.css';
	import { afterNavigate, goto } from '$app/navigation';
	import {
		storeCard,
		storeCategories,
		storeSelectedCategory,
		storeSelectedMainCategory
	} from '../../stores';
	import { browser } from '$app/environment';
	import Swal from 'sweetalert2';
	import { PUBLIC_BACKEND_ENDPOINT_HOST, PUBLIC_FRONTEND_ENDPOINT_HOST } from '$env/static/public';
	import { getUserIdFromAccessToken, redirectWithPageRefresh } from '$lib/utils';
	import { onMount } from 'svelte';
	import store2 from 'store2';
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

	let card = [];

	onMount(() => {
		if (browser) {
			if (store2.get('card')) {
				$storeCard = store2.get('card');
			}
		}
	});

	$: {
		card = $storeCard;
		// console.log('card :', card);
	}

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
				// console.log('el : ', el);
				mainCategories.push(el);
			});
		}
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

	const handleRemoveFromCard = (item) => {
		console.log('remove item : ', item);
		card?.forEach((el) => {
			el?.forEach((el1) => {
				if (el1.id == item[1]?.[0]?.id) {
					el.pop();
					console.log('remove el1 :', el1);
				}
			});
			console.log('remove el :', el);
		});
		card = card?.filter((item) => item?.length != 0);
		store2.set('card', card);
		$storeCard = card;
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
					{:else if email === undefined || !email}
						<div class="d-flex me-2">
							<div class="dropdown">
								<button
									class="btn btn-secondary dropdown-toggle me-2"
									type="button"
									id="dropdownMenuButtonCard"
									data-bs-toggle="dropdown"
									aria-expanded="false">Card : {Object.values(card).length}</button
								>
								<ul class="dropdown-menu" aria-labelledby="dropdownMenuButtonCard">
									{#each Object.entries(card) as item}
										<div class="card" style="width: 18rem;">
											<img src="..." class="card-img-top" alt="..." />
											<div class="card-body">
												<h5 class="card-title">
													{JSON.stringify(item[1][0]?.product?.product?.searchable?.name)}
												</h5>
												<h6 class="card-title">
													{JSON.stringify(item[1][0]?.product?.product?.searchable?.price)}
												</h6>
												<p>count : {JSON.stringify(item[1].length)}</p>

												<button
													on:click={() => {
														handleRemoveFromCard(item);
													}}
													class="btn btn-primary">-</button
												>
											</div>
										</div>
									{/each}
								</ul>
							</div>

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
							$storeSelectedMainCategory = item?.name;
							if (browser) {
								redirectWithPageRefresh('/product');
								localStorage.setItem('mainCategory', item?.name);
								localStorage.setItem('category', item?.name);

								if ($storeSelectedCategory != item?.name) {
									if (localStorage.getItem('searchablePropertyList')) {
										localStorage.removeItem('searchablePropertyList');
									}
								}
								$storeSelectedCategory = item?.name;
							}
						}}
					>
						{item.name}
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
