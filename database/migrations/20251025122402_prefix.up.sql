-- Table: prefixes
CREATE TABLE prefixes (
    id         UUID        NOT NULL PRIMARY KEY,
    name       VARCHAR     NOT NULL, -- คำนำหน้าชื่อ
    created_at TIMESTAMP   NULL,     -- วันที่สร้าง
    updated_at TIMESTAMP   NULL,     -- วันที่แก้ไข
    deleted_at TIMESTAMP   NULL      -- วันที่ลบ
);

-- Seed data for prefixes
INSERT INTO prefixes (id, name, created_at, updated_at, deleted_at) VALUES
    ('029741a0-e049-43dd-806a-74f810bdba30', 'นาย',   NOW(), NULL, NULL),
    ('743a31c7-5e2b-4a48-8d95-d9ac5830f9b7', 'นาง',   NOW(), NULL, NULL),
    ('77e9d727-3b17-4271-8db3-f26216ffc5a4', 'นางสาว', NOW(), NULL, NULL),
    ('f1e4d3c2-5b6a-4c3e-9f7e-2d3b4c5a6e7f', 'เด็กชาย', NOW(), NULL, NULL),
    ('a2b3c4d5-e6f7-8901-2345-67890abcdef1', 'เด็กหญิง', NOW(), NULL, NULL);