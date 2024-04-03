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
  ablity_id ablity
  world_id world
}
WORLD {
  string id
  string name
  account_id creator
}
EXIST {
  world_id world
  book_id book
}
TRADITION {
  world_id world
  saga_id saga
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

ACCOUNT ||--|{ WORLD: create
WORLD ||--o{ EXIST: has
EXIST ||--|| BOOK: has
SAGA ||-- |{ CHARACTER: create
WORLD ||--o{ TRADITION: has
TRADITION ||--|| SAGA: has
ABLITY ||--o{ CHARACTER: has
ACCOUNT ||--|{ CHARACTER: has
WORLD ||--|{ CHARACTER: exist
CHARACTER }|--|| BOOK: create
CHARACTER }|--|| STORY: create
CHARACTER }|--|| SCENE: create
BOOK ||--o{ TOC: belong
TOC }o--|| STORY: belong
STORY }|--|{ PATH: ancestor
STORY }|--|{ PATH: descendant
STORY ||--o{ SCENE: specify
SAGA ||--o{ BOOK: collection
```
