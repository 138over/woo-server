
start() {
    sudo /Library/Application\ Support/VirtualBox/LaunchDaemons/VirtualBoxStartup.sh  start
}

stop() {
    pkill -HUP -f /Applications/VirtualBox.app/Contents/MacOS/VirtualBox || true
    pkill -HUP -f /Applications/VirtualBox.app/Contents/MacOS/VBoxHeadless || true
    sudo /Library/Application\ Support/VirtualBox/LaunchDaemons/VirtualBoxStartup.sh stop
}

driversNotLoaded() {
    local list=""
    kextstat -lb org.virtualbox.kext.VBoxDrv 2>&1 | grep -q org.virtualbox.kext.VBoxDrv || list="VBoxDrv";
    kextstat -lb org.virtualbox.kext.VBoxUSB 2>&1 | grep -q org.virtualbox.kext.VBoxUSB || list="$list VBoxUSB";
    kextstat -lb org.virtualbox.kext.VBoxNetFlt 2>&1 | grep -q org.virtualbox.kext.VBoxNetFlt || list="$list VBoxNetFlt";
    kextstat -lb org.virtualbox.kext.VBoxNetAdp 2>&1 | grep -q org.virtualbox.kext.VBoxNetAdp || list="$list VBoxNetAdp";
    echo $list
}

isRunning() {
    test -z "$(driversNotLoaded)"
}

