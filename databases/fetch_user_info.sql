SELECT
    username,
    email,
    first_name,
    last_name
FROM
    Users
WHERE
    email = ?;
