# การปรับปรุง classroom_id ให้รองรับค่า NULL

## สรุปการเปลี่ยนแปลง

ปรับปรุงระบบให้ `classroom_id` ของทั้ง **Student** และ **Teacher** สามารถเป็น `null` ในฐานข้อมูลได้เมื่อผู้ใช้งานไม่ได้ส่งค่ามาตอนลงทะเบียน

## ไฟล์ที่แก้ไข

### 1. Student Entity (`app/modules/entities/ent/student.ent.go`)
- **แก้ไข**: เอา `notnull` ออกจาก `ClassroomID` field
- **ก่อน**: `ClassroomID uuid.UUID \`bun:"type:uuid,notnull"\``
- **หลัง**: `ClassroomID uuid.UUID \`bun:"type:uuid"\``

### 2. Student Create Controller (`app/modules/student/student-create.ctl.go`)
- **แก้ไข Request**: เอา `binding:"required"` ออกจาก `classroom_id`
- **แก้ไข Response**: เปลี่ยน `ClassroomID` เป็น `*uuid.UUID` พร้อม `omitempty`
- **แก้ไข Logic**: เพิ่มการตรวจสอบว่า `classroom_id` ว่างหรือไม่ก่อน parse UUID
- **แก้ไข Response Assignment**: ใช้ pointer เมื่อ `ClassroomID != uuid.Nil`

### 3. Student Create Service (`app/modules/student/student-create.svc.go`)
- **แก้ไข**: เพิ่มคอมเมนต์อธิบายว่าจะใช้ `uuid.Nil` สำหรับกรณีไม่มีค่า

### 4. Student Update Controller (`app/modules/student/student-update.ctl.go`)
- **แก้ไข Request**: เอา `binding:"required"` ออกจาก `classroom_id`
- **แก้ไข Logic**: เพิ่มการตรวจสอบว่า `classroom_id` ว่างหรือไม่ก่อน parse UUID

### 5. Student DTO (`app/modules/entities/dto/student.dto.go`)
- **แก้ไข**: เอา `validate:"required"` ออกจาก `Classroom` field ใน `StudentCreateRequest` และ `StudentUpdateRequest`

### 6. Student List Service (`app/modules/student/student-list.svc.go`)
- **แก้ไข Response**: เปลี่ยน `ClassroomID` เป็น `*uuid.UUID` พร้อม `omitempty`
- **แก้ไข Logic**: ใช้ pointer เมื่อ `ClassroomID != uuid.Nil`

### 7. Student Info Service (`app/modules/student/student-info.svc.go`)
- **แก้ไข Response**: เปลี่ยน `ClassroomID` และ `ClassroomName` เป็น pointer พร้อม `omitempty`
- **แก้ไข Logic**: จัดการกรณี `ClassroomID == uuid.Nil` โดยไม่ไปดึงข้อมูล classroom

### 8. Teacher Entity (`app/modules/entities/ent/teacher.ent.go`)
- **แก้ไข**: เปลี่ยน `ClassroomID` เป็น `*uuid.UUID` เพื่อรองรับ `NULL` อย่างสมบูรณ์

### 9. Teacher DTO (`app/modules/entities/dto/teacher.dto.go`)
- **แก้ไข**: เปลี่ยน `ClassroomID` เป็น `*uuid.UUID` ใน `TeacherCreateRequest` และ `TeacherUpdateRequest`

### 10. Teacher Create Service (`app/modules/teacher/teacher-create.svc.go`)
- **แก้ไข Logic**: ใช้ pointer แทน `uuid.Nil` เพื่อส่งค่า `nil` ไปยัง database

### 11. Teacher List Service (`app/modules/teacher/teacher-list.svc.go`)
- **แก้ไข Response**: เปลี่ยน `ClassroomID` เป็น `*uuid.UUID` พร้อม `omitempty`
- **แก้ไข Logic**: ใช้ pointer ตรงๆ จาก entity

### 12. Teacher Info Service (`app/modules/teacher/teacher-info.svc.go`)
- **แก้ไข Response**: เปลี่ยน `ClassroomID` และ `ClassroomName` เป็น pointer พร้อม `omitempty`
- **แก้ไข Logic**: จัดการกรณี `ClassroomID == nil` โดยไม่ไปดึงข้อมูล classroom

### 13. Teacher Update Service (`app/modules/teacher/teacher-update.svc.go`)
- **แก้ไข Request**: เปลี่ยน `ClassroomID` เป็น `*uuid.UUID`

