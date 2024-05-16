CREATE TABLE articles (
    id INTEGER PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    body BLOB,
    created_at DATE 
);
