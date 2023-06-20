BEGIN;

ALTER TABLE departments ADD COLUMN photo text;
ALTER TABLE departments ADD CONSTRAINT departments_photo_unique unique (photo);

COMMIT;