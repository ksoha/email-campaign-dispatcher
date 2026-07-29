
--campaign-recipients table 
--bridge table between campaigns and recipients
--A campaign can have many recipients and a recipient can be in many campaigns

CREATE TABLE campaign_recipients(
    campaign_id UUID NOT NULL
    REFERENCES campaigns(id) ON DELETE CASCADE,

    recipient_id UUID NOT NULL
    REFERENCES recipients(id) ON DELETE CASCADE,

    PRIMARY KEY (campaign_id, recipient_id) -- composite primary key 
); 