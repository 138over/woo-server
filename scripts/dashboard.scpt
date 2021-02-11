#!/usr/bin/osascript                                                                            

on run argv
    tell app "Terminal"
        set newTab to do script "sampler -c " & (item 1 of argv) 
        set current settings of newTab to settings set "Pro"
        set the position of the front window to {600, 22}
        set the size of the front window to {1200, 600}
    end tell
end run
