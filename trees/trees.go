package trees

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/adamsanghera/trees/trees/treespb"

	"github.com/adamsanghera/trees/trees/storage"
	"github.com/adamsanghera/trees/trees/storage/rds"
)

// Trees serves data needed to support the trees app
type Trees struct {
	strg storage.Storage

	srv *http.Server
}

// New creates a new trees server
func New(password string) (*Trees, error) {
	rds, err := rds.New(password)
	if err != nil {
		return nil, err
	}

	t := &Trees{
		strg: rds,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/near-me", func(w http.ResponseWriter, req *http.Request) {
		latStr := req.URL.Query().Get("latitude")
		lonStr := req.URL.Query().Get("longitude")
		radStr := req.URL.Query().Get("radius")
		lat, err := strconv.ParseFloat(latStr, 32)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Couldn't parse latitude from url"))
			return
		}
		lon, err := strconv.ParseFloat(lonStr, 32)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Couldn't parse longitude from url"))
			return
		}
		rad, err := strconv.ParseFloat(radStr, 32)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Couldn't parse radius from url"))
			return
		}

		var limit int
		limitStr := req.URL.Query().Get("limit")
		if limitStr != "" {
			limit, err = strconv.Atoi(limitStr)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Couldn't parse limit from url"))
				return
			}
		} else {
			limit = 100
		}

		spcCommon := req.URL.Query().Get("spc_common")
		spcLatin := req.URL.Query().Get("spc_latin")
		treeID := req.URL.Query().Get("tree_id")

		searchReq := &treespb.SearchRequest{
			Origin: &treespb.Location{
				Lat: float32(lat),
				Lon: float32(lon),
			},
			Limit:   int32(limit),
			Radius:  float32(rad),
			Filters: []*treespb.Filter{},
		}

		if spcCommon != "" {
			searchReq.Filters = append(searchReq.Filters, &treespb.Filter{
				Key:   treespb.FilterKey_spc_common,
				Value: spcCommon,
			})
		}
		if spcLatin != "" {
			searchReq.Filters = append(searchReq.Filters, &treespb.Filter{
				Key:   treespb.FilterKey_spc_latin,
				Value: spcLatin,
			})
		}
		if treeID != "" {
			searchReq.Filters = append(searchReq.Filters, &treespb.Filter{
				Key:   treespb.FilterKey_tree_id,
				Value: treeID,
			})
		}

		resp, err := t.strg.Search(searchReq)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to search database " + err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to encode JSON response " + err.Error()))
			return
		}
	})

	t.srv = &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	return t, nil
}

// Start starts the http server
func (t *Trees) Start() error {
	return t.srv.ListenAndServe()
}
