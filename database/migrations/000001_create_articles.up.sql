CREATE TABLE IF NOT EXISTS articles (
    id INTEGER PRIMARY KEY,
    title TEXT,
    body BLOB,
    created_on TEXT DEFAULT CURRENT_DATE
)
