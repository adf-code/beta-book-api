-- Create index on books.title if not exists
DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_indexes WHERE indexname = 'idx_books_title') THEN
CREATE INDEX idx_books_title ON books(title);
END IF;
END$$;

-- Create index on books.author if not exists
DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_indexes WHERE indexname = 'idx_books_author') THEN
CREATE INDEX idx_books_author ON books(author);
END IF;
END$$;

-- Create index on books.year if not exists
DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_indexes WHERE indexname = 'idx_books_year') THEN
CREATE INDEX idx_books_year ON books(year);
END IF;
END$$;

-- Create index on books.created_at if not exists
DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_indexes WHERE indexname = 'idx_books_created_at') THEN
CREATE INDEX idx_books_created_at ON books(created_at);
END IF;
END$$;
