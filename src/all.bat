:: Copyright 2012 The Go Authors. All rights reserved.
:: Use of this source code is governed by a BSD-style
:: license that can be found in the LICENSE file.

@echo off

setlocal

if exist make.bat goto ok
echo all.bat must be run from go\src
exit /b 1
:ok

call .\make.bat --no-banner || goto fail
call .\run.bat --no-rebuild || goto fail
..\bin\go tool dist banner
goto :eof

:fail
exit /b 1
