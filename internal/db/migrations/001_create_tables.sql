-- Base
CREATE TABLE IF NOT EXISTS currencies  (
    id SERIAL PRIMARY KEY,
    code VARCHAR(10) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL
);
CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    from_account_id INT,
    to_account_id INT,
    amount BIGINT NOT NULL,
    currency_id INT REFERENCES currencies(id),
    transaction_type VARCHAR(32)
);

-- Banks
CREATE TABLE IF NOT EXISTS banks (
    id SERIAL PRIMARY KEY,
    planet_id INT NOT NULL,
    code INT UNIQUE NOT NULL,
    name VARCHAR(255) UNIQUE NOT NULL,
    owner_id INT NOT NULL,
    currency_id INT REFERENCES currencies(id) NOT NULL
);
CREATE TABLE IF NOT EXISTS accounts (
    id INT PRIMARY KEY,
    owner_id INT NOT NULL,
    owner_type VARCHAR(10) CHECK (owner_type IN ('user', 'bank')),
    bank_id INT REFERENCES banks(id),
    currency_id INT REFERENCES currencies(id),
    balance BIGINT DEFAULT 0.0,
    is_reserve BOOLEAN DEFAULT FALSE
);
CREATE TABLE IF NOT EXISTS bank_exchange_rates (
    id SERIAL PRIMARY KEY,
    bank_id INT REFERENCES banks(id),
    currency_from_id INT REFERENCES currencies(id),
    currency_to_id INT REFERENCES currencies(id),
    date TIMESTAMP NOT NULL,
    rate BIGINT NOT NULL
);

-- Market
CREATE TABLE IF NOT EXISTS market_orders (
    id SERIAL PRIMARY KEY,
    user_id INT,
    currency_sell_id INT REFERENCES currencies(id) ON DELETE CASCADE,
    amount_sell BIGINT NOT NULL,
    currency_buy_id INT REFERENCES currencies(id) ON DELETE CASCADE,
    price BIGINT NOT NULL,
    order_type VARCHAR(10) CHECK (order_type IN ('buy', 'sell')),
    status VARCHAR(20) CHECK (status IN ('open', 'filled', 'cancelled')),
    created_at TIMESTAMP DEFAULT NOW()
);
CREATE TABLE IF NOT EXISTS market_trades (
    id SERIAL PRIMARY KEY,
    order_sell_id INT REFERENCES market_orders(id) ON DELETE CASCADE,
    order_buy_id INT REFERENCES market_orders(id) ON DELETE CASCADE,
    currency_sell_id INT REFERENCES currencies(id) ON DELETE CASCADE,
    amount_sell BIGINT NOT NULL,
    currency_buy_id INT REFERENCES currencies(id) ON DELETE CASCADE,
    price BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
