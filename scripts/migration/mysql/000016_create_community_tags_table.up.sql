CREATE TABLE p_community_tags (
    id INT AUTO_INCREMENT PRIMARY KEY,
    community_id INT NOT NULL,
    FOREIGN KEY (community_id) REFERENCES p_community(id),
    tag_name VARCHAR(255) NOT NULL
);
