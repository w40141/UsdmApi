# Task ER diagram

```marmeid
erDiagram
CHARACTER {
  account_id account
  ability_id ability
  saga_id saga
}
SCENE {
  string id
  string title
  story_id story
  character_id character
  date created_at
  date updated_at
}
PROP {
  string id
  string title
  string description
  string reason
  status_id status
  character_id creator
  date created_at
  date updated_at
}
PROPRIETOR {
  string id
  character_id user
  prop_id prop
  date start
  date end
}
SPRINT {
  string id
  date start
  date end
}
STATUS {
  string id
  string name
}
INSPRINT {
  sprint_id sprint
  prop_id prop
}
INCORPORATE {
  scene_id scene
  prop_id prop
}

PROPRIETOR }|--|| PROP: use
CHARACTER ||--|| PROP: create
CHARACTER ||--o{ SCENE: create
CHARACTER ||--o{ PROPRIETOR: use
SCENE ||--|| INCORPORATE: has
INSPRINT ||--|| PROP: has
SPRINT ||--|| INSPRINT: has
INCORPORATE ||--|| PROP: has
PROP }o--|| STATUS: has
```
