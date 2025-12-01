#!/bin/bash

# ==========================================
# PROJECT MANAGEMENT SCRIPT v1.2
# ==========================================
# Usage:
#   ./manage.sh init 02 prime_checker
#   ./manage.sh deploy "feat: solved day 2"
#   ./manage.sh sync 02 "Prime Numbers"
# ==========================================

USER_PREFIX="10xdev4u"
UPSTREAM_REMOTE="upstrm"
PRIMARY_REMOTE="github"
BACKUP_REMOTE="codeberg"
PRIMARY_BRANCH="main"
SUBMISSION_BRANCH="submission"

# ANSI Colors
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

    local FOLDER="December - $DAY"
    if [ ! -d "$FOLDER" ]; then
        echo -e "${LOG_ERROR} Directory '$FOLDER' not found."
        exit 1
    fi
    local FILE_PATH="${FOLDER}/${USER_PREFIX}_${TASK_NAME}.go"
    if [ -f "$FILE_PATH" ]; then
        echo -e "${LOG_ERROR} File '$FILE_PATH' already exists."
        exit 1
    fi
    echo -e "${LOG_INFO} Generating boilerplate..."
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
# Function: Deploy (Commit & Push to Lab)
# ------------------------------------------
function deploy() {
    local MSG=$1
    if [ -z "$MSG" ]; then MSG="chore: update daily progress"; fi

    echo -e "${LOG_INFO} Staging changes on MAIN..."
    git add .
    git commit -m "$MSG"

    if git remote | grep -q "$PRIMARY_REMOTE"; then
        echo -e "${LOG_INFO} Pushing to Primary ($PRIMARY_REMOTE)..."
        git push $PRIMARY_REMOTE $PRIMARY_BRANCH
    fi
    if git remote | grep -q "$BACKUP_REMOTE"; then
        echo -e "${LOG_INFO} Pushing to Backup ($BACKUP_REMOTE)..."
        git push $BACKUP_REMOTE $PRIMARY_BRANCH
    fi
    echo -e "${LOG_SUCCESS} Main branch deployed."
}

# ------------------------------------------
# Function: Sync to Submission (Sniper Mode)
# ------------------------------------------
function sync_submission() {
    local DAY=$1
    local TASK_TITLE=$2  # Capture the formal name (e.g., "Perfect Squares")

    if [ -z "$DAY" ]; then
        echo -e "${LOG_ERROR} Usage: ./manage.sh sync <DayNumber> \"Problem Title\""
        exit 1
    fi

    FOLDER="December - $DAY"
    echo -e "${LOG_INFO} 🎯 Preparing Sniper Sync for '$FOLDER'..."

    # 1. Check for dirty state
    if [ -n "$(git status --porcelain)" ]; then 
        echo -e "${LOG_ERROR} Working directory dirty. Run ./manage.sh deploy first."
        exit 1
    fi

    # 2. Switch to submission branch
    echo -e "${LOG_INFO} Switching to $SUBMISSION_BRANCH..."
    git checkout $SUBMISSION_BRANCH
    git pull $PRIMARY_REMOTE $SUBMISSION_BRANCH

    # 3. Grab EVERYTHING from main
    git checkout $PRIMARY_BRANCH -- "$FOLDER"

    # 4. PURGE THE EXTRAS
    echo -e "${LOG_INFO} 🧹 Cleaning up non-code files..."
    find "$FOLDER" -name "*_test.go" -delete
    find "$FOLDER" -name "README.md" -delete

    # 5. Commit and Push to 'github'
    if [ -n "$(git status --porcelain)" ]; then
        git add "$FOLDER"
        
        # FORMAT THE COMMIT MESSAGE
        if [ -z "$TASK_TITLE" ]; then
            COMMIT_MSG="Day $DAY: Solution (Code Only)"
        else
            COMMIT_MSG="Day $DAY: $TASK_TITLE (Code Only)"
        fi

        git commit -m "$COMMIT_MSG"
        
        echo -e "${LOG_INFO} Pushing to $PRIMARY_REMOTE/$SUBMISSION_BRANCH..."
        git push $PRIMARY_REMOTE $SUBMISSION_BRANCH
        
        echo -e "${LOG_SUCCESS} Synced Day $DAY to submission."
    else
        echo -e "${LOG_INFO} No changes detected."
    fi

    # 6. Return Home
    echo -e "${LOG_INFO} Returning to $PRIMARY_BRANCH..."
    git checkout $PRIMARY_BRANCH
}

case "$1" in
    init)   init_task "$2" "$3" ;;
    deploy) deploy "$2" ;;
    sync)   sync_submission "$2" "$3" ;;
    *)
        echo "Available Commands:"
        echo "  init <Day> <Name>         : Start a new day"
        echo "  deploy \"Message\"          : Save work to Main"
        echo "  sync <Day> \"Problem Name\" : Push Clean Code to Submission"
        ;;
esac
