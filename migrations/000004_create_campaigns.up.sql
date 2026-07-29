
--campaign table 
--a user can have multiple campaigns
--every campaign must belong to a single user 

CREATE TABLE campaigns(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), 
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,  --foreign key to users table
    name TEXT NOT NULL, 
    subject TEXT NOT NULL, 
    body TEXT NOT NULL, 
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
)