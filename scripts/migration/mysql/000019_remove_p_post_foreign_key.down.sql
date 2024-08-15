ALTER TABLE paopao.p_post
    ADD CONSTRAINT p_post_ibfk_1
        FOREIGN KEY (community_id)
            REFERENCES p_community(id);