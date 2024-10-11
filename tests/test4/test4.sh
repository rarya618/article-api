# TEST 4: Posting one article to the API and getting one tag
# Checks:
# - If the API accepts a standard article through the POST request
# - If the API returns the same article through the GET request
# - If the API returns the correct tag data through the GET request


# Expected output for both checks
expectedOutput=$(cat input.json)
expectedTagOutput=$(cat output.json)

# Actual outputs received
postOutput=$(curl -s -d @input.json -X POST http://localhost:8080/articles)
getOutput=$(curl -s -X GET http://localhost:8080/articles/2)
tagOutput=$(curl -s -X GET http://localhost:8080/tags/fitness/20181022)

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
        
        # Tag output does not match
        if ! echo "$tagOutput" | jq --sort-keys . | diff - <(echo "$expectedTagOutput" | jq --sort-keys .); then
            echo "Tag check failed: returned incorrect Tag object"
        else
            echo "Tag check passed: returned correct Tag object"
            testIsSuccessful=true
        fi
    fi
fi

# Print test outcome
if [ "$testIsSuccessful" = "true" ]; then
    echo "Test 4 successful"
else
    echo "Test 4 failed"
fi