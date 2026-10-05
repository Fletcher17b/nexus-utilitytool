# Nexus CLI

`nexus` is a command-line tool written in Go for quick navigation between directories through aliases linked to specific paths saving up the small headeache of running a bunch of cd commands or one long cd string. This is a personal tool, use at your own risk.

## Setup & Installation

### Prerequisites

* [Go](https://go.dev/doc/install) (1.18 or higher recommended)
* Windows

No linux support, go touch grass or smt

### Build and Run

1. Clone or extract the source repository.
```git
mkdir nexus
cd nexus
git clone ....
```

2. Build the executable binary:
```bash
go build -o nexus-cli ./cmd/nexus
```

3. Write down your aliases on the aliases.json file in this format:
```json
{
  "alias": "Z:\\Path\\to\\your\\Project"
}
```

4. Run it:
```
C:\nexus> ./nexus-cli alias
Z:\Path\to\your\Project>
```

And thats it. Or is it?
You see the fun part about this small sidequest of a project was really about tinkering with windows and some powershell scripting. 

Right now the cli only works on your project root folder which kinda defeats de purpose. You may think adding it to your PATH in the enviormental variables may solve this dilema but that is not the case since on windows when you run an executable it runs on a child process that is unable to change the dir of the parent terminal session. In order for the cli to work correctly we're gonna need a small script that works around that:

nexus.cmd: (or whatever you want the keyword to be named .cmd)
```powershell
@echo off
for /f "delims=" %%i in ('"%~dp0nexus-cli.exe" %*') do (
    cd /d "%%i"
)
```

This small script is a wrapper of the cli that circumvents the previous inconvenience. The script should live in your project root folder (dont forget to add the dir to PATH) alongside the executable. If you did everything correctly you should be able to call your .cmd script on the cmd and switch between your specified aliases.

What if I use powershell instead of the ancient cmd? Fortunately I also had that question, in order to also make it work in your powershell you just need to add this to your ps profile:

if you dont know how to open you powershell profile, go look it up 

Once you have it open add this snippet:
```powershell
function nexus {
    $path = & "A:\Users\user\Tools\nexus-cli.exe" @args

    if ($LASTEXITCODE -eq 0 -and (Test-Path $path)) {
        Set-Location $path
    } else {
        Write-Host $path
    }
}
```
Now open your powershell and regresh your settings with: . $PROFILE

if you did it correctly it should work.