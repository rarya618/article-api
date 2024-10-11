postOutput=$(curl -s -d @input.json -X POST http://localhost:8080/articles)
postExpectedOutput=$(cat postOutput.json)

getOutput=$(curl -s -X GET http://localhost:8080/articles/1)
getExpectedOutput=$(cat getOutput.json)

# Compare the outputs
if ! echo "$postOutput" | jq --sort-keys . | diff - <(echo "$postExpectedOutput" | jq --sort-keys .); then
    echo "POST check failed: unexpected response"
    echo "Test 3 failed"
else
    echo "POST check passed: valid error thrown"
    if ! echo "$getOutput" | jq --sort-keys . | diff - <(echo "$getExpectedOutput" | jq --sort-keys .); then
        echo "GET test failed: returned incorrect article"
        echo "Test 3 failed"
    else
        echo "GET check passed: returned original article"
        echo "Test 3 successful"
    fi
fi