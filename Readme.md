# Test1:

## This service provides two endpoints. 


Endpoint 1: ```http://localhost:8080/users ``` returns in JSON format. E.g., 
```
[ { "id": 1, "name":"John", "age":31, "city":"New York" }, 
  { "id": 2, "name":"Doe", "age":22, "city":"Vancouver" } ]
```

Endpoint 2 : ```http://localhost:8080/user/2``` returns a user with id=2 in JSON format. E.g., 
```
{ "id": 2, "name":"Doe", "age":22, "city":"Vancouver" }
```