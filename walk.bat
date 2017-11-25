@echo off
if [%1]==[] (
	echo must use as "walk gocmd"
	goto :end
)
set root=%CD%
::echo %root%
for /r /d %%d in (*main *src *outofgopath) do (
	echo %%d|findstr "main$ src$" > nul && (
		cd %%d
		echo %%d
		go %1
	)
)
::echo %root%
cd %root%

echo press any key close...
pause

:end