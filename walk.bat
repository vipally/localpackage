@echo off
if [%1]==[] (
	echo must use as "walk gocmd"
	goto :end
)
set GOPATH=%CD%\gopath\;%GOPATH%
set root=%CD%
::echo %root%
for /r /d %%d in (*main *src) do (
	echo %%d|findstr "main$ src$" > nul && (
		if exist %%d\*.go (
			cd %%d
			if [%2]==[run] (
				echo ________________________________________________________________
			)
			
			echo ====processing %%d 
			go %1
			if [%2]==[run] (
				call %%~nd.exe
			)
		) else (
			echo %%d invalid 
		)
	)
)
::echo %root%
cd %root%

::echo press any key close...
::pause

:end