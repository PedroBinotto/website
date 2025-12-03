INSERT INTO blog (id, title, synopsis, url, body, created_at)
VALUES 
  (1, 'my first blog post',
      'A short introductory post marking the start of the blog, with simple placeholder content.',
      'first_blog_entry',
      'lorem ipsum dolor',
      '2024-01-01'),

  (2, 'Platform Update',
      'An overview of new features and improvements released in the latest version of the platform.',
      'https://example.com/post-2',
      'Exploring the new release of our platform.',
      '2024-01-02'),

  (3, 'Getting Started with Go',
      'A beginner-friendly guide that walks through setting up a development environment using Go and Docker.',
      'https://example.com/how-to-go',
      'A quick tutorial on getting started with Go and Docker.',
      '2024-01-03'),

  (4, 'SEO Tips',
      'Practical and lightweight SEO techniques that help your site become more discoverable.',
      'https://example.com/seo-tips',
      'Improve discoverability using simple SEO techniques.',
      '2024-01-04'),

  (5, 'HTMX Explained',
      'An introduction to HTMX and how it enables simpler, more dynamic frontends without heavy JavaScript.',
      'https://example.com/htmx',
      'Why HTMX can simplify your frontend drastically.',
      '2024-01-05'),

  (6, 'Templ + Go',
      'A walkthrough of using Templ for clean server-side rendering and structured component-based layouts in Go.',
      'https://example.com/templ',
      'Using Templ for clean server-side rendering in Go.',
      '2024-01-06'),

  (7, 'DB Sync Workflow',
      'A guide on keeping markdown-based blog posts synchronized with your database through automated workflows.',
      'https://example.com/db-sync',
      'Synchronizing markdown blog posts with your database.',
      '2024-01-07'),

  (8, 'TS Migration Pitfalls',
      'Common issues developers face when migrating to TypeScript, along with tips to avoid them.',
      'https://example.com/typescript',
      'Common pitfalls when migrating to TypeScript.',
      '2024-01-08'),

  (9, 'Docker Networking',
      'An explanation of how Docker networks work and how to connect multiple containers effectively.',
      'https://example.com/containers',
      'Connecting multiple containers through Docker networks.',
      '2024-01-09'),

  (10, 'Centering in Tailwind',
      'A quick guide to the most reliable ways to center elements in TailwindCSS, including common mistakes.',
      'https://example.com/tailwind',
      'How to correctly center elements using TailwindCSS.',
      '2024-01-10'),

  (11, 'ZKSync Overview',
      'A simplified overview of ZKSync, its token model, and how the ecosystem around it is evolving.',
      'https://example.com/zksync',
      'Understanding ZKSync token supply and holders.',
      '2024-01-11');

INSERT INTO tag (display_name, url)
VALUES 
  ('Subject 1', '/tags/subject-1'),
  ('Subject 2', '/tags/subject-2');

INSERT INTO blog_tag (blog_id, tag_id)
VALUES 
  (1, 2),
  (2, 2),
  (2, 1);
