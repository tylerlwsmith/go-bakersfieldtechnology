import { defineConfig } from "vite";

export default defineConfig({
  root: "assets/",
  base: "/assets/",
  publicDir: false,
  build: {
    ourDir: "dist",
    manifest: "manifest.json",
    assetsDir: "public",
    emptyOutDir: true,
    rollupOptions: {
      input: ["assets/scripts/app.ts", "assets/styles/app.scss"],
    },
  },
});
