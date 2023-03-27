create table if not exists users (
    id bigserial primary key,
    name varchar(40) not null,
    surname varchar(60) not null,
    father_name varchar(40) not null,
    email varchar(100) not null unique,
    login varchar(100) not null unique,
    password varchar(200) not null 
);

create table if not exists roles (
    role varchar(30) not null,
    user_id bigint references users(id) on delete cascade
);

create table if not exists departments (
    id bigserial primary key,
    name varchar(200) not null unique,
    short_name varchar(8) not null unique
);

create table if not exists groups(
    id bigserial primary key,
    name varchar(20) not null unique,
    budget boolean not null,
    education_form varchar(20) not null,
    education_start_date date not null,
    year int not null default 1,
    education_finish_date date not null,
    department_id bigint references departments(id)
);

create table if not exists studentDocuments(
    id bigserial primary key,
    doc_number varchar(60) not null unique,
    order_number varchar(60) not null,
    order_date date not null,
    study_start_date date not null
);

create table if not exists students(
    id bigserial primary key,
    user_id bigint references users(id) on delete cascade unique,
    group_id bigint references groups(id),
    student_document_id bigint references studentDocuments(id) unique,
    added_info varchar(80)
);

create table if not exists studentRequests(
    id bigserial primary key,
    student_id bigint references students(id) on delete cascade,
    status varchar(16) not null,
    department_id bigint references departments(id),
    count int not null,
    date date not null
);

