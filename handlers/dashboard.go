package handlers

import "net/http"

func DashboardHandle(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/dashboard.html")
}
