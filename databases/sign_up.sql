INSERT INTO Users(
    username,
    email,
    first_name,
    last_name,
    password_hash
) VALUES
    (?, ?, ?, ?, ?)
RETURNING ROWID;
