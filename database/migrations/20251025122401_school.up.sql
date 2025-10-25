CREATE TABLE schools (
    id         UUID      NOT NULL,
    name       VARCHAR   NOT NULL,
    address    TEXT      NULL,
    phone      VARCHAR   NULL,
    created_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id)
);

-- Add table comment
COMMENT ON TABLE schools IS 'ข้อมูลโรงเรียน';

-- Add column comments
COMMENT ON COLUMN schools.name IS 'ชื่อโรงเรียน';
COMMENT ON COLUMN schools.address IS 'ที่อยู่โรงเรียน';
COMMENT ON COLUMN schools.phone IS 'เบอร์โทรศัพท์โรงเรียน';
COMMENT ON COLUMN schools.created_at IS 'วันที่สร้าง';
COMMENT ON COLUMN schools.updated_at IS 'วันที่แก้ไข';
COMMENT ON COLUMN schools.deleted_at IS 'วันที่ลบ';
