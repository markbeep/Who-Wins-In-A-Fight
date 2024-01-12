CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    token TEXT NOT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS cards (
    id SERIAL PRIMARY KEY,
    filename TEXT NOT NULL,
    wins INT DEFAULT 0,
    battles INT DEFAULT 0,
    elo INT DEFAULT 1000,
    name TEXT,
    category_id INT NOT NULL REFERENCES categories(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS battles (
    id SERIAL PRIMARY KEY,
    start TIMESTAMP DEFAULT NOW(),
    card1_id INT NOT NULL REFERENCES cards(id) ON DELETE CASCADE,
    card2_id INT NOT NULL REFERENCES cards(id) ON DELETE CASCADE,
    token TEXT NOT NULL,
    category_id INT NOT NULL REFERENCES categories(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS user_tokens (
    id SERIAL PRIMARY KEY,
    token TEXT NOT NULL,
    category_id INT NOT NULL REFERENCES categories(id) ON DELETE CASCADE
);
