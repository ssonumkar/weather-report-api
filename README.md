Database setup:
1.db used: MySQL
2.db name: weather_report
3.db-setup.sql has the script to create all the necessary tables and stored procedures.

Api:
1)All the configurations are fetched from config package
2)run the api using: **go run main.go**
3)All endpoints can be fetched from routes package.


**Architechure:**
1. All the apis i.e Weather, User Auth, Weather history are been implemented in kind of MVC layered architechtural pattern. So, all of these apis have packages and every package will have contoller, service, repository and model files.
3. For every user, when logged in is issued a jwt token which can be used futher for authentication until user is logged out. These tokens are stateless and not persisted on system so user will need to again login if server is re-started.
4. All the passwords are encrypted and then stored in the database. 
5. All the configurations are fetched from config package.
6. Every package has a test folder which will have all the respective test files of the respective packages.
7. The entry point for application is main.go at base path.
8. Packages:
   **middleware**: Middleware package is a wrapper over each api to validate the auth token. So, every api call first goes through this middlewarwe and a user is     validated.
   **routes**: All the api routes of api are configured in routes package.
   **config**:  All the configurations are fetched from config package.
   **weather**: Has all the logic implemented for weather api. It has controller -> handler function for api endpoint and then has service and repository level go files.
   **weatherhistory**: Has all the logic implemented for weather search history api. It has controller -> handler function for all the api endpoints and then has service and repository level go files.
   **auth**: Has all the logic implemented for user authentication and management. It has controller -> handler function for auth api endpoints and then has service and repository level go files.
   **log**: Custom logging framework is implemented.
   **encrypt**: this package is used for password encryption and managing the user auth tokens.

**Note**: DOB for user register api accepts date in YYYY-MM-DD format
