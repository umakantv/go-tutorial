module hello

go 1.18

replace greetings => ../greetings

require greetings v0.0.0-00010101000000-000000000000

require github.com/razorpay/goutils/uniqueid v1.0.1
