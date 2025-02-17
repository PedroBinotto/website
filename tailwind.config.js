/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/**/*.{html,templ}"],
  theme: {
    extend: {},
  },
  plugins: [require("@tailwindcss/typography")],
}

