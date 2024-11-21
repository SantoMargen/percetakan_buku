CREATE TYPE type_ticket AS ENUM ('Reguler', 'VIP','PVIP');

CREATE TABLE tikets (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    event_id UUID NOT NULL, 
    type_ticket type_ticket not null,
    price int not null, 
    available_tickets int not null,  
    created_at timestamptz NOT NULL DEFAULT current_timestamp,
    updated_at timestamptz NULL
)