# Domain model

```marmeid
erDiagram
    USER {
      string id
      string name
      string type
    }
    PROJECT {
      string id
      string title
      string description
      user_id owner
    }
    EPIC {
      string id
      string title
      string description
      user_id owner
    }
    STORY {
      string id
      string title
      string description
      user_id owner
    }
    TALE {
      string id
      string title
      string description
      user_id owner
    }
    REQUIREMENT {
      string id
      string title
      string description
      user_id owner
    }
    TASK {
      string id
      string title
      string description
      status_id status
      project_id project_id
      task_id parent_id
      user_id worker_id
      user_id reporter_id
    }
    USER }|--|| AUTHORITY: has
    AUTHORITY ||--|| PROJECT: has
    USER }|--|{ PROJECT: participant
    TASK }o--|| USER: manage
    TASK }o--|| USER: work
    PROJECT ||--o{ SPRINT: has
    PROJECT ||--|| BACKLOG: has
    SPRINT ||--|| BELONG: has
    BACKLOG ||--|| BELONG: has
    BELONG |o--o{ TASK: belong
    PROJECT ||--o{ EPIC: has
    PROJECT |o--o{ STORY: has
    EPIC |o--o{ STORY: has
    STORY ||--o{ TALE: has
    STORY |o--o{ REQUIREMENT: has
    TALE ||--o{ REQUIREMENT: has
    TASK ||--o{ TASK: has
    STATUS ||--o{ TASK: has
    ATTACHMENT ||--|| TASK: has
    ATTACHMENT }o--o| EPIC: has
    ATTACHMENT }o--o| STORY: has
    ATTACHMENT }o--o| TALE: has
    ATTACHMENT }o--o| REQUIREMENT: has
```
