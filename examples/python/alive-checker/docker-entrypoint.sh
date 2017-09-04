#!/bin/sh
set -x

rm -rf hosts.list /output/*

CMSTEP="2"
PYAPP="web1check.py"

APPOUT=${PYAPP/py/out}

while [[ $# -gt 0 ]]; do
    case "$1" in
        -m|--periodic)
            # for crond
			shift 
            CMSTEP= $1
            ;;
	    -a|--python-file)
		    shift
            PYAPP=$1
	        shift
			if [[ $# -gt 0 ]]; then
	            for var in $@; do
	                echo $var >> hosts.list
	            done
		    fi
			;;
	esac
	shift       # Check next set of parameters.
done

if [[ ! -e hosts.list ]]; then
    echo "warn: add localhost into hosts list"
	echo "127.0.0.1" > hosts.list
fi 

touch crontab.tmp
echo "*/$CMSTEP * * * * python /app/$PYAPP>/output/$APPOUT' > crontab.tmp
crontab -r
crontab crontab.tmp
rm -rf crontab.tmp

/usr/sbin/crond -f -d 0