/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./ui/**/**/*.{js,html,tmpl,css}",
    "./assets/**/**/*.{js,css,svg}",
    "./public/**/**/*.js",
  ],
  theme: {
    extend: {
      spacing: {
        4.5: "1.125rem",
        15: "3.75rem",
        tiny: "calc(0.125 * var(--space-base))",
        vSmall: "calc(0.25 * var(--space-base))",
        small: "calc(0.5 * var(--space-base))",
        base: "var(--space-base)",
        large: "calc(1.5 * var(--space-base))",
        vLarge: "calc(2 * var(--space-base))",
        huge: "calc(3 * var(--space-base))",
      },

      fontSize: {
        vSmall: "calc(0.5 * var(--fontSize-base))",
        small: "calc(0.75 * var(--fontSize-base))",
        base: "var(--fontSize-base)",
        large: "calc(1.25 * var(--fontSize-base))",
        vLarge: "calc(1.5 * var(--fontSize-base))",
        huge: "calc(2.25 * var(--space-base))",
      },
    },
  },
  plugins: [],
};
