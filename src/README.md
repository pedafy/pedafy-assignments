# API Routes

> Here will be defined all the API's routes using the Richardson Maturity Model, more info [here](https://martinfowler.com/articles/richardsonMaturityModel.html).

____

### Assignments

| Action | Method | Route | Resp code |
| ---- | ---- | ---- | ---- |
| Get all assignments | `GET` | `/assignments` | `200`,`204`,`403`,`503` |
| Get assignments by attribut's value | `GET` | `/assignments/{filter}/{value}` | `200`,`204`,`403`,`422`,`503` |
| Get assignment by its ID | `GET` | `/assignment/{id}` | `200`,`403`,`410`,`503` |
| Create a new assignment | `PUT` | `/assignment` | `201`,`403`,`422` |
| Modify an assignment | `POST` | `/assignment/{id}` | `200`,`403`,`410`,`422` |
| Archive an assignment | `POST` | `/assignment/{id}` | `200`,`403`,`410`,`422` |

### Status
| Action | Method | Route | Resp code |
| ---- | ---- | ---- | ---- |
| Get all the available status | `GET` | `/status` | `200`,`204`,`403` |
| Get status by ID | `GET` | `/status/{id}` | `200`,`410` |
| Get status by name | `GET` | `/status/{name}` | `200`,`410` |
