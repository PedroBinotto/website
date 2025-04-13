-- name: GetBlogs :many
SELECT * FROM blogs;

-- name: CreateBlog :execresult
INSERT INTO blogs (
    title,
    url,
    body
) VALUES (
  'Test',
  '/1_test',
  'Test test test'
);
