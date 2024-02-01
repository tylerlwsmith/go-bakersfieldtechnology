// vite.config.js
export default defineConfig({
  build: {
    // generate ./assets/manifest.json in outDir
    manifest: "./assets/manifest.json",
    rollupOptions: {
      // overwrite default .html entry
      input: "/path/to/main.js",
    },
  },
});
