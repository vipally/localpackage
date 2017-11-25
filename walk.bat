@echo off
if [%1]==[] (
	echo must use as "walk gocmd"
	goto :end
)
set root=%CD%
::echo %root%
for /r /d %%d in (*main *src) do (
	echo %%d|findstr "main$ src$" > nul && (
		cd %%d
		echo processing %%d ...
		go %1
		if [%2]==[run] (
			call %%~nd.exe
		)
	)
)
::echo %root%
cd %root%

echo press any key close...
pause

:end