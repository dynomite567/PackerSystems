@echo off

@"%SystemRoot%\System32\WindowsPowerShell\v1.0\powershell.exe" -NoProfile -InputFormat None -ExecutionPolicy Bypass -Command "iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))" && SET "PATH=%PATH%;%ALLUSERSPROFILE%\chocolatey\bin"
C:\Windows\system32\WindowsPowerShell\v1.0\powershell.exe -ExecutionPolicy ByPass -File Update-Needed.ps1 -username "GingerTech" -password "UberPassword"