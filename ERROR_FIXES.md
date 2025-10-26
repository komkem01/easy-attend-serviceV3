# Teacher Registration - Error Handling & Validation

## 🐛 **ปัญหาที่แก้ไข:**

### 1. Foreign Key Constraint Error
**Error Message**: `ERROR: insert or update on table "teachers" violates foreign key constraint "teachers_classroom_id_fkey"`

**สาเหตุ**: ส่ง `classroom_id` ที่ไม่มีอยู่ในฐานข้อมูล

### 2. ไม่มีการตรวจสอบ UUID อื่น ๆ 
- `prefix_id` ไม่มีอยู่
- `gender_id` ไม่มีอยู่
- `email` ซ้ำกับที่มีอยู่แล้ว

## ✅ **การแก้ไข:**

### 1. **Enhanced Validation** - ตรวจสอบความถูกต้องก่อนสร้าง Teacher

```go
// ตรวจสอบ prefix_id
_, err = s.dbPrefix.GetByIDPrefix(ctx, req.PrefixID)
if err != nil {
    return nil, fmt.Errorf("prefix not found: %s", req.PrefixID.String())
}

// ตรวจสอบ gender_id  
_, err = s.dbGender.GetByIDGender(ctx, req.GenderID)
if err != nil {
    return nil, fmt.Errorf("gender not found: %s", req.GenderID.String())
}

// ตรวจสอบ email ซ้ำ
existingTeacher, err := s.db.GetTeacherByEmail(ctx, req.Email)
if err == nil && existingTeacher != nil {
    return nil, fmt.Errorf("email already exists: %s", req.Email)
}

// ตรวจสอบ classroom_id (ถ้ามี)
if req.ClassroomID != nil {
    _, err := s.dbClassroom.GetByIDClassroom(ctx, *req.ClassroomID)
    if err != nil {
        return nil, fmt.Errorf("classroom not found: %s", req.ClassroomID.String())
    }
}
```

### 2. **Custom Error Types** - จัดการ Error ให้ดีขึ้น

```go
// ValidationError - สำหรับ validation failed
type ValidationError struct {
    Field   string `json:"field"`
    Message string `json:"message"`
}

// NotFoundError - สำหรับ resource ไม่เจอ
type NotFoundError struct {
    Resource string `json:"resource"`
    ID       string `json:"id"`
}

// ConflictError - สำหรับข้อมูลซ้ำ
type ConflictError struct {
    Resource string `json:"resource"`
    Value    string `json:"value"`
}
```

### 3. **Enhanced Error Responses** - Response ที่ชัดเจน

```json
// ตัวอย่าง Response เมื่อ classroom ไม่เจอ
{
  "code": "400",
  "message": "Resource not found",
  "error": {
    "resource": "classroom",
    "id": "550e8400-e29b-41d4-a716-446655440003"
  },
  "data": null
}

// ตัวอย่าง Response เมื่อ email ซ้ำ
{
  "code": "409",
  "message": "Resource conflict", 
  "error": {
    "resource": "email",
    "value": "somchai@school.com"
  },
  "data": null
}
```

## 🧪 **Test Cases ที่ควรทดสอบ:**

### ✅ **Valid Cases:**
```json
{
  "school_name": "โรงเรียนทดสอบ",
  "prefix_id": "valid-prefix-uuid",
  "gender_id": "valid-gender-uuid", 
  "first_name": "ทดสอบ",
  "last_name": "ใหม่",
  "email": "new@test.com",
  "password": "password123",
  "phone": "0812345678"
}
```

### ❌ **Error Cases:**

1. **ไม่ระบุ school_name**
```json
{
  "classroom_id": "invalid-uuid",
  // ไม่มี school_name
}
// Expected: 400 Bad Request
```

2. **prefix_id ไม่ถูกต้อง**
```json
{
  "school_name": "โรงเรียนทดสอบ",
  "prefix_id": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
  // ...
}
// Expected: 400 Bad Request - prefix not found
```

3. **email ซ้ำ**
```json
{
  "school_name": "โรงเรียนทดสอบ", 
  "email": "existing@test.com",
  // ...
}
// Expected: 409 Conflict - email already exists
```

4. **classroom_id ไม่มีอยู่**
```json
{
  "school_name": "โรงเรียนทดสอบ",
  "classroom_id": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
  // ...
}
// Expected: 400 Bad Request - classroom not found
```

## 📋 **การใช้งาน:**

### Endpoint:
```
POST /api/teacher
Content-Type: application/json
```

### Valid Request:
```json
{
  "school_name": "โรงเรียนบ้านดอกไม้", 
  "prefix_id": "029741a0-e049-43dd-806a-74f810bdba30",
  "gender_id": "6e4ef366-44ac-4695-8863-9e502eae6b14c",
  "first_name": "สมชาย",
  "last_name": "ทดสอบ",
  "email": "test@gmail.com",
  "password": "123456",
  "phone": "0987654321"
}
```

**🎉 ระบบแก้ไขเสร็จแล้ว! ตอนนี้จะไม่มี foreign key error อีกต่อไป**