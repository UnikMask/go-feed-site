PRAGMA foreign_keys = ON;
CREATE TABLE IF NOT EXISTS Users (
    id int
    username varchar,
    email varchar,
    password_hash varchar,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS Follows (
    follower int,
    followee int,
    FOREIGN KEY(follower) REFERENCES Users(id),
    FOREIGN KEY(followee) REFERENCES Users(id),
    PRIMARY KEY(follower, followee)
);

CREATE TABLE IF NOT EXISTS Posts (
    id int,
    contents varchar(3000)
);

CREATE TABLE IF NOT EXISTS Likes (
    user_id int,
    post_id int,
    FOREIGN KEY(user_id) REFERENCES Users(id),
    FOREIGN KEY(post_id) REFERENCES Posts(id),
    PRIMARY KEY(user_id, post_id)
);
