-- gender.sql

-- Create genders table
CREATE TABLE genders (
    id         UUID        NOT NULL PRIMARY KEY,
    name       VARCHAR     NOT NULL,
    created_at TIMESTAMP   NULL,
    updated_at TIMESTAMP   NULL,
    deleted_at TIMESTAMP   NULL
);

-- Add table comment
COMMENT ON TABLE genders IS 'ข้อมูลเพศ';

-- Add column comments
COMMENT ON COLUMN genders.name IS 'เพศ';
COMMENT ON COLUMN genders.created_at IS 'วันที่สร้าง';
COMMENT ON COLUMN genders.updated_at IS 'วันที่แก้ไข';
COMMENT ON COLUMN genders.deleted_at IS 'วันที่ลบ';

-- Seed initial genders
INSERT INTO genders (id, name, created_at)
VALUES
    ('6e4ef366-44ac-4695-8063-9e502eaeb14c', 'ชาย', NOW()),
    ('0bdad91e-8972-479e-aab9-ea05f87cdf80', 'หญิง', NOW()),
    ('dd738d83-6f7f-4977-92f7-943ab04858cb', 'ไม่ระบุ', NOW());