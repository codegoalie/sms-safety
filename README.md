# SMS Safety

A SMS proxy to handle one time codes and other SMS that you don't want to use
your own personal phone number for.

Host this server somewhere and setup a twilio integration to `POST /sms`.

## Configuration

There are several environment variables which need to be set to get the app to run:

`TWILIO_ACCOUNT_SID` - Your Account SID from Twilio
`TWILIO_AUTH_TOKEN` - Your Auth Token from Twilio
`RECEIVER_NUMBER` - Your real number to forward SMS messages to
`TWILIO_NUMBER` - Your Twilio SMS number to send forwarded messages from. This is also the number to enter into 3rd party services
