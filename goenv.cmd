@echo off
if not exist bin mkdir bin
if not exist pkg mkdir pkg
if not exist src mkdir src

set GOPATH=%1
echo GOPATH is %1
set PATH=%GOPATH%bin;%PATH%

