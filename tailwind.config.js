/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./templates/**/*.{html,templ}'],
  theme: {
    extend: {
      fontFamily: {
        sans: ['IBM Plex Sans', 'system-ui', 'sans-serif'],
        serif: ['DM Serif Text', 'serif'],
        mono: ['Inconsolata', 'mono'],
      },
      fontSize: {
        'header-1': '8rem',
        'header-2': '3.5rem',
        'header-3': '1.5rem'
      },
      lineHeight: {
        'paragraph': '150%',
        'heading': '90%',
      },
      letterSpacing: {
        'paragraph': '-1%',
        'heading': '-5%',
      },
      colors: {
        background: {
          light: '#f8f9fa',
          dark: '#1e1e1e',
        },
        primary: {
          light: '#2563eb',
          dark: '#9333ea',
        },
        text: {
          light: '#111827',
          dark: '#e5e7eb',
        },
      },
      typography: (theme) => ({
        DEFAULT: {
          css: {
            'h1, h2, h3, h4': {
              fontFamily: theme('fontFamily.serif'),
              fontWeight: theme('fontWeight.bold'),
              lineHeight: theme('lineHeight.heading'),
              letterSpacing: theme('letterSpacing.heading'),
            },
            h1: {
              fontSize: theme('fontSize.header-1'),
            },
            h2: {
              fontSize: theme('fontSize.header-2'),
            },
            h3: {
              fontSize: theme('fontSize.header-3'),
            },
            p: {
              fontFamily: theme('fontFamily.sans'),
              fontSize: theme('fontSize.base'),
              fontWeight: theme('fontWeight.light'),
            },
            'p, blockquote': {
              lineHeight: theme('lineHeight.paragraph'),
              letterSpacing: theme('letterSpacing.paragraph'),
            },
            maxWidth: '90ch',
          },
        },
      }),
    },
  },
  plugins: [require('@tailwindcss/typography')],
  darkMode: 'class',
}
