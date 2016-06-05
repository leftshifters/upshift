function SendEmail {

  StartAction "SendEmail"

  # If 'next' is false, exit
  if [ ${next} == false ]; then
    ShowPreviousFailed
    return
  fi

  MAIL_TO=$1
  MAIL_SUBJECT=$2
  MAIL_BODY=$3

  SEND_EMAIL=$(echo "${MAIL_BODY}" | mail -s "${MAIL_SUBJECT}" "${MAIL_TO}")
  printf "${SEND_EMAIL}\n"

}

