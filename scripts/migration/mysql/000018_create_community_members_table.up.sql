CREATE TABLE p_community_members (
    id INT AUTO_INCREMENT PRIMARY KEY,
    community_id INT NOT NULL,
    user_id BIGINT NOT NULL,
    role ENUM('member', 'moderator', 'admin') NOT NULL,
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (community_id) REFERENCES p_community(id),
    FOREIGN KEY (user_id) REFERENCES p_user(id)
);
