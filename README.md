# article-api
A simple API with three endpoints.

## Setup
To set the environment up with Mac:
- clone the repository
- open the directory in Terminal
- run `./run.sh`

If you run into permission issues, run:
- `chmod +x run.sh`

## Description
The main language used to build this API is Golang. It was chosen because of its simplicity and it would be easier to extend and maintain in the future.

Looking at efficiency, the code runs in O(n^2) time.

### Endpoints
- POST `/articles` handles the receipt of some article data in json format, and store it within the service.
- GET `/articles/{id}` returns the JSON representation of the article.
- GET `/tags/{tagName}/{date}` returns the list of articles that have that tag name on the given date and some summary data about that tag for that day.

## Assumptions
- The API only accepts and responds with JSON objects.
- The API requests received will be in continuous (no breaks) and increasing order of id i.e. 1, 2, 3,...
- The articles are not subject to deletion.
- Each POST request only contains one article.
- The articles are posted in order of chronological order of dates (earliest comes first).
- Every article posted is of the correct format, containing all fields in the correct data type.
- Tag name has been assumed to exist for all cases and will search for it when provided.

## Libraries
- fmt
- strconv
- net/http
- github.com/gin-gonic/gin

Personally I chose Gin instead of Mux solely because of more experience with Gin. 

## Error Handling
The API is meant to respond with JSON objects, so the errors are also thrown as JSON objects.

### List of errors
- Invalid tag: Tag name not provided (Error 400)
- Invalid date: should have exactly 8 characters (Error 400)
- Invalid date: should be a valid number (Error 400)
- Invalid date: should be greater than 0 (Error 400)
- Invalid date: should be 8 digits (Error 400)
- Invalid date: year invalid (Error 400)
- Invalid date: month invalid (Error 400)
- Invalid date: day invalid (Error 400)
- Invalid Article ID: should have a unique Article ID (Error 400)
- Invalid Article ID: Article ID needs to be a number (Error 400)
- Article not found (Error 404)

## Tests
Each test will run a set of checks and if all the checks pass, the test is successful.

NOTE: The server needs to be running independently and has to be restarted before every test.

### Test 1: Posting one article to the API
#### Checks
- If the API accepts a standard article through the POST request
- If the API returns the same article through the GET request

### Test 2: Posting three articles to the API
#### Checks
- If the API accepts multiple standard article through POST requests
- If the API returns correct articles through the GET requests

### Test 3: Posting article with existing ID
#### Checks
- If the API returns a valid error after the POST request
- If the API returns the correct article through the GET request

### Test 4: Posting one article to the API and getting one tag
#### Checks
- If the API accepts a standard article through the POST request
- If the API returns the same article through the GET request
- If the API returns the correct tag data through the GET request

### Test 5: Posting multiple articles on the same date to the API and getting two tags
#### Checks:
- If the API accepts multiple standard article through POST requests
- If the API returns correct articles through the GET requests
- If the API returns the correct tag data through the GET request