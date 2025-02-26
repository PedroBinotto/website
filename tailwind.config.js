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
        'xs': '0.5rem',
        'sm': '0.7rem',
        'lg': '1.4rem',
        'xl': '2rem',
        '2xl': '2.8rem',
        '3xl': '4rem',
        '4xl': '5.6rem',
        '5xl': '8rem',
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
            'h1, h2, h3, h4, h5, h6': {
              fontFamily: theme('fontFamily.serif'),
              // fontWeight: theme('fontWeight.bold'),
              lineHeight: theme('lineHeight.heading'),
              letterSpacing: theme('letterSpacing.heading'),
            },
            h1: {
              fontSize: theme('fontSize.5xl'),
            },
            h2: {
              fontSize: theme('fontSize.4xl'),
            },
            h3: {
              fontSize: theme('fontSize.3xl'),
            },
            h4: {
              fontSize: theme('fontSize.2xl'),
            },
            h5: {
              fontSize: theme('fontSize.xl'),
            },
            h6: {
              fontSize: theme('fontSize.lg'),
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
            maxWidth: 'none',
          },
        },
      }),
    },
  },
  plugins: [require('@tailwindcss/typography')],
  darkMode: 'class',
}


