# get env file path from the first special arg
env=$1/env
shift

# get message from arg
msg=$1

# pass message to env, format: "key \t value"
echo -e "mymsg\t$msg" >> $env
echo "Sent: $msg"
