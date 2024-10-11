# TEST 1: Posting one article to the API
# Checks:
# - If the API accepts a standard article through the POST request
# - If the API returns the same article through the GET request


# Expected output for both checks
expectedOutput=$(cat input.json)

# The actual outputs received
postOutput=$(curl -s -d @input.json -X POST http://localhost:8080/articles)
getOutput=$(curl -s -X GET http://localhost:8080/articles/2)

# Store if test was successful
testIsSuccessful=false

# If post output does not match
if ! echo "$postOutput" | jq --sort-keys . | diff - <(echo "$expectedOutput" | jq --sort-keys .); then
    echo "POST check failed: returned incorrect article"

# Post check successful
else
    echo "POST check passed: returned correct article"
    
    # If get output does not match
    if ! echo "$getOutput" | jq --sort-keys . | diff - <(echo "$expectedOutput" | jq --sort-keys .); then
        echo "GET check failed: returned incorrect article"
    
    # Both checks successful
    else
        echo "GET check passed: returned correct article"
        testIsSuccessful=true
    fi
fi

# Print test outcome
if [ "$testIsSuccessful" = "true" ]; then
    echo "Test 1 successful"
else
    echo "Test 1 failed"
fi