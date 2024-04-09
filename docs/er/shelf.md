# SHELF ER diagram

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
  account_id owner
  date created_at
  date updated_at
}
CHARACTER {
  account_id account
  ablity_id ablity
  book_id appearance
}
ABLITY {
  string id
  string name
}
BOOK {
  string id
  string name
  string description
  character_id creator
  date created_at
  date updated_at
}
COLLECTION {
  saga_id SAGA
  book_id BOOK
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

SAGA ||--o{ COLLECTION: has
COLLECTION |o--|| BOOK: has
CHARACTER }|--|| SAGA: create
ABLITY ||--o{ CHARACTER: has
ACCOUNT ||--|{ CHARACTER: has
CHARACTER }|--|| BOOK: create
CHARACTER }|--|| STORY: create
CHARACTER }|--|| SCENE: create
BOOK ||--o{ TOC: belong
TOC }o--|| STORY: belong
STORY }|--|{ PATH: ancestor
STORY }|--|{ PATH: descendant
STORY ||--o{ SCENE: specify
```