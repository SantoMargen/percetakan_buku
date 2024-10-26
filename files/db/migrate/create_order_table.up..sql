CREATE TYPE payment_status AS ENUM  ('PAID', 'CANCEL','FAILED');

CREATE TABLE orders (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    event_id UUID NOT NULL, 
    user_id UUID NOT NULL,
    ticket_id UUID NOT NULL,
    price int not null, 
    quantity int not null,
    status payment_status NOT  NULL,
    created_at timestamptz NOT NULL DEFAULT current_timestamp,
    updated_at timestamptz NULL
)