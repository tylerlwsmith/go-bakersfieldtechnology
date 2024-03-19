const defaultTheme = require("tailwindcss/defaultTheme");

/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./components/**/*.templ"
  ],
  theme: {
    screens: {
      xs: "500px",
      ...defaultTheme.screens,
    },
    extend: {
      screens: {
        sm: "574px",
      },
    },
  },
  plugins: [require("@tailwindcss/typography")],
}
