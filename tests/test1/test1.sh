postOutput=$(curl -s -d @input.json -X POST http://localhost:8080/articles)
postExpectedOutput=$(cat input.json)

getOutput=$(curl -s -X GET http://localhost:8080/articles/2)
getExpectedOutput=$(cat input.json)

# Compare the outputs
if ! echo "$postOutput" | jq --sort-keys . | diff - <(echo "$postExpectedOutput" | jq --sort-keys .); then
    echo "POST check failed"
    echo "Test 1 failed"
else
    echo "POST check passed"
    if ! echo "$getOutput" | jq --sort-keys . | diff - <(echo "$getExpectedOutput" | jq --sort-keys .); then
        echo "GET test failed"
        echo "Test 1 failed"
    else
        echo "GET check passed"
        echo "Test 1 successful"
    fi
fi