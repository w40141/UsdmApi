
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
  saga_id saga
  date created_at
  date updated_at
}
BELONG {
  book_id book
  story_id story
}
STORY {
  string id
  string title
  string description
  string reason
  saga_id saga
  date created_at
  date updated_at
}
PATH {
  string ancestor
  string descendant
}
SCENE {
  string id
  string title
  saga_id saga
  story_id story
  date created_at
  date updated_at
}

SAGA ||--o{ BOOK: has
SAGA ||--o{ STORY: has
SAGA ||--o{ SCENE: has
BOOK |o--o{ BELONG: belong
BELONG }o--o| STORY: belong
STORY }|--o{ PATH: ancestor
STORY }|--o{ PATH: descendant
STORY ||--o{ SCENE: specify
```
