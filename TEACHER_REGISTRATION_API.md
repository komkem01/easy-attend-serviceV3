# Teacher Registration API Examples

## การลงทะเบียนครู (Teacher Registration)

### ✅ **ข้อมูลใหม่ที่รองรับ:**

1. **school_name** - ชื่อโรงเรียน (บังคับกรอก)
   - ระบบจะหาโรงเรียนที่มีชื่อนี้อยู่แล้ว
   - ถ้าไม่มี ระบบจะสร้างโรงเรียนใหม่อัตโนมัติ

2. **classroom_id** - รหัสห้องเรียน (ไม่บังคับกรอก)
   - สามารถเป็น null หรือไม่ส่งมาก็ได้

### 📋 **API Endpoint:**
```
POST /api/teacher
```

### 🔧 **Request Body Examples:**

#### ตัวอย่างที่ 1: ลงทะเบียนโรงเรียนใหม่ (ไม่มีห้องเรียน)
```json
{
  "school_name": "โรงเรียนบ้านดอกไม้",
  "prefix_id": "550e8400-e29b-41d4-a716-446655440001",
  "gender_id": "550e8400-e29b-41d4-a716-446655440002", 
  "first_name": "สมชาย",
  "last_name": "ใจดี",
  "email": "somchai@school.com",
  "password": "password123",
  "phone": "0812345678"
}
```

#### ตัวอย่างที่ 2: ลงทะเบียนโรงเรียนเก่า (มีห้องเรียน)
```json
{
  "school_name": "โรงเรียนบ้านดอกไม้",
  "classroom_id": "550e8400-e29b-41d4-a716-446655440003",
  "prefix_id": "550e8400-e29b-41d4-a716-446655440001",
  "gender_id": "550e8400-e29b-41d4-a716-446655440002",
  "first_name": "สมหญิง", 
  "last_name": "เก่งมาก",
  "email": "somying@school.com",
  "password": "password456",
  "phone": "0823456789"
}
```

#### ตัวอย่างที่ 3: ลงทะเบียนโรงเรียนใหม่ (ไม่ระบุห้องเรียน)
```json
{
  "school_name": "โรงเรียนวัดใหญ่",
  "prefix_id": "550e8400-e29b-41d4-a716-446655440001", 
  "gender_id": "550e8400-e29b-41d4-a716-446655440002",
  "first_name": "ครูใหม่",
  "last_name": "มาสอน",
  "email": "teacher@watyard.com", 
  "password": "newpassword789",
  "phone": "0834567890"
}
```

### ✨ **Response Example:**
```json
{
  "code": "200",
  "message": "Success",
  "data": {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "school_id": "456e7890-e89b-12d3-a456-426614174001", 
    "school_name": "โรงเรียนบ้านดอกไม้",
    "classroom_id": "789e0123-e89b-12d3-a456-426614174002",
    "prefix_id": "550e8400-e29b-41d4-a716-446655440001",
    "gender_id": "550e8400-e29b-41d4-a716-446655440002",
    "first_name": "สมชาย", 
    "last_name": "ใจดี",
    "email": "somchai@school.com",
    "phone": "0812345678",
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token": "def50200a1b2c3d4e5f6...",
    "token_type": "Bearer",
    "expires_at": 1698331200
  }
}
```

### 🔄 **ระบบการทำงาน:**

1. **รับ Request** → ตรวจสอบข้อมูลที่จำเป็น
2. **ค้นหาโรงเรียน** → `GetSchoolByName(school_name)`
   - **ถ้าเจอ** → ใช้ school_id ที่มีอยู่
   - **ถ้าไม่เจอ** → สร้างโรงเรียนใหม่ด้วย `CreateSchool(school_name, "", "")`
3. **จัดการห้องเรียน** → 
   - **ถ้ามี classroom_id** → ตรวจสอบและใช้
   - **ถ้าไม่มี** → เก็บเป็น null
4. **สร้างครู** → ใช้ school_id และ classroom_id (อาจเป็น null)
5. **Generate Token** → สร้าง JWT token สำหรับ login อัตโนมัติ
6. **Response** → ส่งข้อมูลครูและ token กลับ

### 📝 **ข้อดี:**

- **ง่ายต่อการใช้งาน** - หน้าบ้านแค่ส่งชื่อโรงเรียนมา
- **Auto-creation** - ไม่ต้องสร้างโรงเรียนล่วงหน้า
- **Flexible** - classroom_id เป็น optional
- **Immediate Access** - ได้ token เลยหลังลงทะเบียน

### 🛡️ **Security Features:**

- Password hashing ด้วย Argon2id
- JWT token generation
- Email validation
- Input validation

**🎉 ระบบพร้อมใช้งาน! ครูสามารถลงทะเบียนได้โดยระบุเพียงชื่อโรงเรียน**