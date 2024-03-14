SELECT
    p.ROWID,
    u.ROWID,
    u.username,
    p.contents,
    p.posted_at,
    COUNT(l.post_id) AS likes
FROM
    Posts AS p
JOIN Users AS u
    ON p.user_id = u.ROWID
LEFT JOIN Likes AS l
    ON l.post_id = p.ROWID
WHERE
    p.ROWID = ?
GROUP BY
    p.ROWID;
