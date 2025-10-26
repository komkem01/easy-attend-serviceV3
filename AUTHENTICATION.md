# Teacher Authentication System

This document describes the authentication system implemented for the teacher module.

## Features

- **Password Hashing**: Uses Argon2id algorithm for secure password storage
- **Token Generation**: Simple token-based authentication system
- **Middleware Protection**: Authentication middleware for protected routes
- **Registration & Login**: Complete teacher registration and login flow

## Password Security

The system uses Argon2id with the following parameters:
- Memory: 64 MB
- Iterations: 3
- Parallelism: 2
- Salt Length: 16 bytes
- Key Length: 32 bytes

## API Endpoints

### Public Endpoints (No Authentication Required)

#### Teacher Registration
```http
POST /api/teacher
Content-Type: application/json

{
  "school_id": "uuid",
  "classroom_id": "uuid",
  "prefix_id": "uuid",
  "gender_id": "uuid",
  "first_name": "John",
  "last_name": "Doe",
  "email": "john.doe@example.com",
  "password": "securepassword123",
  "phone": "0123456789"
}
```

**Response:**
```json
{
  "message": "Teacher created successfully",
  "data": {
    "id": "uuid",
    "school_id": "uuid",
    "classroom_id": "uuid",
    "prefix_id": "uuid",
    "gender_id": "uuid",
    "first_name": "John",
    "last_name": "Doe",
    "email": "john.doe@example.com",
    "phone": "0123456789",
    "access_token": "random_token_string",
    "refresh_token": "random_refresh_token_string",
    "token_type": "Bearer",
    "expires_at": "2024-01-26T12:00:00Z"
  }
}
```

#### Teacher Login
```http
POST /api/teacher/login
Content-Type: application/json

{
  "email": "john.doe@example.com",
  "password": "securepassword123"
}
```

**Response:**
```json
{
  "message": "Login successful",
  "data": {
    "id": "uuid",
    "school_id": "uuid",
    "classroom_id": "uuid",
    "prefix_id": "uuid",
    "gender_id": "uuid",
    "first_name": "John",
    "last_name": "Doe",
    "email": "john.doe@example.com",
    "phone": "0123456789",
    "access_token": "random_token_string",
    "refresh_token": "random_refresh_token_string",
    "token_type": "Bearer",
    "expires_at": "2024-01-26T12:00:00Z"
  }
}
```

### Protected Endpoints (Authentication Required)

All protected endpoints require the `Authorization` header:

```http
Authorization: Bearer <access_token>
```

#### Get Teachers List
```http
GET /api/teacher
Authorization: Bearer <access_token>
```

#### Get Teacher by ID
```http
GET /api/teacher/{id}
Authorization: Bearer <access_token>
```

#### Update Teacher
```http
PATCH /api/teacher/{id}
Authorization: Bearer <access_token>
Content-Type: application/json

{
  "first_name": "Updated Name",
  "last_name": "Updated Lastname"
}
```

#### Delete Teacher
```http
DELETE /api/teacher/{id}
Authorization: Bearer <access_token>
```

## Error Responses

### Authentication Errors

#### Missing Authorization Header
```json
{
  "error": "Unauthorized",
  "message": "Authorization header is required"
}
```

#### Invalid Token Format
```json
{
  "error": "Unauthorized",
  "message": "Invalid authorization header format"
}
```

#### Invalid or Expired Token
```json
{
  "error": "Unauthorized",
  "message": "Invalid or expired token"
}
```

### Login Errors

#### Invalid Credentials
```json
{
  "error": "Login failed",
  "message": "invalid email or password"
}
```

## Implementation Details

### Password Hashing

The system uses the `PasswordHasher` utility:

```go
passwordHasher := auth.NewPasswordHasher()
hashedPassword, err := passwordHasher.HashPassword(plainPassword)
```

### Token Generation

The system uses the `TokenManager` utility:

```go
tokenMgr := auth.NewTokenManager("your-secret-key-here")
tokens, err := tokenMgr.GenerateTokenPair(userID, email, firstName, lastName, "teacher")
```

### Middleware Usage

Protected routes use the authentication middleware:

```go
teacherProtected := r.Group("/teacher")
teacherProtected.Use(auth.RequireAuth())
```

## Security Notes

1. **Password Storage**: Passwords are never stored in plain text
2. **Token Expiry**: Access tokens expire after 24 hours
3. **Refresh Tokens**: Refresh tokens expire after 7 days
4. **Email Validation**: Login requires valid email format
5. **Secure Headers**: Authentication uses Bearer token format

## TODO

1. Move secret key to configuration file
2. Implement proper token storage (Redis/Database)
3. Add token refresh endpoint
4. Add password reset functionality
5. Implement role-based access control (RBAC)
6. Add rate limiting for login attempts
7. Add email verification for registration

## Usage Example

1. **Register a new teacher**:
   ```bash
   curl -X POST http://localhost:8080/api/teacher \
     -H "Content-Type: application/json" \
     -d '{
       "school_id": "uuid",
       "classroom_id": "uuid", 
       "prefix_id": "uuid",
       "gender_id": "uuid",
       "first_name": "John",
       "last_name": "Doe",
       "email": "john.doe@example.com",
       "password": "securepassword123",
       "phone": "0123456789"
     }'
   ```

2. **Login with registered teacher**:
   ```bash
   curl -X POST http://localhost:8080/api/teacher/login \
     -H "Content-Type: application/json" \
     -d '{
       "email": "john.doe@example.com",
       "password": "securepassword123"
     }'
   ```

3. **Access protected endpoint**:
   ```bash
   curl -X GET http://localhost:8080/api/teacher \
     -H "Authorization: Bearer <access_token_from_login>"
   ```