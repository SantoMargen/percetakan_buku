CREATE TYPE user_role AS ENUM ('ADMIN', 'PENYELENGGARA', 'PENGUNJUNG');

CREATE TABLE mst_users (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    first_name varchar(50) NULL,
    last_name varchar(50) NULL,
    email varchar(32) NOT NULL UNIQUE,
    "password" varchar(10485760) NULL,
    age int4 NOT NULL,
    username varchar(32) NOT NULL UNIQUE,
    phone_number varchar(15) NULL,
    address varchar(10485760) NULL,
    role user_role NOT NULL,
    created_at timestamptz NOT NULL DEFAULT current_timestamp,
    updated_at timestamptz NULL
);