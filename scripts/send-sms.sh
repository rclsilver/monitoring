#!/usr/bin/env bash
#
# push notification via smsapi.free-mobile.fr
#
LOGGEROPT="-p security.notice -t SMS"
REPLOGG=1
REPMAIL=1
MAILTO=root
SEND_SMS_URL=${SEND_SMS_URL:-https://smsapi.free-mobile.fr/sendmsg}

report() {
    [ ${REPLOGG} -eq 1 ] && echo $@ | logger ${LOGGEROPT}
    [ ${REPMAIL} -eq 1 ] && echo $@ | mail -s "SMS $@" ${MAILTO}
}

if [ -z "${SEND_SMS_USER_ID}" ]; then
    echo "missing variable 'SEND_SMS_USER_ID'" >&2
    exit 1
fi

if [ -z "${SEND_SMS_API_KEY}" ]; then
    echo "missing variable 'SEND_SMS_API_KEY'" >&2
    exit 1
fi

eval $(${CURL} --insecure -G --write-out "c=%{http_code} u=%{url_effective}" \
    -o /dev/null -L -d user=${SEND_SMS_USER_ID} -d pass=${SEND_SMS_API_KEY} --data-urlencode msg="$(cat|tr '\n' '\r')" ${SEND_SMS_URL} 2>/dev/null | tr '&' ';')

echo ${c} ${u}\&pass=***\&msg=${msg} | logger ${LOGGEROPT}

case ${c} in
    "200")
        exit 0
        ;;
    "400")
        report "parameter missing" ; exit 400
        ;;
    "402")
        report "too many SMS..." ; exit 402
        ;;
    "403")
        report "service unavailable for user, or wrong credentials" ; exit 403
        ;;
    "500")
        report "server error, try later" ; exit 500
        ;;
    *)
        report "unexpected result" ; exit 600
        ;;
esac
