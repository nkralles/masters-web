package routes

import (
	"encoding/csv"
	"fmt"
	"github.com/nkralles/masters-web/internal/persistence"
	"html/template"
	"net/http"
	"time"
)

func GetEntries(w http.ResponseWriter, r *http.Request) {
	entries, err := persistence.DefaultDriver().GetEntries(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJSON(w, entries, http.StatusOK)
}

type TableEntries struct {
	Entries []TableEntry
}

type TableEntry struct {
	Name           string
	Top1           string
	Top1Score      string
	Top2           string
	Top2Score      string
	Top3           string
	Top3Score      string
	Wildcard1      string
	Wildcard1Score string
	Wildcard2      string
	Wildcard2Score string
	Wildcard3      string
	Wildcard3Score string
	Wildcard4      string
	Wildcard4Score string
	Wildcard5      string
	Wildcard5Score string
	Winning        string
	Total          string
}

func GetEntriesHtml(w http.ResponseWriter, r *http.Request) {
	entries, err := persistence.DefaultDriver().GetEntries(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := TableEntries{}
	for _, entry := range *entries {
		golfers := *entry.Golfers
		if golfers == nil || len(golfers) != 8 {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("shit broke"))
			return
		}
		data.Entries = append(data.Entries, TableEntry{
			Name:           entry.Name,
			Top1:           fmt.Sprintf("%s %s", golfers[0].FirstName, golfers[0].LastName),
			Top1Score:      fmt.Sprintf("%d", golfers[0].Score),
			Top2:           fmt.Sprintf("%s %s", golfers[1].FirstName, golfers[1].LastName),
			Top2Score:      fmt.Sprintf("%d", golfers[1].Score),
			Top3:           fmt.Sprintf("%s %s", golfers[2].FirstName, golfers[2].LastName),
			Top3Score:      fmt.Sprintf("%d", golfers[2].Score),
			Wildcard1:      fmt.Sprintf("%s %s", golfers[3].FirstName, golfers[3].LastName),
			Wildcard1Score: fmt.Sprintf("%d", golfers[3].Score),
			Wildcard2:      fmt.Sprintf("%s %s", golfers[4].FirstName, golfers[4].LastName),
			Wildcard2Score: fmt.Sprintf("%d", golfers[4].Score),
			Wildcard3:      fmt.Sprintf("%s %s", golfers[5].FirstName, golfers[5].LastName),
			Wildcard3Score: fmt.Sprintf("%d", golfers[5].Score),
			Wildcard4:      fmt.Sprintf("%s %s", golfers[6].FirstName, golfers[6].LastName),
			Wildcard4Score: fmt.Sprintf("%d", golfers[6].Score),
			Wildcard5:      fmt.Sprintf("%s %s", golfers[7].FirstName, golfers[7].LastName),
			Wildcard5Score: fmt.Sprintf("%d", golfers[7].Score),
			Winning:        fmt.Sprintf("%d", entry.WinningScore),
			Total:          fmt.Sprintf("%d", entry.Total),
		})
	}

	tmpl, err := template.ParseFiles("./internal/routes/table-template.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
func GetEntriesCSV(w http.ResponseWriter, r *http.Request) {
	entries, err := persistence.DefaultDriver().GetEntries(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=entries%s.csv", time.Now().UTC().Format("2006-01-02-15_04_05")))
	w.Header().Set("Transfer-Encoding", "chunked")
	writer := csv.NewWriter(w)
	writer.Write([]string{"Entry Name", "Top 12", "Top 12", "Top 12",
		"Wildcard", "Wildcard", "Wildcard", "Wildcard", "Wildcard", "Winning Score", "Total"})

	for _, entry := range *entries {
		golfers := *entry.Golfers
		if golfers == nil || len(golfers) != 8 {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("shit broke"))
			return
		}
		writer.Write([]string{entry.Name,
			fmt.Sprintf("%s %s\n%d", golfers[0].FirstName, golfers[0].LastName, golfers[0].Score),
			fmt.Sprintf("%s %s\n%d", golfers[1].FirstName, golfers[1].LastName, golfers[1].Score),
			fmt.Sprintf("%s %s\n%d", golfers[2].FirstName, golfers[2].LastName, golfers[2].Score),
			fmt.Sprintf("%s %s\n%d", golfers[3].FirstName, golfers[3].LastName, golfers[3].Score),
			fmt.Sprintf("%s %s\n%d", golfers[4].FirstName, golfers[4].LastName, golfers[4].Score),
			fmt.Sprintf("%s %s\n%d", golfers[5].FirstName, golfers[5].LastName, golfers[5].Score),
			fmt.Sprintf("%s %s\n%d", golfers[6].FirstName, golfers[6].LastName, golfers[6].Score),
			fmt.Sprintf("%s %s\n%d", golfers[7].FirstName, golfers[7].LastName, golfers[7].Score),
			fmt.Sprintf("%d", entry.WinningScore),
			fmt.Sprintf("%d", entry.Total),
		})
		writer.Flush()
	}
}
