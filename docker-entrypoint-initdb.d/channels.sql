CREATE TABLE IF NOT EXISTS channels(
    channel_id varchar(50) primary key,
    title varchar(50) NOT NULL,
    view_count bigint NOT NULL,
    subscriber_count int NOT NULL,
    video_count int NOT NULL
);