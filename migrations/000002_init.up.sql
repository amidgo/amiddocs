BEGIN;

ALTER TABLE student_documents RENAME COLUMN study_start_date TO education_start_date;

COMMIT;

