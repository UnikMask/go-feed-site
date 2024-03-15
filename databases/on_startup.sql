PRAGMA foreign_keys = ON;
CREATE TABLE IF NOT EXISTS Users (
    username varchar,
    email varchar,
    first_name varchar,
    last_name varchar,
    password_hash varchar,
    UNIQUE(email)
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_users_email ON Users(email);

CREATE TABLE IF NOT EXISTS Sessions (
    user_id int,
    expiration date
);

CREATE TABLE IF NOT EXISTS Follows (
    follower int,
    followee int,
    FOREIGN KEY(follower) REFERENCES Users(ROWID) ON DELETE CASCADE,
    FOREIGN KEY(followee) REFERENCES Users(ROWID) ON DELETE CASCADE,
    PRIMARY KEY(follower, followee)
);

CREATE TABLE IF NOT EXISTS Posts (
    contents varchar(3000),
    user_id int,
    posted_at datetime,
    FOREIGN KEY(user_id) REFERENCES Users(ROWID) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS Likes (
    user_id int,
    post_id int,
    FOREIGN KEY(user_id) REFERENCES Users(ROWID) ON DELETE CASCADE,
    FOREIGN KEY(post_id) REFERENCES Posts(ROWID) ON DELETE CASCADE,
    PRIMARY KEY(user_id, post_id)
);
