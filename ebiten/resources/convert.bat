cd /d %~dp0
set TL=file2byteslice-master\cmd\file2byteslice\file2byteslice.exe

%TL% -input mascot_t16.png -output sgimg/mascot16.go -package sgimg -var Mascot16_png
%TL% -input mascot_t32.png -output sgimg/mascot32.go -package sgimg -var Mascot32_png
%TL% -input mascot_t64.png -output sgimg/mascot64.go -package sgimg -var Mascot64_png
%TL% -input mascot_t128.png -output sgimg/mascot128.go -package sgimg -var Mascot128_png
%TL% -input mascot_t256.png -output sgimg/mascot256.go -package sgimg -var Mascot256_png

%TL% -input gopher_t32.png -output sgimg/gopher32.go -package sgimg -var Gopher32_png
%TL% -input gopher_t64.png -output sgimg/gopher64.go -package sgimg -var Gopher64_png
%TL% -input gopher_t128.png -output sgimg/gopher128.go -package sgimg -var Gopher128_png

%TL% -input mascot_anger_t32.png -output sgimg/mascotA32.go -package sgimg -var MascotA32_png

%TL% -input ebiten_t32.png -output sgimg/ebiten32.go -package sgimg -var Ebiten32_png
%TL% -input ebiten_power_t32.png -output sgimg/ebitenP32.go -package sgimg -var EbitenP32_png
%TL% -input ebiten_time_t32.png -output sgimg/ebitenT32.go -package sgimg -var EbitenT32_png


pause