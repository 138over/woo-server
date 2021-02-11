
# ------------------------------------------------------------------------
# add isobox boxname...
#   Given a path to an iso-box and a list of box names
#   Register each box with the iso-box
# ------------------------------------------------------------------------
add() {
    local isofile=$1; shift
    local errno;

    if [ ! -f $isofile ]; then
        echo "$isofile does not exist"
        return 1
    fi
    
    set +e
    for i in $@; do
        echo vagrant add --name $i $isofile
    done
    set -e
    return errno
}

# ------------------------------------------------------------------------
# boxes registered|unregistered boxname...
#   Given a list of box names and a state of registered or unregistered
#   Return box names from the list that match the specified state
# ------------------------------------------------------------------------
boxes() {
    local found;
    if [ $1 = "registered" ]; then 
        found=0
    elif [ $1 = "unregistered" ]; then
        found=1
    fi
    shift
    local result=""
    local list=$(vagrant box list)
    for i in $@; do
        set +e
        echo "$list" | grep -q ^$i[[:space:]]
        if [ $? -eq $found ]; then 
            result="$result $i"; 
        fi
        set -e
    done

    echo $result
}

# ------------------------------------------------------------------------
# vms running|poweredoff boxname...
#   Given a list of box names and a vm state of running or poweredoff
#   Return box names from the list that match the specified state
# ------------------------------------------------------------------------
vms() {
    local state=$1; shift
    local boxes=$@
    local result=""
    local global_status=$(vagrant global-status)
    for i in $boxes; do
        set +e
        echo "$global_status" | grep -q $i'[[:blank:]]*virtualbox[[:blank:]]'$state
        if [ $? -eq 0 ]; then
            result="$result $i"
        fi
        set -e

    done

    echo $result
}

# ------------------------------------------------------------------------
# boxes_not_running 
#   Given a list of box names 
#   Return box names from the list that are powered off or never powered on
# ------------------------------------------------------------------------
boxes_not_running() {
    local state=poweroff;
    local boxes=$@
    local result=""
    local global_status=$(vagrant global-status)

    for i in $boxes; do
        set +e
        echo "$global_status" | grep -q $i'[[:blank:]]*virtualbox[[:blank:]]'
        if [ $? -ne 0 ]; then
            # box has never been powered on, add to result
            result="$result $i"
            continue
        fi

        echo "$global_status" | grep -q $i'[[:blank:]]*virtualbox[[:blank:]]'$state
        if [ $? -eq 0 ]; then
            # box is powered off, add to result
            result="$result $i"
        fi
        set -e
    done

    echo $result
}
