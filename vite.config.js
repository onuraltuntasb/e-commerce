import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	plugins: [sveltekit()],
	optimizeDeps: { include: ['objection', 'knex', 'pg'], exclude: ['pg-native'] }
	// server: {
	// 	hmr: false
	// }
});
