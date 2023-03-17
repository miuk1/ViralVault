/*Sql Script to Initialize the Database*/
-- Create the Database
CREATE DATABASE IF NOT EXISTS `viralvault`;
USE `viralvault`;
-- Create the Table
CREATE TABLE vulnerable_machines (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    os TEXT NOT NULL,
    docker_image TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE vulnerabilities (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    difficulty TEXT NOT NULL,
    machine_id INTEGER NOT NULL REFERENCES vulnerable_machines(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    username TEXT NOT NULL,
    email TEXT NOT NULL,
    password_hash TEXT NOT NULL,
    password_salt TEXT NOT NULL,
    two_factor_enabled BOOLEAN NOT NULL DEFAULT false,
    session_token TEXT NOT NULL DEFAULT '',
    session_expiration BIGINT NOT NULL DEFAULT 0
);

-- Insert sample data
INSERT INTO vulnerable_machines (name, description, os, docker_image) VALUES ('Vulnerable Machine 1', 'This is a vulnerable machine', 'Linux', 'vulnerablemachine1');
INSERT INTO vulnerable_machines (name, description, os, docker_image) VALUES ('Vulnerable Machine 2', 'This is a vulnerable machine', 'Linux', 'vulnerablemachine2');

INSERT INTO vulnerabilities (name, description, difficulty, machine_id) VALUES ('Vulnerability 1', 'This is a vulnerability', 'Easy', 1);
INSERT INTO vulnerabilities (name, description, difficulty, machine_id) VALUES ('Vulnerability 2', 'This is a vulnerability', 'Medium', 1);


/*Sql Script to Initialize the Test Database*/
--Create test database and switch to it
CREATE DATABASE IF NOT EXISTS `viralvault_test`;
USE `viralvault_test`;

-- Create the Table
CREATE TABLE vulnerable_machines (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    os TEXT NOT NULL,
    docker_image TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE vulnerabilities (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    difficulty TEXT NOT NULL,
    machine_id INTEGER NOT NULL REFERENCES vulnerable_machines(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    username TEXT NOT NULL,
    email TEXT NOT NULL,
    password_hash TEXT NOT NULL,
    password_salt TEXT NOT NULL,
    two_factor_enabled BOOLEAN NOT NULL DEFAULT false,
    session_token TEXT NOT NULL DEFAULT '',
    session_expiration BIGINT NOT NULL DEFAULT 0
);


--Insert sample data
INSERT INTO vulnerable_machines (name, description, os, docker_image) VALUES ('Vulnerable Machine 1', 'This is a vulnerable machine', 'Linux', 'vulnerablemachine1');
INSERT INTO vulnerable_machines (name, description, os, docker_image) VALUES ('Vulnerable Machine 2', 'This is a vulnerable machine', 'Linux', 'vulnerablemachine2');

INSERT INTO vulnerabilities (name, description, difficulty, machine_id) VALUES ('Vulnerability 1', 'This is a vulnerability', 'Easy', 1);
INSERT INTO vulnerabilities (name, description, difficulty, machine_id) VALUES ('Vulnerability 2', 'This is a vulnerability', 'Medium', 1);

