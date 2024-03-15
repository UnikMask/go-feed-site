SELECT
    u.username,
    u.email,
    u.first_name,
    u.last_name
FROM
    Users AS u
WHERE
    u.ROWID = ?;
