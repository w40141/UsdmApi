# Domain model

```marmeid
erDiagram
    USER {
      string id
      string name
    }
    PROJECT {
      string id
      string title
      owner id
    }
    USER ||--|{ PROJECT: has
    %%USER ||--|{ TASK: tantou
    PROJECT ||--o{ TASK: has
    PROJECT ||--o{ EPIC: has
    PROJECT ||--o{ STORY: has
    EPIC |o--o{ STORY: has
    %%EPIC |o--o{ TASK: has
    STORY ||--o{ TALE: has
    %%STORY |o--o{ TASK: has
    STORY |o--o{ REQUIREMENT: has
    TALE ||--o{ REQUIREMENT: has
    %%TALE |o--o{ TASK: has
    %%REQUIREMENT |o--o{ TASK: has
    %%TASK ||--o{ TASK: has
    %%SPRINT |o--o{ TASK: has
    %%BACKLOG |o--o{ TASK: has
    %%TASK }o--|| STATUS: has
```
