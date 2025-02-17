-- Base
CREATE TABLE IF NOT EXISTS currencies  (
    id SERIAL PRIMARY KEY,
    code VARCHAR(5) UNIQUE NOT NULL,
    name VARCHAR(64) NOT NULL,
    organization_id INT NOT NULL
);

-- Banks
CREATE TABLE IF NOT EXISTS banks (
    id int2 PRIMARY KEY CHECK (id >= 1000 AND id < 10000),
    planet_id INT NOT NULL,
    organization_id INT NOT NULL,
    name VARCHAR(255) NOT NULL
);
CREATE TABLE IF NOT EXISTS bank_exchange_rates (
    id SERIAL PRIMARY KEY,
    bank_id int2 REFERENCES banks(id),
    currency_from_id INT REFERENCES currencies(id) NOT NULL,
    currency_to_id INT REFERENCES currencies(id) NOT NULL,
    date DATE NOT NULL,
    rate INT NOT NULL
);
CREATE TABLE IF NOT EXISTS accounts (
    id INT PRIMARY KEY CHECK (id >= 1000000000000000 AND id < 10000000000000000),
    owner_id INT,
    currency_id INT REFERENCES currencies(id),
    balance BIGINT DEFAULT 0
);
CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    from_account_id INT NOT NULL,
    to_account_id INT NOT NULL,
    amount BIGINT NOT NULL,
    currency_id INT REFERENCES currencies(id) NOT NULL,
    reason VARCHAR(32),
    user_comment VARCHAR(512),
    created_at TIMESTAMP DEFAULT NOW()
);

-- Market
CREATE TABLE IF NOT EXISTS market_pairs (
    id SERIAL PRIMARY KEY,
    currency_sell_id INT REFERENCES currencies(id) NOT NULL,
    currency_buy_id INT REFERENCES currencies(id) NOT NULL
);
CREATE TABLE IF NOT EXISTS market_orders (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    pair_id INT REFERENCES market_pairs(id) NOT NULL,
    amount BIGINT NOT NULL,
    price BIGINT NOT NULL,
    status bool,
    created_at TIMESTAMP DEFAULT NOW()
);
CREATE TABLE IF NOT EXISTS market_trades (
    id SERIAL PRIMARY KEY,
    order_sell_id INT REFERENCES market_orders(id) NOT NULL,
    order_buy_id INT REFERENCES market_orders(id) NOT NULL,
    amount BIGINT NOT NULL,
    price BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
