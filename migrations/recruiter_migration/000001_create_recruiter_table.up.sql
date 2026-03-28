CREATE TABLE recruiters (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,

    license_no BIGINT UNIQUE NOT NULL,

    location TEXT NOT NULL,

    phone_no BIGINT UNIQUE NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);