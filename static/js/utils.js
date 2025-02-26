if (localStorage.theme === 'dark' || (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
  document.documentElement.classList.add('dark')
}

function toggleDarkMode() {
  const root = document.documentElement;
  if (root.classList.contains('dark')) {
    root.classList.remove('dark');
    localStorage.theme = 'light';
  } else {
    root.classList.add('dark');
    localStorage.theme = 'dark';
  }
}

function toggleNavbar() {
  document.getElementById('menu-icon').classList.toggle('rotate-90');
  document.getElementById('mobile-menu').classList.toggle('hidden');
}
