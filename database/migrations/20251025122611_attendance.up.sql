-- สร้าง ENUM type สำหรับ attendance_status
CREATE TYPE attendance_status AS ENUM (
    'pending',
    'present',
    'absent',
    'late',
    'excused'
);

CREATE TABLE attendances (
    id            UUID              NOT NULL,
    student_id    UUID              NULL REFERENCES students(id),
    teacher_id    UUID              NULL REFERENCES teachers(id),
    classroom_id  UUID              NULL REFERENCES classrooms(id),
    date          DATE              NOT NULL,
    time          TIME              NOT NULL,
    status        attendance_status NOT NULL DEFAULT 'pending',
    created_at    TIMESTAMP         NULL,
    updated_at    TIMESTAMP         NULL,
    deleted_at    TIMESTAMP         NULL,
    PRIMARY KEY (id)
);

-- Add table comment
COMMENT ON TABLE attendances IS 'ข้อมูลการเข้าเรียน';

-- Add column comments
COMMENT ON COLUMN attendances.student_id IS 'รหัสนักเรียน';
COMMENT ON COLUMN attendances.teacher_id IS 'รหัสครู';
COMMENT ON COLUMN attendances.classroom_id IS 'รหัสห้องเรียน';
COMMENT ON COLUMN attendances.date IS 'วันที่เข้าเรียน';
COMMENT ON COLUMN attendances.time IS 'เวลาที่เข้าเรียน';
COMMENT ON COLUMN attendances.status IS 'สถานะการเข้าเรียน';
COMMENT ON COLUMN attendances.created_at IS 'วันที่สร้าง';
COMMENT ON COLUMN attendances.updated_at IS 'วันที่แก้ไข';
COMMENT ON COLUMN attendances.deleted_at IS 'วันที่ลบ';