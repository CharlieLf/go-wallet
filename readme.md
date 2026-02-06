Simple wallet withdraw, topup and get balance with Golang

Using clean and scalable architecture

Header :

Content-Type : application/json

x-api-key	: {{your-api-key}}


API List :

METHOD : GET

URL : {{server_url}}/api/v1/wallet/balance?user_id={value}


METHOD : POST
URL : {{server_url}}/api/v1/wallet/withdraw

Body Example :

{

    "user_id":"1",

    "amount":26000.00,

    "source":"Tarik tunai"

}



URL : {{server_url}}/api/v1/wallet/add

Body Example :

{

    "user_id":"1",

    "amount":125000.00,

    "source":"Transfer"

}
