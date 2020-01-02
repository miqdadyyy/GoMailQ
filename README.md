# Mail Queue
### Description
Microservice to send email asyncrhonusly by sending `address`, `subject`, and `content`.
### Usage 
#### Instalation : 
- Clone this project
- Make sure your machine has Go with minimum version 1.13
- Copy .env.example to .env and fill the requirements.
- Run `go get` on this directory
- Run `go mod download` too
- Type `go run main.go`
#### Routes : 
`/` (GET) : index  
`/send` (POST) : sending email
#### Environment : 
copy `.env.example` to `.env` and fill the requirements.
> - SMTP_HOST : host of your SMTP server
> - SMTP_PORT : port of your SMTP server
> - SMTP_USERNAME : username of your SMTP server
> - SMTP_PASSWORD : password of your SMTP server
> - SENDER_EMAIL : email of sender
> - PORT : port of this microservice
#### Example : 
In this case i used GuzzleHttp on Laravel for example
```php
$g = new \GuzzleHttp\Client();
$addresses = ["miqdad.farcha@gmail.com", "cibaicode321@gmail.com"];
$content = new \App\Mail\InvoiceCreated(\App\Invoice::first());
$data = [
    "addresses" => $addresses,
        "email" => [
            "subject" => "Testing World",
            "content" => $content->render()
        ]
    ];
$g->postAsync("http://0.0.0.0:3000/send", [\GuzzleHttp\RequestOptions::JSON => $data])->wait();
``` 
Data : 
```
[
  "addresses" => $addresses,
          "email" => [
              "subject" => $subject
              "content" => $content
          ]
      ]
]
```
`$addresses` : Array of target email  
`$content` : Content of email (HTML format)  
`$subject` : Subject email
