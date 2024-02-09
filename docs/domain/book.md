# Request ER diagram

```marmeid
erDiagram
    PARTICIPANT {
      member_id member
      authority_id authority
      book_id BOOK
    }
    BOOK {
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
      book_id BOOK
      date created_at
      date updated_at
    }
    STORY {
      string id
      string title
      string description
      string reason
      book_id BOOK
      parent_id parent
      date created_at
      date updated_at
    }
    EPISODE {
      string id
      string title
      string description
      string reason
      book_id BOOK
      story_id story
      date created_at
      date updated_at
    }
    SCENE {
      string id
      string title
      string description
      string reason
      book_id BOOK
      parent_id parent
      date created_at
      date updated_at
    }
    AUTHORITY {
      string id
      string name
    }
    PARTICIPANT }|--|{ BOOK: participant
    AUTHORITY ||--o{ PARTICIPANT: has
    BOOK ||--o{ NARRATIVE: belong
    BOOK ||--o{ STORY: belong
    BOOK ||--o{ EPISODE: belong
    BOOK ||--o{ SCENE: belong
    NARRATIVE |o--o{ STORY: has
    NARRATIVE |o--o{ SCENE: has
    STORY }o--o| BOOK: has
    STORY ||--o{ EPISODE: has
    STORY |o--o{ SCENE: has
    EPISODE |o--o{ SCENE: has
    SCENE }o--o| BOOK: has
```
