set GOPATH=%~dp0\..\..

ECHO mkdir bin and pkg
md %GOPATH%\bin
md %GOPATH%\pkg

ECHO go get packages
go get github.com/akavel/rsrc
go get github.com/lxn/walk

ECHO rsrc generate .syso
%GOPATH%\bin\rsrc.exe -manifest %~dp0\main.manifest -ico %~dp0\icon\DeviceCenter.ico -o %~dp0\main.syso

ECHO go build
CD %~dp0
%~d0
go build -ldflags "-H=windowsgui" -o goWinGui.exe
del main.syso