# Entity Relationship Diagram (ERD)

## Users
- **name**: string
- **position**: string
- **role**: string
- **email**: string
- **password**: string
- **isActive**: boolean
- **createdAt**: timestamp
- **updatedAt**: timestamp

## Tasks
- **title**: string
- **date**: date
- **priority**: string (enum: `high`, `medium`, `normal`, `low`)
- **stage**: string (enum: `todo`, `in progress`, `completed`)
- **teams**: array of user IDs
- **isDelete**: boolean
- **createdAt**: timestamp
- **updatedAt**: timestamp

## Activities
- **type**: string (enum: `assigned`, `started`, `in progress`, `bug`, `completed`)
- **activity**: string
- **date**: date
- **by**: user ID (string)
- **createdAt**: timestamp
- **updatedAt**: timestamp

## Notifications
- **teams**: array of user IDs
- **text**: string
- **notiType**: string (enum: `alert`, `message`)
- **isRead**: array of user IDs
- **createdAt**: timestamp
- **updatedAt**: timestamp

