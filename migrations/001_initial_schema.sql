CREATE TABLE emails (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    sender TEXT NOT NULL,
    recipients TEXT NOT NULL, -- Could be normalized if time permits
    subject TEXT,
    timestamp DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP 
);

