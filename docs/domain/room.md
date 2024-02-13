# Room ER diagram

```marmeid
erDiagram
ACCOUNT {
  string id
  string name
  string email
}
MEMBER {
  string id
  account_id account
  shelf_id shelf
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
SHELF {
  string id
  saga_id saga
}
AUTHORITY {
  string id
  string name
}

ACCOUNT ||--|| SHELF: has
PARTICIPANT }|--|{ SAGA: participant
ACCOUNT }|--|| PARTICIPANT: join
AUTHORITY ||--o{ PARTICIPANT: has
```
