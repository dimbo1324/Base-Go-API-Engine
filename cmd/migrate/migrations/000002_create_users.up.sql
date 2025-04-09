CREATE TABLE IF NOT EXISTS posts (
  id bigserial PRIMARY KEY,
  user_id bigint NOT NULL REFERENCES users(id),
  title text NOT NULL,
  content text NOT NULL,
  tags text[] NOT NULL,
  created_at timestamptz NOT NULL DEFAULT NOW(),
  updated_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_posts_updated_at
BEFORE UPDATE ON posts
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();
