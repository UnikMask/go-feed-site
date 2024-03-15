SELECT
    COUNT(l.post_id) AS likes
FROM
    Posts AS p
LEFT JOIN Likes AS l
    ON l.post_id = p.ROWID
WHERE
    p.ROWID = ?1
GROUP BY
    p.ROWID;
