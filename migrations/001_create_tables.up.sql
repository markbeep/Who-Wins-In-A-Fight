CREATE TABLE IF NOT EXISTS battle_card (
    id SERIAL PRIMARY KEY,
    url TEXT,
    wins INT,
    losses INT,
    elo INT,
    name TEXT
);

CREATE TABLE IF NOT EXISTS category (
    id SERIAL PRIMARY KEY,
    token TEXT NOT NULL,
    title TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS battle (
    id SERIAL PRIMARY KEY,
    start TIMESTAMP DEFAULT NOW(),
    card1_id INT REFERENCES battle_card(id),
    card2_id INT REFERENCES battle_card(id),
    token TEXT,
    category_id INT REFERENCES category(id)
);

