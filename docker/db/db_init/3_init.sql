-- CREATE TABLE quetions (
--     id int auto_increment not null primary key, 
--     title VARCHAR(32), 
--     content VARCHAR(256),
--     created_at TIMESTAMP,
--     modifyed_at TIMESTAMP,
--     group_id VARCHAR (256),
--     privated BOOLEAN,
--     create_user_id int
-- );

-- CREATE TABLE quetion_tags {
--     id int auto_increment not null primary key,
--     target_id int,
--     tag_id int
-- }

-- CREATE TABLE tags {
--     id int auto_increment not null primary key,
--     tag_name VARCHAR (256)
-- }

-- CREATE TABLE user_ids {
--     id int auto_increment not null primary key,
--     group_id int,
--     name_info VARCHAR (256)
-- }

-- CREATE TABLE group_ids {
--     id int auto_increment not null primary key,
--     name_info  VARCHAR (256)
-- }

CREATE TABLE users {
    id int auto_increment not null primary key,
    login_id VARCHAR (256),
    login_pass VARCHAR (256),
    benesse_id VARCHAR (256)
}

CREATE TABLE user_stamps {
    id int auto_increment not null primary key,
    user_id VARCHAR (256),
    stamp_count int,
    stamp_type int
}

CREATE TABLE user_studys {
    user_id int,
    japanese_score int,
    society_score int,
    math_score int,
    science_score int,
    english_score int
}

CREATE TABLE items {
    id int auto_increment not null primary key,
    stamp_type int,
    stamp_count int,
    item_name VARCHAR (256)
}