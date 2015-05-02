rem Note the manifest file must exist with the executable unless it's compiled differently (not below)

go build -ldflags="-H windowsgui"

dtcnotif.exe
