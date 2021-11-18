CREATE TABLE IF NOT EXISTS films (
    id SERIAL PRIMARY KEY, 
    published_at timestamp NOT NULL,
    channel_id varchar(50) NOT NULL REFERENCES channels ON DELETE CASCADE ON UPDATE CASCADE,
    title varchar(100) NOT NULL,
    descript text NOT NULL,
    view_count bigint NOT NULL,
    like_count int NOT NULL,
    dislike_count int NOT NULL,
    comment_count int NOT NULL
);

