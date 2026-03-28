CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE seekers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,

    gender TEXT CHECK (gender IN ('male', 'female', 'other')) NOT NULL,

    age INT CHECK (age >= 18 AND age <= 100),

    qualification TEXT NOT NULL,

    adhar_no TEXT UNIQUE NOT NULL,
    phone_no TEXT UNIQUE NOT NULL,

    location TEXT NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);