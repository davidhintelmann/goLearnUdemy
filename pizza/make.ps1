# PowerShell file for building building and running
# go code in this directory.
# This file will execute the following steps:
#   1. build go binary
#   2. run go binary
#   3. delete go binary
$binary = "app.exe"
go build -o $binary main.go
Start-Process -FilePath $binary
# .\app.exe
Remove-Item $binary