-- +goose Up
SELECT 'up SQL query';
CREATE TABLE users ( id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) not null,
    email VARCHAR(255) not null unique,
    password VARCHAR(255) not null
);
create table teams (
    id int primary key AUTO_INCREMENT,
    name VARCHAR(255) not null,
    created_by int,
    foreign key (created_by) references users(id)
);
create table team_members(
    id int primary key auto_increment,
    team_id int,
    user_id int,
    role VARCHAR(255) default('member'),
    foreign key (team_id) references teams(id),
    foreign key (user_id) references users(id)

);
create table tasks (
    id int primary key auto_increment,
    name VARCHAR(255) not null ,
    description VARCHAR(255),
    team_id int,
    assignee_id int,
    created_by int,
    status VARCHAR(255),
    foreign key (team_id) references teams(id),
    foreign key (assignee_id) references users(id),
    foreign key (created_by) references users(id)
);
create table task_history(
    id int primary key auto_increment,
    status VARCHAR(255),
    name VARCHAR(255),
    finished_date date,
    task_id int,
    changed_by int,
    foreign key (task_id) references tasks(id),
    foreign key (changed_by) references users(id)
);
create table task_comments(
    id int primary key auto_increment,
    comment VARCHAR(255),
    task_id int,
    user_id int,
    foreign key (task_id) references tasks(id),
    foreign key (user_id) references users(id)
);


-- +goose Down
SELECT 'down SQL query';
drop table users;
drop table teams;
drop table team_members;
drop table tasks;
drop table task_history;
drop table task_comments;
