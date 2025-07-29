CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS book_covers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    book_id UUID NOT NULL,
    file_name TEXT NOT NULL,
    file_url TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_book FOREIGN KEY(book_id) REFERENCES books(id) ON DELETE CASCADE
);

-- Create index on book_covers.book_id if not exists
DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_indexes WHERE indexname = 'idx_book_covers_book_id') THEN
CREATE INDEX idx_book_covers_book_id ON book_covers(book_id);
END IF;
END$$;