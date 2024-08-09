ALTER TABLE p_post
ADD COLUMN community_id INT NULL,
ADD FOREIGN KEY (community_id) REFERENCES p_community(id);
