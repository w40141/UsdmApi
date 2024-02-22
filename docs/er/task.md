# Task ER diagram

```marmeid
erDiagram
SCENE {
  string id
  string title
  saga_id saga
  story_id story
  character_id creator
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
  saga_id saga
  date created_at
  date updated_at
}
WORKER {
  prop_id prop
  character_id worker
  date start
  date end
}
SPRINT {
  string id
  saga_id saga
  date start
  date end
}
STATUS {
  string id
  string name
}
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
INSPRINT {
  sprint_id sprint
  prop_id prop
}
INCORPORATE {
  scene_id scene
  prop_id prop
}

CHARACTER }|--|| SAGA: apper
CHARACTER }|--|| SCENE: create
CHARACTER }|--|| PROP: create
CHARACTER }|--|| WORKER: has
WORKER }|--|| PROP: work
SAGA ||--o{ SPRINT: has
SAGA ||--|| PROP: has
SAGA ||--o{ SCENE: has
SCENE ||--|| INCORPORATE: has
INSPRINT ||--|| PROP: has
SPRINT ||--|| INSPRINT: has
INCORPORATE ||--|| PROP: has
PROP }o--|| STATUS: has
```
