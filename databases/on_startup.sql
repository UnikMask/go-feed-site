PRAGMA foreign_keys = ON;
CREATE TABLE IF NOT EXISTS Users (
    username varchar,
    email varchar,
    first_name varchar,
    last_name varchar,
    password_hash varchar
);

CREATE TABLE IF NOT EXISTS Sessions (
    user_id int,
    expiration date
);

CREATE TABLE IF NOT EXISTS Follows (
    follower int,
    followee int,
    FOREIGN KEY(follower) REFERENCES Users(ROWID),
    FOREIGN KEY(followee) REFERENCES Users(ROWID),
    PRIMARY KEY(follower, followee)
);

CREATE TABLE IF NOT EXISTS Posts (
    contents varchar(3000)
);

CREATE TABLE IF NOT EXISTS Likes (
    user_id int,
    post_id int,
    FOREIGN KEY(user_id) REFERENCES Users(ROWID),
    FOREIGN KEY(post_id) REFERENCES Posts(ROWID),
    PRIMARY KEY(user_id, post_id)
);
