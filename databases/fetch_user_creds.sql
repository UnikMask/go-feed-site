SELECT
    username,
    password_hash
FROM
    Users
WHERE
    email = ?;
