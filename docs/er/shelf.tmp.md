
# Shelf ER diagram

```marmeid
erDiagram

MEMBER {
  account_id account
  world_id world
  authority_id authority
}
CHARACTER {
  account_id account
  ability_id ability
  saga_id saga
}
SAGA {
  string id
  string title
  string description
  member_id creator
  date created_at
  date updated_at
}
BOOK {
  string id
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
  story_id story
  character_id creator
  date created_at
  date updated_at
}

MEMBER ||--o{ SAGA: create
CHARACTER }|--|| SAGA: apper
CHARACTER }|--|| BOOK: create
CHARACTER }|--|| STORY: create
CHARACTER }|--|| SCENE: create
BOOK ||--o{ TOC: belong
TOC }o--|| STORY: belong
STORY }|--|{ PATH: ancestor
STORY }|--|{ PATH: descendant
STORY ||--o{ SCENE: specify
```
