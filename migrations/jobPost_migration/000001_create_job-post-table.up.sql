CREATE TABLE job_posts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    comp_id UUID NOT NULL,
    name TEXT NOT NULL,
    description TEXT,
    amount INTEGER,
    time TEXT,
    time_range TEXT,
    period TEXT,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_recruiter
        FOREIGN KEY (comp_id)
        REFERENCES recruiters(id)
        ON DELETE CASCADE
);