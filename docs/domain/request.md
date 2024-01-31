# Request ER diagram

```marmeid
erDiagram
    MEMBER {
      string id
      string name
      string authority_id
    }
    PARTICIPANT {
      member_id member
      authority_id authority
      legend_id legend
    }
    LEGEND {
      string id
      string title
      string description
      date created_at
      date updated_at
    }
    NARRATIVE {
      string id
      string title
      string description
      string reason
      legend_id legend
      date created_at
      date updated_at
    }
    STORY {
      string id
      string title
      string description
      string reason
      legend_id legend
      parent_id parent
      date created_at
      date updated_at
    }
    EPISODE {
      string id
      string title
      string description
      string reason
      legend_id legend
      story_id story
      date created_at
      date updated_at
    }
    SCENE {
      string id
      string title
      string description
      string reason
      legend_id legend
      parent_id parent
      date created_at
      date updated_at
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
    LEGEND ||--o{ NARRATIVE: has
    LEGEND |o--o{ STORY: has
    NARRATIVE |o--o{ STORY: has
    EPISODE }o--|| STORY: has
    STORY |o--o{ SCENE: has
    EPISODE ||--o{ SCENE: has
```
