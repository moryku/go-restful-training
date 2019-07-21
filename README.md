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
| `/api/users?name={na}`      | GET     | Get {na} like in users              |

List of book routes:

| Route                       | HTTP    | Description                         |
| --------------------------- |:-------:| -----------------------------------:|
| `/api/books`                | GET     | Get all books data                  |
| `/api/books/:id`            | GET     | Get singel book by `id`             |
| `/api/books`                | POST    | Create new book                     |
| `/api/books/:id`            | DELETE  | Delete book by `id`                 |
| `/api/books/:id`            | PUT     | Update book information             |
| `/api/books?name={na}`      | GET     | Get {na} like in books              |