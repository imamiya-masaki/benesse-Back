
CREATE TABLE users (
    id int auto_increment not null primary key,
    login_id VARCHAR (256),
    login_pass VARCHAR (256),
    benesse_id VARCHAR (256)
);

CREATE TABLE user_stamps (
    id int auto_increment not null primary key,
    user_id VARCHAR (256),
    stamp_count int,
    stamp_type int
);

CREATE TABLE user_stus(
    user_id int,
    japanese_score int,
    society_score int,
    math_score int,
    science_score int,
    english_score int
);

CREATE TABLE items (
    id int auto_increment not null primary key,
    stamp_type int,
    stamp_count int,
    item_name VARCHAR (256)
);