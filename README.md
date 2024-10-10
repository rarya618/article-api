# article-api
A simple API with three endpoints.

## Setup
To set the environment up with Mac:
- clone the repository
- open the directory in Terminal
- run `./run.sh`

If you run into permission issues, run:
- `chmod +x run.sh`

## Error Handling
The API is meant to respond with JSON objects, which is why the errors are also thrown as JSON objects.

### List of errors
- Tag name not provided
- Invalid date: should have exactly 8 characters