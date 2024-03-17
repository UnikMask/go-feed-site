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
    END AS liked,
    CASE 
        WHEN COUNT(f.follower) THEN 1
        ELSE 0
    END AS followed
FROM
    Posts AS p
JOIN Users AS u
    ON p.user_id = u.ROWID
LEFT JOIN Likes AS l
    ON l.post_id = p.ROWID
LEFT JOIN Follows f ON f.followee = u.ROWID AND f.follower = ?2
WHERE
    p.ROWID = ?1
GROUP BY
    p.ROWID, f.follower;
