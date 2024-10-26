CREATE TYPE status_event AS ENUM  ('ACTIVE', 'INACTIVE');

CREATE TABLE mst_events (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    "name" varchar(50) NULL,
    start_date timestamptz NOT NULL,
    end_date timestamptz NOT NULL,
    "description" TEXT  NOT NULL,
    "location" varchar(100) NOT NULL,
    status status_event NOT NULL, 
    created_at timestamptz NOT NULL DEFAULT current_timestamp,
    updated_at timestamptz NULL
);