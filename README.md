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
The main language used to build this API is Golang. It was chosen because of its simplicity. It would be easier to extend and maintain in the future.

## Endpoints
- POST `/articles` handles the receipt of some article data in json format, and store it within the service.
- GET `/articles/{id}` returns the JSON representation of the article.
- GET `/tags/{tagName}/{date}` returns the list of articles that have that tag name on the given date and some summary data about that tag for that day.

## Assumptions
- The API only accepts and responds with JSON objects
- Efficiency was taken into consideration, so the code runs in O(n^2) time


## Error Handling
The API is meant to respond with JSON objects, which is why the errors are also thrown as JSON objects.

### List of errors
- Tag name not provided
- Invalid date: should have exactly 8 characters

## Tests
### Test 1: Standard POST requests
- To test if the API accepts a standard JSON POST request
