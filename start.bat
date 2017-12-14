@echo off

title BoomCase - Google App Engine

echo ========== ========== ========== ========== ==========
echo Starting BoomCase on Google App Engine Locally..
echo ========== ========== ========== ========== ==========

echo ========== ========== ========== ========== ==========
echo Adding Current Folder to GOPATH
echo ========== ========== ========== ========== ==========

set GOPATH=%GOPATH%;%cd%
echo %GOPATH%

rem pause

echo ========== ========== ========== ========== ==========
echo First, Check for Google App Engine Updates
echo ========== ========== ========== ========== ==========

start cmd.exe /k "gcloud components update"

pause

echo ========== ========== ========== ========== ==========
echo Second, Open Browser to the Admin and App itself
echo ========== ========== ========== ========== ==========

rem Start the Backend
start http://localhost:8000/

rem Start the frontend
start http://localhost:8080/

echo ========== ========== ========== ========== ==========
echo Finally, Start Google App Engine
echo ========== ========== ========== ========== ==========

rem Start Google App Engine
dev_appserver.py app.yaml

rem bat file references:
rem http://www.makeuseof.com/tag/write-simple-batch-bat-file/
