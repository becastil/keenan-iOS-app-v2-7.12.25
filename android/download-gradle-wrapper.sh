#!/bin/bash

# Download gradle wrapper jar if it doesn't exist
WRAPPER_JAR="gradle/wrapper/gradle-wrapper.jar"

if [ ! -f "$WRAPPER_JAR" ]; then
    echo "Downloading gradle-wrapper.jar..."
    mkdir -p gradle/wrapper
    wget https://github.com/gradle/gradle/raw/v8.3.0/gradle/wrapper/gradle-wrapper.jar -O "$WRAPPER_JAR"
    echo "gradle-wrapper.jar downloaded successfully"
else
    echo "gradle-wrapper.jar already exists"
fi