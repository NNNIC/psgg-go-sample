if exist testControl.go (
go run testControl.go game.go mainControl.go  gameFont.go main.go
) else (
    go run game.go mainControl.go  gameFont.go main.go
)
pause