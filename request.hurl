POST http://localhost:8000/payments
Content-Type: application/json
{
	"amount": 10000,
	"currency": "BRL",
	"description": "Pagamento do Pix",
	"paymentType": "card",
	"card": {
	    "number": "0000 0000 0000 0000",
	    "holderName": "Teste Teste",
	    "cvv": "123",
	    "expirationDate": "12/2345",
	    "installments": 1
	}
}
HTTP 200
