package rds

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/adamsanghera/trees/trees/treespb"

	// pgsql driver
	_ "github.com/lib/pq"
)

var baseSearch = `
SELECT tree_id, spc_common, spc_latin, ST_X(location), ST_Y(location)
	FROM trees
WHERE ST_DistanceSphere(
	ST_SetSRID(
		St_MakePoint($1, $2), 4326
	), location
) / 1609.344 <= $3`

var tailSearch = `
ORDER BY ST_DistanceSphere(
	ST_SetSRID(
		St_MakePoint($1, $2), 4326
	), location
)`

func andKeyValue(key, val int) string {
	return fmt.Sprintf("AND $%d = $%d", key, val)
}

// Postgres wraps a db connection
type Postgres struct {
	db *sql.DB
}

// New creates a new postgres app driver, and returns it
func New(password string) (*Postgres, error) {
	db, err := sql.Open("postgres", "host=trees.cfuvv65qlkp8.us-east-2.rds.amazonaws.com "+
		"port=5432 user=cc_trees password="+password+" dbname=trees sslmode=require")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Safe connection")

	return &Postgres{
		db: db,
	}, nil
}

// Search does what it says
func (pg *Postgres) Search(req *treespb.SearchRequest) (*treespb.SearchResponse, error) {
	q := baseSearch

	// Start with lat, lon, radius, and limit
	numArgs := 4
	args := []interface{}{req.Origin.Lat, req.Origin.Lon, req.Radius, req.Limit}
	for _, f := range req.Filters {
		q += andKeyValue(numArgs+1, numArgs+2)
		numArgs += 2
		args = append(args, f.Key, f.Value)
	}
	q += tailSearch + " LIMIT $4;"

	rows, err := pg.db.Query(q, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp := make([]*treespb.CondensedTree, 0)
	for rows.Next() {
		var (
			treeID              int64
			spcCommon, spcLatin string
			lat, lon            float32
		)
		err = rows.Scan(&treeID, &spcCommon, &spcLatin, &lat, &lon)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		resp = append(resp, &treespb.CondensedTree{
			TreeId:    treeID,
			SpcCommon: spcCommon,
			SpcLatin:  spcLatin,
			Location: &treespb.Location{
				Lat: lat,
				Lon: lon,
			},
		})
	}

	return &treespb.SearchResponse{
		Trees: resp,
	}, nil
}
