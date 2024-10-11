# TEST 5: Posting multiple articles on the same date to the API and getting two tags
# Checks:
# - If the API accepts multiple standard article through POST requests
# - If the API returns correct articles through the GET requests
# - If the API returns correct tag data through the GET requests


# Expected output for both checks
expectedOutput1=$(cat input1.json)
expectedOutput2=$(cat input2.json)
expectedTagOutput1=$(cat output1.json)
expectedTagOutput2=$(cat output2.json)

# Actual outputs received
postOutput1=$(curl -s -d @input1.json -X POST http://localhost:8080/articles)
postOutput2=$(curl -s -d @input2.json -X POST http://localhost:8080/articles)
getOutput1=$(curl -s -X GET http://localhost:8080/articles/2)
getOutput2=$(curl -s -X GET http://localhost:8080/articles/3)
tagOutput1=$(curl -s -X GET http://localhost:8080/tags/fitness/20181022)
tagOutput2=$(curl -s -X GET http://localhost:8080/tags/wellness/20181022)

# Store if test was successful
testIsSuccessful=false

# If post output 1 does not match
if ! echo "$postOutput1" | jq --sort-keys . | diff - <(echo "$expectedOutput1" | jq --sort-keys .); then
    echo "POST check 1 failed: returned incorrect article"

# Post check successful
else
    echo "POST check 1 passed: returned correct article"

    # If post output 2 does not match
    if ! echo "$postOutput2" | jq --sort-keys . | diff - <(echo "$expectedOutput2" | jq --sort-keys .); then
        echo "POST check 2 failed: returned incorrect article"

    # Post check successful
    else
        echo "POST check 2 passed: returned correct article"
    
        # If get output 1 does not match
        if ! echo "$getOutput1" | jq --sort-keys . | diff - <(echo "$expectedOutput1" | jq --sort-keys .); then
            echo "GET check 1 failed: returned incorrect article"
        
        # Both checks successful
        else
            echo "GET check 1 passed: returned correct article"

            # If get output 2 does not match
            if ! echo "$getOutput2" | jq --sort-keys . | diff - <(echo "$expectedOutput2" | jq --sort-keys .); then
                echo "GET check 2 failed: returned incorrect article"
            
            # Both checks successful
            else
                echo "GET check 2 passed: returned correct article"
            
                # Tag output 1 does not match
                if ! echo "$tagOutput1" | jq --sort-keys . | diff - <(echo "$expectedTagOutput1" | jq --sort-keys .); then
                    echo "Tag check 1 failed: returned incorrect Tag object"
                else
                    echo "Tag check 1 passed: returned correct Tag object"
                    # Tag output 1 does not match
                    if ! echo "$tagOutput2" | jq --sort-keys . | diff - <(echo "$expectedTagOutput2" | jq --sort-keys .); then
                        echo "Tag check 2 failed: returned incorrect Tag object"
                    else
                        echo "Tag check 2 passed: returned correct Tag object"
                        testIsSuccessful=true
                    fi
                fi
            fi
        fi
    fi
fi

# Print test outcome
if [ "$testIsSuccessful" = "true" ]; then
    echo "Test 5 successful"
else
    echo "Test 5 failed"
fi