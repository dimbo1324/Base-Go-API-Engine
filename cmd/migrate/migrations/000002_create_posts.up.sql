CREATE TABLE IF NOT EXISTS posts (
    id bigserial PRIMARY KEY,
    title text NOT NULL,
    user_id bigint NOT NULL,
    content text NOT NULL,
    tags text[],
    created_at timestamp(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id)
);