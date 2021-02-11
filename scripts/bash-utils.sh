# ------------------------------------------------------------------------
# Given variable x set to 1
# Then is_var_set x || die 1 "Error: variable x not set"
# ------------------------------------------------------------------------
is_var_set() {
    local errno=0
    if [[ -z "$1" ]] || [[ -z "${!1}" ]]; then
        errno=1
    fi
    return errno
}

# ------------------------------------------------------------------------
# Given path /foo
# Then dir_exist /foo || die 1 "Error: directory /foo does not exist"
# ------------------------------------------------------------------------
dir_exist() {
    local errno=0
    local dir="$1"
    if [[ -z "$dir" ]] || [[ ! -d $dir ]]; then
        errno=1
    fi
    return errno
}

# ------------------------------------------------------------------------
# Given required environment variable names
# Then ensure_env_set TEST_HOME JAVA_HOME || die 1 "Error: required environment variables not set"
# ------------------------------------------------------------------------
ensure_env_set() {
    local errno=0
    local i
    for i in $@; do
        # echo "environment $i = ${!i}"
        if [[ -z "${!i}" ]] || [[ "${!i}" = "null" ]]; then
            echo "Error: $i not set"
            errno=1
        fi
    done
    return $errno
}

# ------------------------------------------------------------------------
# Given a collection of variables containing options
#   opt_config=${opt_config:-''}
#   opt_logfile=${opt_logfile:-''}
#   opt_ipaddress=${opt_ipaddress:-'127.0.0.1'}
#   opt_port=${opt_port:-'8200'}
# Then ensure_arguments_set opt_ || die 1 "Error: required options are not set"
# ------------------------------------------------------------------------
ensure_options_set() {
    local errno=0
    local i
    for i in `compgen -A variable | grep $1`; do
        if [[ -z "${!i}" ]]; then
            echo "Error: $i not set" | sed 's/opt_/--/'
            errno=1
        fi
    done
    return $errno
}

# ------------------------------------------------------------------------
#
# ------------------------------------------------------------------------
logger() {
  echo "[$(date +'%Y-%m-%dT%H:%M:%S')]: $@" >&2
}

# ------------------------------------------------------------------------
# Given a path /foo
# Then file_exist /foo || die 1 "Error: file /foo does not exist"
# ------------------------------------------------------------------------
file_exist() {
    local errno=0
    local file="$1"
    if [[ -z "$file" ]] || [[ ! -f $file ]]; then
        errno=1
    fi
    return $errno
}

# ------------------------------------------------------------------------
# ------------------------------------------------------------------------
if_grep() {
    local errno=0
    grep "$@" || errno=$?
    return $(( errno == 1 ? 0 : ernno ))       
}

# ------------------------------------------------------------------------
# ------------------------------------------------------------------------
get_json_data() {
    local filter="$1"
    local data="$2"
    $(echo $data | jq ''${filter}'')
}

# ------------------------------------------------------------------------
# ------------------------------------------------------------------------
read_json_data() {
    local filter="$1"
    local file="$2"
    local data="$(cat $file)"
    get_json_data $filter "$data"
}

# ------------------------------------------------------------------------
# ------------------------------------------------------------------------
get_json_value() {
    local filter="$1"
    local data="$2"
    $(echo $data | jq ''${filter}'' | sed 's/"//g')
}

# ------------------------------------------------------------------------
# ------------------------------------------------------------------------
get_json_key_value() {
    local filter="$1"
    local data="$2"
    local list=""
    local line=""

    while read line; do
        set $line
        list="$1:$2 $list"
    done < <(echo $data | jq ''${filter}' | to_entries' | jq '.[] | [.key, .value] | join(" ") ' | sed 's/"//g')

    echo "$list"
}

# ------------------------------------------------------------------------
# ------------------------------------------------------------------------
is_json_valid() {
    local file="$1"
    local exit_status=$(jq '.' $file 2>&1 > /dev/null)
    if [[ ! -z "$exit_status" ]]; then return 1; else return 0; fi
}

# ------------------------------------------------------------------------
# ------------------------------------------------------------------------
is_tcp_listening() {
    errno=$(lsof -t -i:$1 -sTCP:LISTEN 2>&1)
    if [[ $? = 0 ]]; then
        return 1
    else
        return 0
    fi
}

# ------------------------------------------------------------------------
# ------------------------------------------------------------------------
trap_subcommand() {
    trap_handler="$1"; 
    shift
    for signal in "$@"; do
        trap "$trap_handler $signal" "$signal"
    done
}
