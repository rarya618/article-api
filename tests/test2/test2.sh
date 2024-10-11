# TEST 2: Posting three articles to the API
# Checks:
# - If the API accepts multiple standard article through POST requests
# - If the API returns correct articles through the GET requests

# Expected output for the three input files
expectedOutput1=$(cat input1.json)
expectedOutput2=$(cat input2.json)
expectedOutput3=$(cat input3.json)

# Outputs from POST requests
postOutput1=$(curl -s -d @input1.json -X POST http://localhost:8080/articles)
postOutput2=$(curl -s -d @input2.json -X POST http://localhost:8080/articles)
postOutput3=$(curl -s -d @input3.json -X POST http://localhost:8080/articles)

# Outputs from GET requests
getOutput1=$(curl -s -X GET http://localhost:8080/articles/2)
getOutput2=$(curl -s -X GET http://localhost:8080/articles/3)
getOutput3=$(curl -s -X GET http://localhost:8080/articles/4)

# Store if test was successful
testIsSuccessful=false

# If post output 1 does not match
if ! echo "$postOutput1" | jq --sort-keys . | diff - <(echo "$expectedOutput1" | jq --sort-keys .); then
    echo "POST check 1 failed: returned incorrect article"

# Post check 1 successful
else
    echo "POST check 1 passed: returned correct article"
    # If post output 2 does not match
    if ! echo "$postOutput2" | jq --sort-keys . | diff - <(echo "$expectedOutput2" | jq --sort-keys .); then
        echo "POST check 2 failed: returned incorrect article"
    
    # Post check 2 successful
    else
        echo "POST check 2 passed: returned correct article"
        # If post output 3 does not match
        if ! echo "$postOutput3" | jq --sort-keys . | diff - <(echo "$expectedOutput3" | jq --sort-keys .); then
            echo "POST check 3 failed: returned incorrect article"
        
        # Post check 3 successful
        else
            echo "POST check 3 passed: returned correct article"

            # If get output 1 does not match
            if ! echo "$getOutput1" | jq --sort-keys . | diff - <(echo "$expectedOutput1" | jq --sort-keys .); then
                echo "GET check 1 failed: returned incorrect article"
            
            # Get check 1 successful
            else
                echo "GET check 1 passed: returned correct article"

                # If get output 2 does not match
                if ! echo "$getOutput2" | jq --sort-keys . | diff - <(echo "$expectedOutput2" | jq --sort-keys .); then
                    echo "GET check 2 failed: returned incorrect article"
                
                # Get check 2 successful
                else
                    echo "GET check 2 passed: returned correct article"
                    
                    # If get output 3 does not match
                    if ! echo "$getOutput3" | jq --sort-keys . | diff - <(echo "$expectedOutput3" | jq --sort-keys .); then
                        echo "GET check 3 failed: returned incorrect article"
                    
                    # Get check 3 successful
                    else
                        echo "GET check 3 passed: returned correct article"
                        testIsSuccessful=true
                    fi
                fi
            fi
        fi
    fi
fi

# Print test outcome
if [ "$testIsSuccessful" = "true" ]; then
    echo "Test 2 successful"
else
    echo "Test 2 failed"
fi