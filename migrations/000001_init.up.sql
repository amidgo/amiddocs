create table if not exists users (
    id bigserial primary key,
    name varchar(40) not null,
    surname varchar(60) not null,
    "fatherName" varchar(40) not null,
    email varchar(100) not null unique,
    login varchar(100) not null unique,
    password varchar(200) not null 
);

create table if not exists roles (
    role varchar(30) not null,
    "userId" bigint references users(id) on delete cascade
);

create table if not exists departments (
    id bigserial primary key,
    name varchar(200) not null unique,
    "shortName" varchar(8) not null unique
);

create table if not exists groups(
    id bigserial primary key,
    name varchar(20) not null unique,
    budget boolean not null,
    "educationForm" varchar(20) not null,
    "educationStartDate" date not null,
    year int not null default 1,
    "educationFinishDate" date not null,
    "departmentId" bigint references departments(id)
);

create table if not exists studentDocuments(
    id bigserial primary key,
    "docNumber" varchar(60) not null unique,
    "orderNumber" varchar(60) not null,
    "orderDate" date not null,
    "studyStartDate" date not null
);

create table if not exists students(
    id bigserial primary key,
    "userId" bigint references users(id) on delete cascade unique,
    "groupId" bigint references groups(id),
    "studentDocumentId" bigint references studentDocuments(id) unique,
    "addedInfo" varchar(80)
);

create table if not exists studentRequests(
    id bigserial primary key,
    "studentId" bigint references students(id) on delete cascade,
    status varchar(16) not null,
    "departmentId" bigint references departments(id),
    count int not null,
    date date not null
);

