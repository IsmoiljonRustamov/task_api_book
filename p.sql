create database BooK;

create table books(
    id SERIAL PRIMARY KEY,
    title VARCHAR(50),
    author_name VARCHAR(30),
    price decimal(18,2),
    amount int,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP 
) 