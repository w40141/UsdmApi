# Task ER diagram

```marmeid
erDiagram
    MEMBER {
      string id
      string name
      string authority_id
    }
    PROP {
      string id
      string title
      string description
      string reason
      status_id status
      prop_id parent_id
      member_id user
      member_id creator
      legend_id legend
    }
    PARTICIPANT {
      string member_id
      string authority_id
      string legend_id
    }
    LEGEND {
      string id
      string title
      string description
    }
    SPRINT {
      string id
      legend_id legend
      date start
      date end
    }
    NARRATIVE {
      string id
      string title
      string description
      string reason
      legend_id legend
      date start
      date end
    }
    STORY {
      string id
      string title
      string description
      string reason
      legend_id legend
      parent_id parent
    }
    EPISODE {
      string id
      string title
      string description
      string reason
      legend_id legend
      story_id story
    }
    SCENE {
      string id
      string title
      string description
      string reason
      legend_id legend
      parent_id parent
    }
    INCORPORATION {
      parent_id belonging
      prop_id id
    }
    STATUS {
      string id
      string name
    }
    AUTHORITY {
      string id
      string name
    }
    MEMBER ||--|{ PARTICIPANT: has
    PARTICIPANT }|--|{ LEGEND: participant
    AUTHORITY ||--o{ PARTICIPANT: has
    LEGEND ||--|{ NARRATIVE: belong
    LEGEND ||--|{ STORY: belong
    LEGEND ||--|{ EPISODE: belong
    LEGEND ||--|{ SCENE: belong
    SPRINT }o--|| LEGEND: has
    PROP }o--o| SPRINT: has
    PROP }o--|| LEGEND: has
    LEGEND ||--o{ NARRATIVE: has
    LEGEND |o--o{ STORY: has
    NARRATIVE |o--o{ STORY: has
    STORY ||--o{ EPISODE: has
    STORY |o--o{ SCENE: has
    EPISODE ||--o{ SCENE: has
    MEMBER ||--o{ PROP: create
    MEMBER ||--o{ PROP: use
    PROP ||--o{ PROP: has
    STATUS ||--o{ PROP: has
    PROP ||--|| INCORPORATION: has
    INCORPORATION }o--o| NARRATIVE: has
    INCORPORATION }o--o| STORY: has
    INCORPORATION }o--o| EPISODE: has
    INCORPORATION }o--o| SCENE: has
```
