CREATE TABLE classroom_members (
    id           UUID      NOT NULL,
    classroom_id UUID      NULL,
    student_id   UUID      NULL,
    teacher_id   UUID      NULL,
    created_at   TIMESTAMP NULL,
    updated_at   TIMESTAMP NULL,
    deleted_at   TIMESTAMP NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (classroom_id) REFERENCES classrooms(id),
    FOREIGN KEY (student_id) REFERENCES students(id),
    FOREIGN KEY (teacher_id) REFERENCES teachers(id)
);

-- Add table comment
COMMENT ON TABLE classroom_members IS 'ข้อมูลสมาชิกห้องเรียน';

-- Add column comments
COMMENT ON COLUMN classroom_members.classroom_id IS 'รหัสห้องเรียน';
COMMENT ON COLUMN classroom_members.student_id IS 'รหัสนักเรียน';
COMMENT ON COLUMN classroom_members.teacher_id IS 'รหัสครู';
COMMENT ON COLUMN classroom_members.created_at IS 'วันที่สร้าง';
COMMENT ON COLUMN classroom_members.updated_at IS 'วันที่แก้ไข';
COMMENT ON COLUMN classroom_members.deleted_at IS 'วันที่ลบ';