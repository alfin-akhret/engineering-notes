# Timestamp with timezone

**Main principle:**  Always store and process time in UTC at backend and database. Convert to local time only in the frontend.

Follow this practice when using datetime:
1. Database: always use `timestamptz` and `UTC` as default timezone
```
ALTER TABLE orders
ALTER COLUMN created_at
SET DATA TYPE TIMESTAMPTZ
USING created_at AT TIME ZONE 'UTC';
```

Reminder: Ensure your database default timezone is UTC, run this on your DB
```SHOW timezone;```


2. Backend code: always use UTC before inserting to db
for example in golang you can use:
```
// This ensures all timestamps stored in the DB are UTC
t := time.Now().UTC()
```

3. When sending response to the front end use `ISO 8601 UTC String format`
```
// Always use ISO 8601 format with 'Z' to indicate UTC
{
  "id": "user-1",
  "created_at": "2026-03-19T15:29:57Z",
  "updated_at": "2026-03-19T15:29:57Z"
} 
```
Reminder: frontend can convert to local timezone as needed:
```
const createdAtLocal = new Date(user.created_at).toLocaleString(); 
```

## Common mistakes

- Storing timestamps in local time instead of UTC
- Using `timestamp` without timezone in DB
- Sending timestamps to frontend without specifying UTC (ISO 8601)
