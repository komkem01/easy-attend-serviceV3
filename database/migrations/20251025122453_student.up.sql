CREATE TABLE students (
    id            UUID        NOT NULL,
    school_id     UUID        NULL,
    classroom_id  UUID        NULL,
    prefix_id     UUID        NULL,
    gender_id     UUID        NULL,
    student_code  VARCHAR     NOT NULL, -- auto-generated, start from 001
    first_name    VARCHAR     NOT NULL,
    last_name     VARCHAR     NOT NULL,
    phone         VARCHAR     NOT NULL,
    created_at    TIMESTAMP   NULL,
    updated_at    TIMESTAMP   NULL,
    deleted_at    TIMESTAMP   NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (school_id)    REFERENCES schools(id),
    FOREIGN KEY (classroom_id) REFERENCES classrooms(id),
    FOREIGN KEY (prefix_id)    REFERENCES prefixes(id),
    FOREIGN KEY (gender_id)    REFERENCES genders(id)
);

-- Add table comment
COMMENT ON TABLE students IS 'ข้อมูลนักเรียน';

-- Add column comments
COMMENT ON COLUMN students.student_code IS 'รหัสนักเรียน';
COMMENT ON COLUMN students.first_name IS 'ชื่อ';
COMMENT ON COLUMN students.last_name IS 'นามสกุล';
COMMENT ON COLUMN students.phone IS 'เบอร์โทรศัพท์';
COMMENT ON COLUMN students.created_at IS 'วันที่สร้าง';
COMMENT ON COLUMN students.updated_at IS 'วันที่แก้ไข';
COMMENT ON COLUMN students.deleted_at IS 'วันที่ลบ';
COMMENT ON COLUMN students.school_id IS 'รหัสโรงเรียน';
COMMENT ON COLUMN students.classroom_id IS 'รหัสห้องเรียน';
COMMENT ON COLUMN students.prefix_id IS 'คำนำหน้าชื่อ';
COMMENT ON COLUMN students.gender_id IS 'เพศ';