set GOOS=js
set GOARCH=wasm
go build -o game.wasm github.com/NNNIC/psgg-go-sample/ebiten/rotatex3

::@echo off
echo :
echo : Copy to psgg-go-sampleweb ?
set /p a="Y or N"
goto :_%a%

:_Y
echo : start Copy
copy game.wasm ..\..\..\psgg-go-sampleweb\docs\*.*
pause
goto eof

:_N
pause
goto eof