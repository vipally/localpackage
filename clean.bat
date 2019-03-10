@call walk clean
@echo off
for /r  %%d in (*.exe *.a) do (
	::@echo %%d
	del /S %%d
)