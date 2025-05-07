// Handling HTTPS requests 
// command to generate key.pem and cert.pem files 
openssl req -x509 -newkey rsa:2048 -nodes -keyout key.pem -out cert.pem -days 365 

To run your Gin server with HTTPS locally, you'll need a self-signed TLS certificate and private key. Here’s how you can do it step by step:

🔧 Step 1: Generate self-signed certificate and key
Run this command in your terminal (requires OpenSSL):

->openssl req -x509 -newkey rsa:2048 -nodes -keyout key.pem -out cert.pem -days 365

It will ask you a few questions. You can skip most of them by pressing Enter. For Common Name, use localhost.

This creates:

cert.pem → certificate file

key.pem → private key

