# GUI Wallet

## Build

1. fyne.exe package -os windows -name govm.net -icon ./assets/govm.png
2. mv gui.exe govm.exe
3. tar zcvf  govm_windows_$(date +'%Y%m%d_%H%M%S').tar.gz govm.exe assets conf.json

## add local language

1. add translation: assets/zh.json
2. copy ttf from C:\Windows\Fonts, rename to assert/zh.ttf
3. add 'zh' to conf.json->langure_list, split by ','
4. use the language: conf.json->langure set value to zh

## others

fyne package -os android -appID net.govm -icon ./assets/govm.png -name govm

gomobile build -target=android

go build -ldflags -H=windowsgui
