-- race,class, user, quest, character, ระดับความยากง่าย

CREATE TABLE IF NOT EXISTS race (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    detail TEXT,
    status VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS class (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    detail TEXT,
    status VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS difficulty_level (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    detail TEXT,
    status VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS user_dnd (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50) DEFAULT 'user',
    status VARCHAR(50) DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS quest (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description VARCHAR(5000),
    difficulty_level_id INT,
    created_by INT,
    status VARCHAR(50) DEFAULT 'active',
    is_public BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (difficulty_level_id) REFERENCES difficulty_level(id),
    FOREIGN KEY (created_by) REFERENCES user_dnd(id)
);

CREATE TABLE IF NOT EXISTS character (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) UNIQUE NOT NULL,
    description VARCHAR(5000),
    class_id INT,
    race_id INT,
    created_by INT,
    status VARCHAR(50) DEFAULT 'active',
    is_public BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (class_id) REFERENCES class(id),
    FOREIGN KEY (race_id) REFERENCES race(id),
    FOREIGN KEY (created_by) REFERENCES user_dnd(id)
);  

create table if not exists image (
    id SERIAL PRIMARY KEY,
    quest_id INT,
    character_id INT,
    url VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);