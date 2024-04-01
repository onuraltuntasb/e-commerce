<script>
	import { browser } from '$app/environment';
	import { onMount } from 'svelte';
	import { storeSelectedCategory } from '../../../stores';

	export let jsonTree;

	let selectedName = '';
	let jst = '';

	onMount(async () => {
		const obj = await import('js-treeview');
		jst = obj.default;
	});

	$: {
		if (browser) {
			if (jst) {
				let tree = new jst([jsonTree], 'tree');

				if (localStorage.getItem('category')) {
					let cat = localStorage.getItem('category');
					selectedName = cat;
					// console.log('first  cat :', selectedName);
					let els = document.getElementsByClassName('tree-leaf-text');
					// console.log('els :', els);
					if (els) {
						for (let i = 0; i < els.length; i++) {
							if (els[i].innerText == selectedName.replaceAll('"', '')) {
								// console.log('el :', els[i]);
								// console.log('selectedName :', selectedName);
								els[i].style.fontWeight = 500;
							} else {
								els[i].style.fontWeight = 400;
							}

							// console.log('el :', els[i]);
						}
					}
				}
				// console.log('JSONtree :', jsonTree);
				// console.log('tree :', tree);

				tree.on('select', function (e) {
					selectedName = e.data.name;
					highlightSelected();
					console.log(e.data.name);
				});

				tree.on('expand', function (e) {
					let name = JSON.stringify(e.target.innerText);
					name = name.slice(4, name.length - 1);
					selectedName = name;
					highlightSelected();
					console.log(name);
				});

				tree.on('collapse', function (e) {
					let name = JSON.stringify(e.target.innerText);
					name = name.slice(4, name.length - 1);
					selectedName = name;
					highlightSelected();
					console.log(name);
				});
			}
		}
	}

	const highlightSelected = () => {
		localStorage.setItem('category', selectedName);
		$storeSelectedCategory = selectedName;

		if (browser) {
			let els = document.getElementsByClassName('tree-leaf-text');
			if (els) {
				for (let i = 0; i < els.length; i++) {
					if (els[i].innerText == selectedName) {
						// console.log('el :', els[i]);
						console.log('selectedName :', selectedName);
						els[i].style.fontWeight = 500;
					} else {
						els[i].style.fontWeight = 400;
					}

					// console.log('el :', els[i]);
				}
			}
		}
	};
</script>

<div id="tree"></div>

<!-- {JSON.stringify(jsonTree)} -->

<style>
</style>
