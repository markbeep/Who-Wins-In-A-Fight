CREATE TABLE IF NOT EXISTS cards (
    id SERIAL PRIMARY KEY,
    wins INT NOT NULL DEFAULT 0,
    battles INT NOT NULL DEFAULT 0,
    name TEXT NOT NULL,
    token TEXT NOT NULL,
    filename TEXT NOT NULL,
    accepted BOOLEAN NOT NULL DEFAULT false
);

CREATE TABLE IF NOT EXISTS battles (
    id SERIAL PRIMARY KEY,
    card1_id INT NOT NULL REFERENCES cards(id) ON DELETE CASCADE,
    card2_id INT NOT NULL REFERENCES cards(id) ON DELETE CASCADE,
    card1_wins INT NOT NULL DEFAULT 0,
    card2_wins INT NOT NULL DEFAULT 0,
    CONSTRAINT card_order CHECK (card1_id < card2_id)
);
    

CREATE TABLE IF NOT EXISTS user_tokens (
    id SERIAL PRIMARY KEY,
    token TEXT NOT NULL,
    superuser BOOLEAN NOT NULL DEFAULT false
);
