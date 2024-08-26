CREATE TABLE p_report (
                          id INT AUTO_INCREMENT PRIMARY KEY,
                          post_id INT NOT NULL,
                          reporter_id INT NOT NULL,
                          comment_id INT NOT NULL,
                          reason VARCHAR(255) NOT NULL,
                          status ENUM('pending', 'reviewed', 'resolved') DEFAULT 'pending',
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
) ENGINE=InnoDB;