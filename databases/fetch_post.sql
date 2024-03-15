SELECT
    p.ROWID,
    u.ROWID,
    u.username,
    p.contents,
    p.posted_at,
    COUNT(l.post_id) AS likes,
    CASE
        WHEN SUM(l.user_id = ?2) THEN 1
        ELSE 0
    END AS liked
FROM
    Posts AS p
JOIN Users AS u
    ON p.user_id = u.ROWID
LEFT JOIN Likes AS l
    ON l.post_id = p.ROWID
WHERE
    p.ROWID = ?1
GROUP BY
    p.ROWID;
