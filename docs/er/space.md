# SPACE ER diagram

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
  member_id creator
  date created_at
  date updated_at
}
CHARACTER {
  account_id account
  ability_id ability
  saga_id saga
}
ABILITY {
  string id
  string name
}
WORLD {
  string id
  string name
  account_id creator
}
MEMBER {
  account_id account
  world_id world
  authority_id authority
}
TRADITION {
  world_id world
  saga_id saga
}
AUTHORITY {
  string id
  string name
}

ABILITY ||--o{ CHARACTER: has
CHARACTER }|--|| SAGA: apper
ACCOUNT ||--|{ WORLD: create
WORLD ||--o{ TRADITION: has
TRADITION ||--|| SAGA: has
AUTHORITY ||--o{ MEMBER: has
ACCOUNT ||--|{ MEMBER: has
WORLD ||--|{ MEMBER: exist
MEMBER ||--o{ SAGA: create
ACCOUNT ||--o{ CHARACTER: has
```
