cd /d %~dp0
set TL=file2byteslice-master\cmd\file2byteslice\file2byteslice.exe

%TL% -input mascot_t16.png -output sgimg/mascot16.go -package sgimg -var Mascot16_png
%TL% -input mascot_t32.png -output sgimg/mascot32.go -package sgimg -var Mascot32_png
%TL% -input mascot_t64.png -output sgimg/mascot64.go -package sgimg -var Mascot64_png
%TL% -input mascot_t128.png -output sgimg/mascot128.go -package sgimg -var Mascot128_png
%TL% -input mascot_t256.png -output sgimg/mascot256.go -package sgimg -var Mascot1256_png

pause