<script>
	import CheckboxGroup from './CheckboxGroup.svelte';

	let group = [];
	let searchablePropertyListParsed = [];

	$: {
		searchablePropertyListParsed = JSON.parse(localStorage.getItem('searchablePropertyList'));
	}
</script>

<div>
	{#if searchablePropertyListParsed}
		{#each searchablePropertyListParsed as item}
			<!-- <p>{JSON.stringify(item[Object.keys(item)])}</p> -->
			<!-- {JSON.stringify(item)} -->
			{#if Object.keys(item) != 'categories' && Object.keys(item) != 'price' && Object.keys(item) != 'name'}
				<CheckboxGroup checkboxes={item[Object.keys(item)]} header={Object.keys(item)} />
			{/if}
		{/each}
	{/if}

	<!-- <CheckboxGroup {checkboxes} bind:group /> -->
	<h1>{group}</h1>
	<br />
	<!-- {JSON.stringify(searchablePropertyListParsed)} -->
</div>
