package main

import (
    "fmt"
    "net/http"
    "github.com/stripe/stripe-go"
    "github.com/stripe/stripe-go/customer"
    "github.com/stripe/stripe-go/charge"
    "encoding/json"
)

type payment struct {
    Amount uint64
    Description string
    EmailAddress string
    StripeToken string
}

func AcceptPayment(w http.ResponseWriter, r *http.Request) {

   w.Header().Set("Access-Control-Allow-Origin", "*")

    decoder := json.NewDecoder(r.Body)
    var payment_details payment
    err := decoder.Decode(&payment_details)
    if err != nil {
        panic(err)
    }
    defer r.Body.Close()

   stripe.Key = "sk_test_L8nrpi6m2KzwGolKsCN86pqJ"

   customerParams := &stripe.CustomerParams{
     Email: payment_details.EmailAddress,
   }
   customerParams.SetSource(payment_details.StripeToken)
   customer, err := customer.New(customerParams)

   params := &stripe.ChargeParams{
     Amount: payment_details.Amount,
     Currency: "gbp",
     Desc: payment_details.Description,
     Token: payment_details.StripeToken,
     Customer: customer.ID,
   }

   _, err = charge.New(params)
}

func main() {
    http.HandleFunc("/", AcceptPayment)
    http.ListenAndServe(":8081", nil)
}
