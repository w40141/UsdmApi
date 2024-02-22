
# Shelf ER diagram

```marmeid
erDiagram

CHARACTER {
  string id
  member_id member
  ability_id ability
  saga_id saga
}
SAGA {
  string id
  string title
  string description
  character_id creator
  date created_at
  date updated_at
}
BOOK {
  string id
  saga_id saga
  string name
  string description
  character_id creator
  date created_at
  date updated_at
}
TOC {
  book_id book
  story_id story
}
STORY {
  string id
  string title
  string description
  string reason
  saga_id saga
  character_id creator
  date created_at
  date updated_at
}
PATH {
  story_id ancestor
  story_id descendant
}
SCENE {
  string id
  string title
  saga_id saga
  story_id story
  character_id creator
  date created_at
  date updated_at
}

CHARACTER }|--|| SAGA: apper
CHARACTER }|--|| BOOK: create
CHARACTER }|--|| STORY: create
CHARACTER }|--|| SCENE: create
SAGA ||--o{ BOOK: has
SAGA ||--o{ STORY: has
BOOK ||--o{ TOC: belong
TOC }o--|| STORY: belong
STORY }|--|{ PATH: ancestor
STORY }|--|{ PATH: descendant
STORY ||--o{ SCENE: specify
SAGA ||--o{ SCENE: has
```
