  It is a simple registration and authentication microservice with the option of adding both open and authenticated endpoints.
- Endpoints
    Open endpoints:
      1) /users - Registration of user in database/
      2) /sessions - Creating session for the user (returns session code in Coockie)
    Closed (/private prefix) endpoints:
      1) /whoami - Returns unic ID of the authorisated user. 
- Database realisation.
    Connection realised by Store and UserRepository interfaces, which allows us to connect diferent DB's. Now created postgres realisation and realisation by map (just for tests).
- Migrations
    Postgres migrations created by: https://github.com/golang-migrate/migrate
