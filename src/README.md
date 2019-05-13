# API Routes

> Here will be defined all the API's routes using the Richardson Maturity Model, more info [here](https://martinfowler.com/articles/richardsonMaturityModel.html).

____

### Assignments

| Action | Method | Route | Resp code |
| ---- | ---- | ---- | ---- |
| Get all assignments | `GET` | `/assignments` | `200`, `204`, `403`, `503` |
| Get all assignments with order | `GET` | `/assignments?sort=status|new|due` | `200`, `204`, `403`, `503` |
| Get assignments by attribut's value | `GET` | `/assignments/{id_type}/{id:[0-9]+}` | `200`, `204`, `403`, `503` |
| Create a new assignment | `PUT` | `/assignment` | `201`, `403`, `422`, `503` |
| Modify an assignment | `POST` | `/assignment/{id:[0-9]+}` | `200`,`403`, `410`, `422`, `503` |
| Archive an assignment | `POST` | `/assignment/archive/{id:[0-9]+}` | `200`, `403`, `410`, `503` |

### Status
| Action | Method | Route | Resp code |
| ---- | ---- | ---- | ---- |
| Get all the available status | `GET` | `/status` | `200`, `204`, `403`, `503` |
| Get status by ID | `GET` | `/status/{id:[0-9]+}` | `200`, `403`, `410`, `503` |
| Get status by name | `GET` | `/status/{name}` | `200`, `403`, `410`, `503` |
