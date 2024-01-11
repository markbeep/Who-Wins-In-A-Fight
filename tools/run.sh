#!/bin/sh

# Rebuilds the application everytime a change is made
# to a file

last_run_time=0
process_id=0

# Loads the environment variables
set -o allexport
source ./.env
set +o allexport

while sleep 1; do
    # Find the most recently updated files
    most_recent_file=$(find . -type f \( -name '*.go' -o -name "*.templ" \) ! -name '*_templ.go' -print 2>/dev/null -printf '%T@ %p\n' | sort -n | tail -1 | cut -f2- -d" ")
    most_recent_time=$(stat -c %Y "$most_recent_file")
    if [ "$most_recent_time" -gt "$last_run_time" ]; then
        last_run_time=$most_recent_time
        
        # Terminate the previous process if it exists
        if [ $process_id -ne 0 ]; then
            (kill -9 $(lsof -t -i:3000 -sTCP:LISTEN))
            echo "Restarting build"
        fi
        
        # Build and execute the binary in the background
        tailwindcss -i static/tw.css -o static/main.css
        templ generate
        go run main.go &
        
        # Store the process ID
        process_id=$!
    fi
done
