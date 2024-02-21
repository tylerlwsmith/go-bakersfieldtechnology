import { defineConfig } from "vite";

export default defineConfig({
  root: "assets/",
  base: "/assets/",
  build: {
    ourDir: "dist",
    manifest: "manifest.json",
    assetsDir: "public",
    rollupOptions: {
      input: ["assets/scripts/app.ts", "assets/styles/app.scss"],
    },
  },
});
