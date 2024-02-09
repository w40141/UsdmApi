# User ER diagram

```marmeid
erDiagram
  ACCOUNT {
    string id
    string name
    string email
    string display_name
  }
  COMPANY {
    string id
    string name
    string email
    string display_name
    enterprise_plan_id enterprise_plan
  }
  ENTERPRISE_PLAN {
    string id
    string name
  }
  SHELF {
    string id
    parent_id owner
  }
  BOOK {

  }
  PARTICIPANT {
    string id
    member_id member
    authority_id authority
  }
```
