Go to: https://windows.php.net/download

Extract it to something like:
C:\php

Add it to your PATH (in powershell, you can run setx PATH "$env:PATH;C:\php")

restart vscode and confirm php by running 'php -v'

to run the file, 'php ./printingDepartment.php'