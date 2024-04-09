-- 001_users_ytable.sql

-- Crear la tabla de usuarios
CREATE TABLE IF NOT EXISTS users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    role_id INT NOT NULL,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    email VARCHAR(255) UNIQUE NOT NULL,
    email_confirmed BOOLEAN DEFAULT FALSE,
    password VARCHAR(255),
    last_login DATETIME,
    created DATETIME DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255),
    updated DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    updated_by VARCHAR(255),
    telephone VARCHAR(20),
    status VARCHAR(20),
    status_updated DATETIME,
    status_updated_by VARCHAR(255),
    FOREIGN KEY (role_id) REFERENCES roles(id)
);