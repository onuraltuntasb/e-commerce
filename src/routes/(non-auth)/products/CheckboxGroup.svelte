<script>
	import { browser } from '$app/environment';
	import { onMount } from 'svelte';
	import { storeSearchablePropertyListParsed, storeSelectedSearchProps } from '../../../stores';

	export let checkboxes = [];
	export let group = [];
	export let header = '';

	$: leastOneTrue = false;

	onMount(() => {
		checkboxes?.forEach((el) => {
			if (el.checked) {
				leastOneTrue = true;
			}
		});
	});

	$: {
		if (leastOneTrue) {
			console.log('leastOneTrue');
			if (browser) {
				let element = document.querySelector(`#flush-collapseThree${header}`);
				bootstrap.Collapse.getOrCreateInstance(element).toggle();
			}
		}
	}

	const addHandler = (event) => {
		let val = event.target.value;
		console.log('event :', event.target.value);

		let tempArr = JSON.parse(localStorage.getItem('searchablePropertyList'));

		$storeSearchablePropertyListParsed = tempArr;

		tempArr.forEach((el) => {
			// console.log('el', el);
			el[Object.keys(el)].forEach((el1) => {
				console.log('Object.keys(el) :', Object.keys(el));
				console.log('header:', header);
				if (Object.keys(el)[0] == header) {
					if (el1.name == val) {
						if (el1.checked == true) {
							leastOneTrue = true;
							el1.checked = false;
						} else if (el1.checked == false) {
							el1.checked = true;
						}
					}
				}
				// console.log('el1', el1);
			});
		});

		//TODO: localstorage sync
		localStorage.setItem('searchablePropertyList', JSON.stringify(tempArr));
		// console.log('tempArr :', tempArr);
	};
</script>

<div class="accordion accordion-flush" id={'accordionFlushExample' + header}>
	<div class="accordion-item">
		<h2 class="accordion-header" id="flush-headingThree">
			<button
				class="accordion-button collapsed"
				type="button"
				data-bs-toggle="collapse"
				data-bs-target={'#flush-collapseThree' + header}
				aria-expanded="false"
				aria-controls="flush-collapseThree"
			>
				{header}
			</button>
		</h2>
		<div
			id={'flush-collapseThree' + header}
			class="accordion-collapse collapse"
			aria-labelledby="flush-headingThree"
			data-bs-parent={'#accordionFlushExample' + header}
		>
			<div class="accordion-body">
				{#each checkboxes as checkbox}
					<input
						on:click={(e) => {
							addHandler(e);
						}}
						type="checkbox"
						value={checkbox?.name}
						checked={checkbox?.checked}
					/>
					<label class="label">{checkbox?.name}</label>
					<br />
				{/each}
				<p>group : {group}</p>
			</div>
		</div>
	</div>
</div>

<!-- {JSON.stringify('checkboxes :', checkboxes)} -->
