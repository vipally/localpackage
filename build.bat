@echo off
set root=%CD%
echo %root%
FOR %%d in (
	/gopath/src
	/gopath/src/main
	/gopath/outofgopath
	/gopath/outofgopath/main
	) DO (
		set dir=%CD%%%d
		cd %dir%
		go build -i  .
		echo %CD% %%d %dir%

	)
echo %root%
cd %root%

rem go build -i
echo press any key close
pause

goto :end
::find all main package
for /r /d %%d in (*) do (
	echo %%d|findstr "main$ src$" > nul && (
		echo %%d ok
	)
)
		set dd = %%d
		set dir = %CD%%dd%
		pushd %dir%
		echo %CD%
		echo %dir%
		::go build -i
		popd
E:\localpackage\gopath\src ok
E:\localpackage\gopath\outofgopath\main o
E:\localpackage\gopath\src\main ok
E:\localpackage\localroot\src ok
E:\localpackage\localroot\src\main ok
E:\localpackage\localroot2\src ok
E:\localpackage\localroot2\src\main ok
E:\localpackage\localroot3\src ok
E:\localpackage\localroot3\src\main ok
:end

		