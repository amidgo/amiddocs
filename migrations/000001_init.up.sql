BEGIN;

create table if not exists users (
    id bigserial primary key,
    name varchar(40) not null,
    surname varchar(60) not null,
    father_name varchar(40),
    login varchar(100) not null,
    email varchar(100),
    password varchar(200) not null,

    constraint users_login_unique unique(login),
    constraint users_email_unique unique(email)
);
create table if not exists refresh_tokens (
    user_id bigserial not null,
    expired timestamp not null,
    token varchar(40) not null,

    primary key(user_id),

    constraint refresh_tokens_unique unique(token), 
    constraint fk_refresh_tokens__users foreign key (user_id) references users(id) on delete cascade
);
create table if not exists roles(
    id smallserial primary key, 
    role varchar(30) not null constraint roles_unique unique
);
create table if not exists user_roles (
    role_id smallint not null,
    user_id bigint not null,

    primary key(role_id,user_id),

    constraint fk_user_roles__roles foreign key(role_id) references roles(id),
    constraint fk_user_roles__users foreign key(user_id) references users(id) on delete cascade
);
create table if not exists departments (
    id bigserial primary key,
    name varchar(200) not null,
    short_name varchar(8) not null,  

    constraint departments_short_name_unique unique (short_name),
    constraint departments_name_unique unique (name)
);
create table if not exists study_departments (
    id bigserial primary key,
    department_id bigint not null,

    constraint study_departments_unique unique(department_id),
    constraint fk_study_departments__departments foreign key (department_id) references departments(id)
);
create table if not exists groups(
    id bigserial primary key,
    name varchar(20) not null,
    is_budget boolean not null,
    education_form varchar(20) not null,
    education_start_date date not null,
    education_year int not null,
    education_finish_date date not null,
    study_department_id bigint not null,

    constraint groups_name_unique unique(name),
    constraint fk_groups__study_departments foreign key (study_department_id) references study_departments(id)
);
create table if not exists students(
    id bigserial primary key,
    user_id bigint not null,
    group_id bigint not null,

    constraint students_user_id_unique unique(user_id),

    constraint fk_students__users foreign key (user_id) references users(id) on delete cascade,
    constraint fk_students__groups foreign key (group_id) references groups(id)
);
create table if not exists student_documents(
    id bigserial primary key,
    student_id bigint not null,
    doc_number varchar(60) not null,
    order_number varchar(60) not null,
    order_date date not null,
    study_start_date date not null,

    constraint student_documents_number_unique unique (doc_number),
    constraint student_documents_student_id_unique unique (student_id),

    constraint fk_student_documents__students foreign key (student_id) references students(id) on delete cascade
);
create table if not exists document_types(
    id serial primary key,
    refresh_time smallint not null,
    type varchar(100) not null,

    constraint document_types_unique unique (type)
);
create table if not exists document_type_roles (
    document_type_id int not null,
    role_id smallint not null,

    primary key (document_type_id, role_id),

    constraint fk_document_type_roles__document_types foreign key (document_type_id) references document_types(id),
    constraint fk_document_type_roles__roles foreign key (role_id) references roles(id)
);
create table if not exists request_status (
    id smallserial primary key,
    status varchar(16) not null
);
create table if not exists requests(
    id bigserial primary key,
    status_id smallint not null,
    count int not null,
    date timestamp not null,
    user_id bigint not null,
    department_id bigint not null,
    document_type_id int not null,

    constraint fk_requests__request_status foreign key (status_id) references request_status(id),
    constraint fk_requests__users foreign key (user_id) references users(id) on delete cascade,
    constraint fk_requests__departments foreign key (department_id) references departments(id) on delete cascade,
    constraint fk_requests__document_types foreign key (document_type_id) references document_types(id) on delete cascade
);
create table if not exists document_templates(
    department_id bigint not null,
    document_type_id int not null,
    data bytea not null,

    primary key (department_id,document_type_id),

    constraint fk_document_templates__departments foreign key (department_id) references departments(id) on delete cascade,
    constraint fk_document_templates__document_types foreign key (document_type_id) references document_types(id) on delete cascade
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

