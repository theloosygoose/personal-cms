package middleware

import(
    "net/http"
    "log"
)

func AllowCors(next http.Handler) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Origin, Accept")
            w.Header().Set("Access-Control-Allow-Origin", "*")

        if (r.Method == "OPTIONS") {
            w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
            w.WriteHeader(http.StatusOK)
            log.Println("OPTIONS ROUTES HIT")
        }

        next.ServeHTTP(w, r)
    })
}
