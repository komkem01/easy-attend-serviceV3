CREATE TABLE teachers (
    id           UUID      NOT NULL,
    school_id    UUID      NULL,
    classroom_id UUID      NULL,
    prefix_id    UUID      NULL,
    gender_id    UUID      NULL,
    first_name   VARCHAR   NOT NULL,
    last_name    VARCHAR   NOT NULL,
    email        VARCHAR   NOT NULL,
    password     VARCHAR   NOT NULL,
    phone        VARCHAR   NOT NULL,
    created_at   TIMESTAMP NULL,
    updated_at   TIMESTAMP NULL,
    deleted_at   TIMESTAMP NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (school_id)    REFERENCES schools(id),
    FOREIGN KEY (classroom_id) REFERENCES classrooms(id),
    FOREIGN KEY (prefix_id)    REFERENCES prefixes(id),
    FOREIGN KEY (gender_id)    REFERENCES genders(id)
);

-- Add table comment
COMMENT ON TABLE teachers IS 'ข้อมูลครูผู้สอน';

-- Add column comments
COMMENT ON COLUMN teachers.first_name IS 'ชื่อ';
COMMENT ON COLUMN teachers.last_name IS 'นามสกุล';
COMMENT ON COLUMN teachers.email IS 'อีเมล';
COMMENT ON COLUMN teachers.password IS 'รหัสผ่าน';
COMMENT ON COLUMN teachers.phone IS 'เบอร์โทรศัพท์';
COMMENT ON COLUMN teachers.created_at IS 'วันที่สร้าง';
COMMENT ON COLUMN teachers.updated_at IS 'วันที่แก้ไข';
COMMENT ON COLUMN teachers.deleted_at IS 'วันที่ลบ';
COMMENT ON COLUMN teachers.school_id IS 'รหัสโรงเรียน';
COMMENT ON COLUMN teachers.classroom_id IS 'รหัสห้องเรียน';
COMMENT ON COLUMN teachers.prefix_id IS 'คำนำหน้าชื่อ';
COMMENT ON COLUMN teachers.gender_id IS 'เพศ';

CREATE UNIQUE INDEX uq_teachers_email ON teachers (email);