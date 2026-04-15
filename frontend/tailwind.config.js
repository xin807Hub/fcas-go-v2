/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  important: true,
  theme: {
    extend: {
      backgroundColor: {
        "main": "#F5F5F5",
      },
      textColor: {
        "active": "var(--el-color-primary)",
      },
      boxShadowColor: {
        "active": "var(--el-color-primary)",
      },
      borderColor: {
        "table-border": "var(--el-border-color-lighter)",
      },
      width: {
        "112": "28rem",
        "120": "30rem",
        "128": "32rem",
        "144": "36rem",
        "160": "40rem",
        "176": "44rem",
        "182": "48rem",
        "198": "52rem",
      },
      height: {
        "112": "28rem",
        "128": "32rem",
        "144": "36rem",
        "160": "40rem",
        "176": "44rem",
        "182": "48rem",
        "198": "52rem",
      }
    },
  },
  darkMode: "class",
  plugins: [],
  corePlugins: {
    preflight: false
  }
}

