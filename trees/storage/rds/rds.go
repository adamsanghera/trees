package rds

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

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

// GetDetails does what it says
func (pg *Postgres) GetDetails(req *treespb.GetDetailsRequest) (*treespb.GetDetailsResponse, error) {
	r := pg.db.QueryRow(`
	SELECT tree_id, created_at, tree_diameter,
		stump_diameter, status, health, spc_latin, spc_common,
		steward, curb_location, guards, sidewalk, user_type, problems,
		root_stone, root_grate, root_other, trunk_other, trunk_wire, trunk_light,
		branch_light, branch_shoe, branch_other, address, zipcode, zip_city,
		borough_name, ST_X(location), ST_Y(location)
	FROM trees
	WHERE tree_id=$1`, req.TreeId)

	var (
		treeID, treeD, stumpD                                              sql.NullInt64
		createdAt                                                          time.Time
		status, health, spcLatin, spcCommon, steward                       sql.NullString
		curbLocation, guards, sidewalk, userType, problems                 sql.NullString
		rootStone, rootGrate, rootOther, trunkOther, trunkWire, trunkLight string
		branchLight, branchShoe, branchOther, address                      string
		zipcode                                                            int64
		zipCity, boroughName                                               string
		lat, lon                                                           float32
	)
	err := r.Scan(&treeID, &createdAt, &treeD, &stumpD,
		&status, &health, &spcLatin, &spcCommon, &steward,
		&curbLocation, &guards, &sidewalk, &userType, &problems,
		&rootStone, &rootGrate, &rootOther, &trunkOther, &trunkWire,
		&trunkLight, &branchLight, &branchShoe, &branchOther,
		&address, &zipcode, &zipCity, &boroughName, &lat, &lon)
	if err != nil {
		return nil, err
	}

	return &treespb.GetDetailsResponse{
		Tree: &treespb.Tree{
			TreeId:        treeID.Int64,
			CreatedAt:     createdAt.Unix(),
			TreeDiameter:  int32(treeD.Int64),
			StumpDiameter: int32(stumpD.Int64),
			Status:        status.String,
			Health:        health.String,
			SpcCommon:     spcCommon.String,
			SpcLatin:      spcLatin.String,
			Steward:       steward.String,
			CurbLocation:  curbLocation.String,
			Guards:        guards.String,
			Sidewalk:      sidewalk.String,
			UserType:      userType.String,
			Problems:      problems.String,
			RootStone:     rootStone,
			RootGrate:     rootGrate,
			RootOther:     rootOther,
			TrunkLight:    trunkLight,
			TrunkOther:    trunkOther,
			TrunkWire:     trunkWire,
			BranchLight:   branchLight,
			BranchOther:   branchOther,
			BranchShoe:    branchShoe,
			Address:       address,
			Zipcode:       strconv.Itoa(int(zipcode)),
			ZipCity:       zipCity,
			BoroughName:   boroughName,
			Location: &treespb.Location{
				Lat: lat,
				Lon: lon,
			},
		},
	}, nil
}
