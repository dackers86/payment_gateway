package main

import (
    "fmt"
    "net/http"
    "github.com/stripe/stripe-go"
    "github.com/stripe/stripe-go/charge"
)

func AcceptPayment(w http.ResponseWriter, r *http.Request) {

   stripe.Key = "sk_test_L8nrpi6m2KzwGolKsCN86pqJ"

   params := &stripe.ChargeParams{
     Amount: 1000,
     Currency: "gbp",
     Desc: "Example charge",
     Token: "ggdffdgfgffff",
   }

  //  Token:    r.PostFormValue("stripeToken")

     fmt.Fprintf(w, "Start charge!")

   _, err := charge.New(params)

   fmt.Fprintf(w, "Got here!")

   if err == nil {
   		fmt.Fprintf(w, "Successful test payment!")
   	} else {
   		fmt.Fprintf(w, "Unsuccessful test payment: "+err.Error())
   	}
}

func main() {
    http.HandleFunc("/", AcceptPayment)
    http.ListenAndServe(":8080", nil)
}
