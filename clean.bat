@echo off
set root=%CD%
echo %root%
for /r /d %%d in (*) do (
	echo %%d|findstr "main$ src$" > nul && (
		cd %%d
		echo %%d
		go clean -i
	)
)
echo %root%
cd %root%
pause