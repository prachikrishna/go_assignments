package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home Prachi!")
}

/*func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	log.Fatal(http.ListenAndServe(":8080", router))
}*/

//creating a dummy database
type contact struct {
	ID   string `json:"ID"`
	Name string `json:"Name"`
	//Title         string `json:"Title"`
	//First_Name       string `json:"First_Name"`
	//Second_Name			string `json:"Second_Name"`
	Number string `json:"Number"`
}

type allContacts []contact

var contactlist = allContacts{
	{
		ID:     "1",
		Name:   "Joey",
		Number: "987654321",
	},
}

//createContact uses POST HTTP Method to create a new contact and accepts data from the user's end
//in the form of HTTP request data
func createContact(w http.ResponseWriter, r *http.Request) {
	var newContact contact
	// Convert r.Body into a readable formart
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the contact id, name and number only in order to update")
	}

	json.Unmarshal(reqBody, &newContact)

	// Add the newly created contact to the contactlist
	contactlist = append(contactlist, newContact)

	// Return the 201 created status code
	w.WriteHeader(http.StatusCreated)
	// Return the newly created contact
	json.NewEncoder(w).Encode(newContact)
}

//this function lists all the contacts present in the contactlist
func getAllContacts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(contactlist)
}

//readContact_By_Id fetches a particular contact with a specific ID and uses the GET HTTP method
func readContact_By_Id(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the url
	contactID := mux.Vars(r)["id"]

	// Get the details from an existing contact
	// Use the blank identifier to avoid creating a value that will not be used
	for _, singleContact := range contactlist {
		if singleContact.ID == contactID {
			json.NewEncoder(w).Encode(singleContact)
		}
	}
}

//updateContact updates an existing contact with a given id and uses the PATCH HTTP method
func updateContact(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the url
	eventID := mux.Vars(r)["id"]
	var updatedContact contact
	// Convert r.Body into a readable formart
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &updatedContact)

	for i, singleContact := range contactlist {
		if singleContact.ID == eventID {
			singleContact.Name = updatedContact.Name
			singleContact.Number = updatedContact.Number
			contactlist[i] = singleContact
			json.NewEncoder(w).Encode(singleContact)
		}
	}
}

//deleteContact deletes an existing contact with a specific id and uses the DELETE HTTP method
func deleteContact(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the url
	contactID := mux.Vars(r)["id"]

	// Get the details from an existing contact
	for i, singleContact := range contactlist {
		if singleContact.ID == contactID {
			contactlist = append(contactlist[:i], contactlist[i+1:]...)
			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", contactID)
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/contact", createContact).Methods("POST")
	router.HandleFunc("/contactlist", getAllContacts).Methods("GET")
	router.HandleFunc("/contactlist/{id}", readContact_By_Id).Methods("GET")
	router.HandleFunc("/contactlist/{id}", updateContact).Methods("PATCH")
	router.HandleFunc("/contactlist/{id}", deleteContact).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
