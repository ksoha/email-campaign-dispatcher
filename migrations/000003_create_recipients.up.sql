
--creating the recipients table (id , email , name , created_at, updated_at)

CREATE TABLE recipients (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), 
     name TEXT , --can be null some csv doesnt have name 
     email TEXT NOT NULL UNIQUE, 
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);