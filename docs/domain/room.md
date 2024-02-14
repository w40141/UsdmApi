# Room ER diagram

```marmeid
erDiagram
ACCOUNT {
  string id
  string name
  string email
}
SAGA {
  string id
  string title
  string description
  date created_at
  date updated_at
}
PARTICIPANT {
  member_id member
  authority_id authority
  saga_id saga
}
AUTHORITY {
  string id
  string name
}

AUTHORITY ||--o{ PARTICIPANT: has
PARTICIPANT }|--|{ SAGA: participant
ACCOUNT ||--|{ PARTICIPANT: join
```
