CREATE TABLE IF NOT EXISTS products (
    guid TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    price REAL NOT NULL,
    description TEXT,
    created_at TEXT NOT NULL
);
