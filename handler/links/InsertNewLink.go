package links

import (
	"encoding/json"
	"github.com/andreabreu76/encurtadorV2/entitie"
	"github.com/andreabreu76/encurtadorV2/handler/utils"
	"log"
	"net/http"
)

func InsertNewLink(w http.ResponseWriter, r *http.Request) {
	var links entitie.Link

	err := json.NewDecoder(r.Body).Decode(&links)
	if err != nil {
		log.Printf("InsertNewLink :: Error decoding json: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	links.ID, err = entitie.CreateNewLink(links);
	if err != nil {
		log.Printf("InsertNewLink :: Error creating new link: %v", err)
		utils.ErrorWithJSON(w, "Error inserting link", http.StatusInternalServerError)
		return
	}

	linkID, err := json.Marshal(links.ID);
	if err != nil {
		log.Printf("InsertNewLink :: Error marshaling link id: %v", err)
		utils.ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
		return
	}

	utils.ResponseWithJSON(w, linkID, http.StatusOK)
}
