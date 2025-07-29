-- Drop index on books.title if exists
DROP INDEX IF EXISTS idx_book_covers_book_id;

DROP TABLE IF EXISTS book_covers;