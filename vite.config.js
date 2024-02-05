import path from "path"
import { defineConfig } from 'vite'

export default defineConfig({
  root:  "assets/",
  base: "/",
  build: {
    ourDir: path.join(__dirname,'dist'),
    // generate ./assets/manifest.json in outDir
    manifest: "assets/manifest.json",
    rollupOptions: {
      // overwrite default .html entry
      input: "assets/scripts/app.ts",
    },
  },
});
