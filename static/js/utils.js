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

function enableLandingAnimation() {
  const container = document.getElementById('landing-container');

  if (!container) return

  const contentHeight = container.scrollHeight;
  container.style.height = '100vh';
  window.addEventListener('scroll', () => {
    const scrollRatio = Math.min(window.scrollY / (window.innerHeight / 2), 1);
    const newHeight = window.innerHeight - (scrollRatio * (window.innerHeight - contentHeight));
    container.style.height = `${newHeight}px`;
  });
}

document.addEventListener('DOMContentLoaded', enableLandingAnimation);
if (localStorage.theme === 'dark' || (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
  document.documentElement.classList.add('dark')
}
