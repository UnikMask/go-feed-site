SELECT
    ROWID,
    password_hash
FROM
    Users
WHERE
    email = ?;
