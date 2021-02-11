

check() {
    local base_version=$1
    shift

    if /usr/libexec/java_home -v ${base_version} >/dev/null 2>&1; then 
        echo `java -version 2>&1 | grep openjdk`; 
    else 
        echo "jdk ${base_version}+ not installed"; 
    fi
}

install() {
    local base_version=$1
    local cask_version=$2
    shift 2
    if /usr/libexec/java_home -v ${base_version} 1>/dev/null; then
        echo "jdk newer than ${base_version}+ already installed"; 
    else
        install_brew $cask_version
    fi
}

install_brew() {
    local cask_version=$1; shift
    brew update
    brew tap
    brew tap homebrew/cask-versions
    brew tap adoptopenjdk/openjdk
    brew search jdk 
    brew cask info homebrew/cask-versions/${cask_version}
    brew cask install homebrew/cask-versions/${cask_version}
    /usr/libexec/java_home -V
    java -version
}

uninstall() {
    local base_version=$1;
    local cask_version=$2;
    shift 2;
    if /usr/libexec/java_home -v ${base_version} >/dev/null 2>&1; then 
        brew cask uninstall homebrew/cask-versions/${cask_version};
    fi
    check $base_version
}

uninstall_oracle() {
    @echo Uninstall Oracle JDK 8 manually
    sudo rm -rf /Library/Java/JavaVirtualMachines/jdk1.8.0_181.jdk
    sudo rm -rf /Library/PreferencePanes/JavaControlPanel.prefPane
    sudo rm -rf /Library/Internet\ Plug-Ins/JavaAppletPlugin.plugin
    sudo rm -rf /Library/Internet\ Plug-Ins/JavaAppletPlugin.plugin
    sudo rm -rf /Library/Internet\ Plug-Ins/JavaAppletPlugin.plugin
    sudo rm -rf /Library/LaunchAgents/com.oracle.java.Java-Updater.plist
    sudo rm -rf /Library/PrivilegedHelperTools/com.oracle.java.JavaUpdateHelper
    sudo rm -rf /Library/LaunchDaemons/com.oracle.java.Helper-Tool.plist
}