### 14. Teacher Update Controller (`app/modules/teacher/teacher-update.ctl.go`)
- **แก้ไข Logic**: สร้าง pointer เมื่อมีค่า หรือส่ง `nil` เมื่อไม่มีค่า

### 15. Teacher Login Service (`app/modules/teacher/teacher-login.svc.go`)
- **แก้ไข Response**: เปลี่ยน `ClassroomID` เป็น `*uuid.UUID` พร้อม `omitempty`

## การทำงานหลังการปรับปรุง

### สำหรับ API Request
1. **Student Create/Update**: `classroom_id` เป็น optional field
   ```json
   {
     "school_id": "uuid",
     "classroom_id": "", // หรือไม่ส่งเลย
     "prefix_id": "uuid",
     "gender_id": "uuid",
     "student_code": "001",
     "first_name": "John",
     "last_name": "Doe"
   }
   ```

2. **Teacher Registration**: `classroom_id` เป็น optional field
   ```json
   {
     "school_name": "โรงเรียนตัวอย่าง",
     "classroom_id": "", // หรือไม่ส่งเลย
     "prefix_id": "uuid",
     "gender_id": "uuid",
     "first_name": "Jane",
     "last_name": "Smith",
     "email": "jane@example.com",
     "password": "password123"
   }
   ```

### สำหรับ API Response
1. **Student Create/List/Info**: `classroom_id` จะเป็น `null` หรือไม่แสดงเมื่อไม่มีค่า
   ```json
   {
     "id": "uuid",
     "school_id": "uuid",
     // "classroom_id": null (หรือไม่แสดง)
     "prefix_id": "uuid",
     "gender_id": "uuid",
     "student_code": "001",
     "first_name": "John",
     "last_name": "Doe"
   }
   ```

2. **Teacher Registration/List/Info**: `classroom_id` จะเป็น `null` หรือไม่แสดงเมื่อไม่มีค่า
   ```json
   {
     "id": "uuid",
     "school_id": "uuid",
     "school_name": "โรงเรียนตัวอย่าง",
     // "classroom_id": null (หรือไม่แสดง)
     // "classroom_name": null (หรือไม่แสดง)
     "prefix_id": "uuid",
     "gender_id": "uuid",
     "first_name": "Jane",
     "last_name": "Smith",
     "email": "jane@example.com"
   }
   ```

### การบันทึกในฐานข้อมูล
- **Student**: เมื่อ `classroom_id` ไม่ได้ส่งมา จะใช้ `uuid.Nil` ซึ่ง Bun ORM จะแปลงเป็น `NULL` ในฐานข้อมูลโดยอัตโนมัติ
- **Teacher**: เมื่อ `classroom_id` ไม่ได้ส่งมา จะส่ง `nil` pointer ซึ่งจะบันทึกเป็น `NULL` ในฐานข้อมูล
- Database schema รองรับ `NULL` อยู่แล้วตาม migration file

## การทดสอบ

ควรทดสอบกรณีต่อไปนี้:

### Student Module:
1. สร้าง student โดยไม่ส่ง `classroom_id`
2. สร้าง student โดยส่ง `classroom_id` เป็น empty string
3. สร้าง student โดยส่ง `classroom_id` เป็น valid UUID
4. อัพเดต student โดยเปลี่ยน `classroom_id` เป็น empty เพื่อลบค่า
5. ตรวจสอบ response ใน list และ info API

### Teacher Module:
1. ลงทะเบียน teacher โดยไม่ส่ง `classroom_id`
2. ลงทะเบียน teacher โดยส่ง `classroom_id` เป็น empty string
3. ลงทะเบียน teacher โดยส่ง `classroom_id` เป็น valid UUID
4. อัพเดต teacher โดยเปลี่ยน `classroom_id` เป็น empty เพื่อลบค่า
5. ตรวจสอบ response ใน list และ info API
6. ตรวจสอบ login response ว่า `classroom_id` แสดงถูกต้อง

## หมายเหตุ

- การเปลี่ยนแปลงนี้ไม่กระทบต่อ database migration เนื่องจาก schema รองรับ `NULL` อยู่แล้ว
- **Student Module**: ใช้ `uuid.Nil` ในการจัดเก็บและใช้ `*uuid.UUID` ใน response เพื่อ JSON serialization
- **Teacher Module**: ใช้ `*uuid.UUID` ทั้งในการจัดเก็บและ response เพื่อให้จัดการ `NULL` ได้อย่างสมบูรณ์
- วิธี Teacher จะดีกว่าเพราะใช้ pointer แท้จริงซึ่ง ORM จะแปลงเป็น `NULL` ได้ถูกต้อง