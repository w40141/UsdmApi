# Task ER diagram

```marmeid
erDiagram
    USER {
      string id
      string name
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
      id parent_id
    }
    STORY {
      string id
      string title
      string description
      user_id owner
      id parent_id
    }
    TALE {
      string id
      string title
      string description
      user_id owner
      id parent_id
    }
    REQUIREMENT {
      string id
      string title
      string description
      user_id owner
      id parent_id
    }
    TASK {
      string id
      string title
      string description
      status_id status
      task_id parent_id
      user_id worker_id
      user_id reporter_id
      project_id project_id
    }
    ATTACHMENT {
      string belonging_id
      task_id id
    }
    STATUS {
      string id
      string name
    }
    AUTHORITY {
      string id
      string name
    }
    PARTICIPANT {
      string user_id
      string name
      string project_id
    }
    BELONG {
      string task_id
      string backlog_id
      string sprint_id
    }
    SPRINT {
      string id
      string project_id
      date start
      date end
    }
    USER ||--|{ PARTICIPANT: has
    TASK }o--o| BELONG: belong
    PARTICIPANT }|--|{ PROJECT: participant
    AUTHORITY ||--o{ PARTICIPANT: has
    SPRINT }o--|| PROJECT: has
    BELONG }o--o{ SPRINT: has
    TASK }o--|| PROJECT: has
    PROJECT ||--o{ EPIC: has
    PROJECT |o--o{ STORY: has
    EPIC |o--o{ STORY: has
    STORY ||--o{ TALE: has
    STORY |o--o{ REQUIREMENT: has
    TALE ||--o{ REQUIREMENT: has
    USER ||--o{ TASK: manage
    USER ||--o{ TASK: work
    TASK ||--o{ TASK: has
    STATUS ||--o{ TASK: has
    TASK ||--|| ATTACHMENT: has
    ATTACHMENT }o--o| EPIC: has
    ATTACHMENT }o--o| STORY: has
    ATTACHMENT }o--o| TALE: has
    ATTACHMENT }o--o| REQUIREMENT: has
```
