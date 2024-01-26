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
    OWNER }o--|{ USER: has
    USER }|--o{ WORKER: has
    USER }|--o{ REPORTER: has
    OWNER }|--|{ PROJECT: has
    REPORTER ||--o{ TASK: manage
    WORKER ||--o{ TASK: work
    TASK }o--|| PROJECT: has
    PROJECT ||--o{ EPIC: has
    PROJECT |o--o{ STORY: has
    EPIC |o--o{ STORY: has
    TASK }o--o| EPIC: has
    STORY ||--o{ TALE: has
    TASK }o--o| STORY: has
    STORY |o--o{ REQUIREMENT: has
    TALE ||--o{ REQUIREMENT: has
    TASK }o--o| TALE: has
    TASK }o--o| REQUIREMENT: has
    TASK ||--o{ TASK: has
    SPRINT |o--o{ TASK: has
    TASK }o--o| BACKLOG: has
    STATUS ||--o{ TASK: has
    BACKLOG ||--|| PROJECT: has
```
