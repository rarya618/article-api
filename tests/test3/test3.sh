# TEST 3: Posting three articles to the API
# Checks:
# - If the API returns a valid error after the POST request
# - If the API returns the correct article through the GET request

# Expected output for the three input files
postExpectedOutput=$(cat postOutput.json)
getExpectedOutput=$(cat getOutput.json)

# Output from POST request
postOutput=$(curl -s -d @input.json -X POST http://localhost:8080/articles)

# Output from GET request
getOutput=$(curl -s -X GET http://localhost:8080/articles/1)

# Store if test was successful
testIsSuccessful=false

# If post output does not match
if ! echo "$postOutput" | jq --sort-keys . | diff - <(echo "$postExpectedOutput" | jq --sort-keys .); then
    echo "POST check failed: unexpected response"
    echo "Test 3 failed"

# Post check successful
else
    echo "POST check passed: valid error thrown"
    
    # If get output does not match
    if ! echo "$getOutput" | jq --sort-keys . | diff - <(echo "$getExpectedOutput" | jq --sort-keys .); then
        echo "GET test failed: returned incorrect article"
        echo "Test 3 failed"
    
    # Get check successful
    else
        echo "GET check passed: returned original article"
        testIsSuccessful=true
    fi
fi

# Print test outcome
if [ "$testIsSuccessful" = "true" ]; then
    echo "Test 3 successful"
else
    echo "Test 3 failed"
fi