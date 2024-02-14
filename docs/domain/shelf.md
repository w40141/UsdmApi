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

SAGA ||--o{ BOOK: belong
SAGA |o--o{ STORY: has
STORY }o--|| SAGA: belong
STORY }o--o| BOOK: has
EPISODE }o--|| SAGA: belong
STORY ||--o{ EPISODE: has
SAGA ||--o{ SCENE: belong
BOOK |o--o{ SCENE: has
STORY |o--o{ SCENE: has
EPISODE |o--o{ SCENE: has
```
