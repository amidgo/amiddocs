BEGIN;

create table if not exists users (
    id bigserial primary key,
    name varchar(40) not null,
    surname varchar(60) not null,
    father_name varchar(40),
    login varchar(100) not null unique,
    email varchar(100) unique,
    password varchar(200) not null
);
create table if not exists refresh_tokens (
    user_id bigserial not null references users(id) on delete cascade,
    expired timestamp not null,
    token varchar(40) not null unique,
    primary key(user_id) 
);
create table if not exists roles(
    id smallserial primary key, 
    role varchar(30) unique not null
);
create table if not exists user_roles (
    role_id smallint not null references roles(id),
    user_id bigint not null references users(id) on delete cascade,
    primary key(role_id,user_id)
);
create table if not exists departments (
    id bigserial primary key,
    name varchar(200) not null unique,
    short_name varchar(8) not null unique
);
create table if not exists study_departments (
    id bigserial primary key,
    department_id bigint not null references departments(id)
);
create table if not exists groups(
    id bigserial primary key,
    name varchar(20) not null unique,
    is_budget boolean not null,
    education_form varchar(20) not null,
    education_start_date date not null,
    education_year int not null,
    education_finish_date date not null,
    study_department_id bigint not null references study_departments(id)
);
create table if not exists students(
    id bigserial primary key,
    user_id bigint not null references users(id) on delete cascade unique,
    group_id bigint not null references groups(id)
);
create table if not exists student_documents(
    id bigserial primary key,
    student_id bigint not null unique references students(id) on delete cascade,
    doc_number varchar(60) not null unique,
    order_number varchar(60) not null,
    order_date date not null,
    study_start_date date not null
);
create table if not exists document_types(
    id serial primary key,
    type varchar(100) not null unique,
    refresh_time smallint not null
);
create table if not exists document_type_roles (
    document_type_id int not null references document_types(id),
    role_id smallint not null references roles(id),
    primary key (document_type_id, role_id)
);
create table if not exists request_status (
    id smallserial primary key,
    status varchar(16) not null
);
create table if not exists requests(
    id bigserial primary key,
    status_id smallint not null references request_status(id),
    count int not null,
    date timestamp not null,
    user_id bigint not null references users(id) on delete cascade,
    department_id bigint not null references departments(id) on delete cascade,
    document_type_id int not null references document_types(id) on delete cascade
);
create table if not exists document_templates(
    department_id bigint not null references departments(id) on delete cascade,
    document_type_id int not null references document_types(id) on delete cascade,
    data bytea not null,
    primary key (department_id,document_type_id)
);



COMMIT;

BEGIN;


INSERT INTO request_status (status) VALUES ('SEND'), ('IN_PROGRESS'), ('DONE');

INSERT INTO roles (role) VALUES ('SECRETARY'), ('ADMIN');

WITH 
student_role_insert AS (
    INSERT INTO roles (role) VALUES('STUDENT') RETURNING id as student_role_id
),
budget_type_insert AS (
    INSERT INTO document_types (type, refresh_time) VALUES ('STUDY_DOCUMENT_BUDGET',7) RETURNING id as budget_type_id
),
no_budget_type_insert AS (
    INSERT INTO document_types (type, refresh_time) VALUES ('STUDY_DOCUMENT_NO_BUDGET',7) RETURNING id as no_budget_type_id
)
INSERT INTO document_type_roles (document_type_id,role_id) VALUES (
    (select budget_type_id from budget_type_insert),
    (select student_role_id from student_role_insert)
),
(
    (select no_budget_type_id from no_budget_type_insert),
    (select student_role_id from student_role_insert)
);

COMMIT;

