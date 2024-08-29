/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./ui/**/**/*.{js,html,tmpl,css}",
    "./assets/**/**/*.{js,css,svg}",
    "./public/**/**/*.js",
  ],
  theme: {
    extend: {
      colors: {
        primary: {
          50: "#E5F0FB",
          100: "#C6DFF5",
          200: "#8ECOEC",
          300: "#5AA3E3",
          400: "#2583D5",
          500: "#1b609d",
          600: "#164D7E",
          700: "#103B60",
          800: "#0A253D",
          900: "#05131E",
          950: "#030B11",
        },
        secondary: {
          50: "#E8E8E8",
          100: "#CFCFCF",
          200: "#A1A1A1",
          300: "#707070",
          400: "#424242",
          500: "#121212",
          600: "#0F0F0F",
          700: "#0A0A0A",
          800: "#080808",
          900: "#030303",
          950: "#030303",
        },
      },
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
