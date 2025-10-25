CREATE TABLE classrooms (
    id         UUID      NOT NULL,
    school_id  UUID      NULL,
    name       VARCHAR   NOT NULL,
    created_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (school_id) REFERENCES schools(id)
);

-- Add table comment
COMMENT ON TABLE classrooms IS 'ข้อมูลห้องเรียน';

-- Add column comments
COMMENT ON COLUMN classrooms.name IS 'ชื่อห้องเรียน';
COMMENT ON COLUMN classrooms.created_at IS 'วันที่สร้าง';
COMMENT ON COLUMN classrooms.updated_at IS 'วันที่แก้ไข';
COMMENT ON COLUMN classrooms.deleted_at IS 'วันที่ลบ';
COMMENT ON COLUMN classrooms.school_id IS 'รหัสโรงเรียน';