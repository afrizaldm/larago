@echo off
setlocal

REM Nama executable
set EXECUTABLE_NAME=main.exe

REM Membuat folder build jika belum ada
if not exist build (
    mkdir build
)

REM Menyalin folder views ke dalam folder build
xcopy views build\views /E /I /Y

REM Membuat executable dalam folder build
go build -o build\%EXECUTABLE_NAME% .\cmd\main.go

endlocal
