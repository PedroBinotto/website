-- name: GetBlogs :many
SELECT
  b.id,
  b.title,
  b.synopsis,
  b.url,
  b.body,
  b.created_at,
  COALESCE(
    json_agg(
      json_build_object(
        'display_name', t.display_name,
        'url', t.url
      )
    ) FILTER (WHERE t.id IS NOT NULL),
    '[]'
  ) AS tags
FROM blog b
LEFT JOIN blog_tag bt ON bt.blog_id = b.id
LEFT JOIN tag t       ON t.id = bt.tag_id
GROUP BY b.id
ORDER BY b.created_at DESC;
