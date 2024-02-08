package handlers

import "net/http"

func (h *Handlers) CarArrival(w http.ResponseWriter, r *http.Request) {
	//change space status?
}

func (h *Handlers) CarDeparture(w http.ResponseWriter, r *http.Request) {
	//find reservation for car
	//add parking details to reservation table
	//send information to client?
}
