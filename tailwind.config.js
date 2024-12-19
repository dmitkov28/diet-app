/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/**/*.templ", "./templates/*.templ"],
  theme: {
    extend: {
      animation: {
        shake: "shake 1s ease-in-out",
      },
      keyframes: {
        shake: {
          "0%, 100%": {
            transform: "rotate(0deg)",
          },
          "25%": {
            transform: "rotate(-30deg)",
          },
          "75%": {
            transform: "rotate(30deg)",
          },
        },
      },
    },
  },
  plugins: [],
};
