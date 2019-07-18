# My App

Demo app with basic RESTFul API.

## REST API

List of basic routes:

| Route                       | HTTP   | Description            |
| --------------------------- |:------:| ----------------------:|
| `/api/hello?name={name}`    | GET    | Hello, `{name}`        |

List of user routes:

| Route                       | HTTP    | Description                         |
| --------------------------- |:-------:| -----------------------------------:|
| `/api/users`                | GET     | Get all users data                  |
| `/api/users/:id`            | GET     | Get singel user by `id`             |
| `/api/users`                | POST    | Create new user                     |
| `/api/users/:id`            | DELETE  | Delete user by `id`                 |
| `/api/users/:id`            | PUT     | Update user information             |
| `/api/users?name={name}`    | GET     | Get {name} match in users           |
| `/api/users?name={na}`      | GET     | Get {na} like in users              |

List of book routes:

| Route                       | HTTP    | Description                  |
| --------------------------- |:-------:| ----------------------------:|
| `/api/book`                | GET     | Get all book data            |
| `/api/book/:id`            | GET     | Get singel user by `id`      |
| `/api/book`                | POST    | Create new user              |
| `/api/book/:id`            | DELETE  | Delete user by `id`          |