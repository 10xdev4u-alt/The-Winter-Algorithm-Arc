#!/bin/bash

# ==========================================
# PROJECT MANAGEMENT SCRIPT
# ==========================================
# Remotes: github, codeberg, upstrm
# ==========================================

USER_PREFIX="10xdev4u"
UPSTREAM_REMOTE="upstrm"
PRIMARY_REMOTE="github"
BACKUP_REMOTE="codeberg"
PRIMARY_BRANCH="main"

# ANSI Colors for Logging
LOG_INFO='\033[0;36m[INFO]\033[0m'
LOG_SUCCESS='\033[0;32m[SUCCESS]\033[0m'
LOG_ERROR='\033[0;31m[ERROR]\033[0m'

# ------------------------------------------
# Function: Initialize New Task
# ------------------------------------------
function init_task() {
    local DAY=$1
    local TASK_NAME=$2

    if [ -z "$DAY" ] || [ -z "$TASK_NAME" ]; then
        echo -e "${LOG_ERROR} Usage: ./manage.sh init <DayNumber> <TaskName>"
        exit 1
    fi

    echo -e "${LOG_INFO} Syncing with upstream repository..."
    git pull $UPSTREAM_REMOTE $PRIMARY_BRANCH

    # Format Folder Name (e.g., "December - 01")
    local FOLDER="December - $DAY"
    
    if [ ! -d "$FOLDER" ]; then
        echo -e "${LOG_ERROR} Directory '$FOLDER' not found. Verify the Day Number."
        exit 1
    fi

    local FILE_PATH="${FOLDER}/${USER_PREFIX}_${TASK_NAME}.go"

    if [ -f "$FILE_PATH" ]; then
        echo -e "${LOG_ERROR} File '$FILE_PATH' already exists."
        exit 1
    fi

    # Generate Go Boilerplate
    echo -e "${LOG_INFO} Generating boilerplate for $FILE_PATH..."
    cat <<EOF > "$FILE_PATH"
package main

import (
	"fmt"
)

func main() {
	// Solution for Day $DAY: $TASK_NAME
	fmt.Println("Executing solution for $TASK_NAME...")
}
EOF

    echo -e "${LOG_SUCCESS} Task initialized. Opening editor..."
    code "$FILE_PATH"
}

# ------------------------------------------
# Function: Deploy (Commit & Push)
# ------------------------------------------
function deploy() {
    local MSG=$1
    if [ -z "$MSG" ]; then
        MSG="chore: update daily progress"
    fi

    echo -e "${LOG_INFO} Staging changes..."
    git add .

    echo -e "${LOG_INFO} Committing with message: '$MSG'..."
    git commit -m "$MSG"

    # Push to Primary (GitHub)
    if git remote | grep -q "$PRIMARY_REMOTE"; then
        echo -e "${LOG_INFO} Pushing to Primary ($PRIMARY_REMOTE)..."
        git push $PRIMARY_REMOTE $PRIMARY_BRANCH
    else
        echo -e "${LOG_ERROR} Remote '$PRIMARY_REMOTE' not found!"
    fi

    # Push to Backup (Codeberg)
    if git remote | grep -q "$BACKUP_REMOTE"; then
        echo -e "${LOG_INFO} Pushing to Backup ($BACKUP_REMOTE)..."
        git push $BACKUP_REMOTE $PRIMARY_BRANCH
    fi

    echo -e "${LOG_SUCCESS} Deployment complete."
}

# ------------------------------------------
# Main Execution
# ------------------------------------------
case "$1" in
    init)
        init_task "$2" "$3"
        ;;
    deploy)
        deploy "$2"
        ;;
    *)
        echo "Available Commands:"
        echo "  init <Day> <Name>   : Pull upstream and create new Go file"
        echo "  deploy \"Message\"    : Commit and push to github & codeberg"
        ;;
esac
