CREATE TABLE campaigns (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    is_active BOOLEAN DEFAULT FALSE,
    start_date TIMESTAMPTZ NOT NULL,
    end_date TIMESTAMPTZ,
    max_vouchers INT DEFAULT 0,
    hold_vouchers INT DEFAULT 0,
    discount_percentage INT DEFAULT 30,
    available_vouchers INT DEFAULT 0
);

CREATE TABLE purchases (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    voucher_code VARCHAR(255),
    subscription_plan_price_details_id BIGINT NOT NULL,
    total_price DECIMAL(10, 2) DEFAULT 0.00
);

CREATE TABLE vouchers (
    id BIGSERIAL PRIMARY KEY,
    code VARCHAR(255) NOT NULL,
    user_id BIGINT,
    campaign_id BIGINT NOT NULL,
    discount_percentage INT DEFAULT 30,
    expired_at TIMESTAMPTZ NOT NULL,
    is_active BOOLEAN DEFAULT TRUE
);

CREATE TABLE subscription_plan (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    features JSONB,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE subscription_plan_details (
    id BIGSERIAL PRIMARY KEY,
    subscription_plan_id BIGINT NOT NULL,
    currency VARCHAR(10) NOT NULL,
    price DECIMAL(10, 2) DEFAULT 0.00,
    plan VARCHAR(50) NOT NULL,
    fee DECIMAL(10, 2) DEFAULT 0.00
);


INSERT INTO subscription_plan 
        (name, created_at, updated_at)
VALUES
        ('Silver', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);


WITH plan_id AS (
    SELECT id FROM subscription_plan WHERE name = 'Silver' LIMIT 1
)
INSERT INTO subscription_plan_details 
        (subscription_plan_id, currency, price, plan, fee)
SELECT id, 'USD', 10.00, 'month', 2.00 FROM plan_id;

WITH plan_id AS (
    SELECT id FROM subscription_plan WHERE name = 'Silver' LIMIT 1
)
INSERT INTO subscription_plan_details 
        (subscription_plan_id, currency, price, plan, fee)
SELECT id, 'USD', 100.00, 'annual', 7.00 FROM plan_id;


INSERT INTO campaigns 
        (name, description, discount_percentage, is_active, start_date, end_date, max_vouchers, hold_vouchers, available_vouchers)
VALUES
        ('Silver Discount Campaign', 
         '30% discount on Silver subscription plan for the first 100 users who register through the campaign link.', 
         30, TRUE, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + INTERVAL '1 month', 100, 0, 100);
