# momo-api

`momo-api` is an unofficial MoMo API written in Go. MoMo is currently #1 digital wallet and mobile payment service in Vietnam.

The API was last updated at `Feb, 2023`. Since I've discontinued the development, I decided to open-source this one, have fun :D

## Disclaimer
**`momo-api` is provided "as is" without warranties. You assume all risks for legal compliance and licensing. No support is offered for installation or troubleshooting.**

_`momo-api` should only be used for educational and personal projects._

## Usage
### Configuration
- Edit `utils/constants.go` for phone configuration:
    + Device: either Android or iPhone (iPhone proves to be more stable)
    + App version / App code: see the below section for guidance

### Test
The test file contains examples to work with the MoMo API.

Require following environment varables:
- test_phone: Your own phone
- test_pass: Password to login
- receiver_phone: The receiver's phone
- receiver_name: The receiver's name

### State persistence
- You have to store the state by your own (e.g: JSON serialization)

## Auth API
### Session validation
- `auth#VerifyPhone` should be sent first to validate session
  + Success: can skip authentication
  + Failure: continue with OTP log-in
- MoMo session will expire after a certain time (currently unknown)

### Log In
- `auth#RequestOTP` is used to request for OTP
- `auth#VerifyOTP` is used to verify the OTP (given the OTP sent to your phone)
- `auth#Login` is used to log in (given the password)

*Note*: MoMo disallows concurrent sessions over multiple devices. If you log in to a new device, the old sessions will be expired.

### Re Login, Logout
- `auth#Relogin` can be used to skip OTP and login (if use the same device)
- `auth#Logout` to log out an account

## Chat API
- `chat/room_api.go` contains stuff to fetch room chats
- `chat/message_api.go` contains stuff fetch message chats in a room

## Notification API
- `noti/api.go` contains stuff to fetch notifications

## Transaction API
- `noti/trans.go` contains stuff to browse transactions
- `noti#InitTransfer` to init the transaction (MoMo will pre-check available balance at this stage)
- `noti#ConfirmTransfer` to confirm the transaction

# Internal
- If you want to extend the API or simply change the app version, you have to touch the internal code of MoMo. This can be done by downloading the apk of MoMo and use a Java decompiler.
- To analyze the connection, install Charles to capture the packets (both in your phone & PC)
  + Phone: connects to your PC's IP
  + PC: acts as middleware to filter and analyze network activities
- All common MoMo endpoints are listed at `utils/constants.go`
- MoMo uses SSL and some secure APIs may require additional encryption
- To find out the app version, check out the info of an arbitrary request to the MoMo servers
- To work on features, you have to check the network and look at the decompiled app code.
- The MoMo app contains a core written in Java and several modules written in React Native. The React Native module uses Hermes VM instructions so it is much harder to analyze. Recommend tool is https://github.com/P1sec/hermes-dec
