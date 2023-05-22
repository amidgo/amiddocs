BEGIN;

ALTER TABLE student_documents RENAME COLUMN education_start_date TO study_start_date;

COMMIT;