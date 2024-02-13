# Shelf ER diagram

```marmeid
erDiagram

SAGA {
  string id
  string title
  string description
  date created_at
  date updated_at
}
BOOK {
  string id
  string title
  string description
  string reason
  saga_id saga
  date created_at
  date updated_at
}
STORY {
  string id
  string title
  string description
  string reason
  saga_id saga
  parent_id parent
  date created_at
  date updated_at
}
EPISODE {
  string id
  string title
  string description
  string reason
  saga_id saga
  story_id story
  date created_at
  date updated_at
}
SCENE {
  string id
  string title
  string description
  string reason
  saga_id saga
  parent_id parent
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
SAGA ||--o{ BOOK: belong
SAGA ||--o{ STORY: belong
SAGA |o--o{ STORY: has
BOOK |o--o{ STORY: has
SAGA ||--o{ EPISODE: belong
STORY ||--o{ EPISODE: has
SAGA ||--o{ SCENE: belong
BOOK |o--o{ SCENE: has
STORY |o--o{ SCENE: has
EPISODE |o--o{ SCENE: has
```
