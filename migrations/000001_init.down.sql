BEGIN;

drop table if exists document_types cascade;
drop table if exists document_templates cascade;
drop table if exists request_status cascade;
drop table if exists requests;
drop table if exists students cascade;
drop table if exists student_documents cascade;
drop table if exists groups cascade;
drop table if exists study_departments cascade;
drop table if exists departments cascade;
drop table if exists user_roles cascade;
drop table if exists roles cascade;
drop table if exists refresh_tokens cascade;
drop table if exists users cascade; 
drop table if exists document_type_roles cascade;

COMMIT;




