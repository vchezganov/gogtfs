package main

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/vchezganov/gogtfs"
	"github.com/vchezganov/gogtfs/model"
)

func main() {
	data := `stop_id,level_id,stop_name,stop_lat,stop_lon,location_type,parent_station
F12,,5 Av/53 St,40.760167,-73.975224,1,
E1,L0,5 Av/53 St SW,40.760474,-73.976099,2,F12
N1,L1,,40.760457,-73.975912,3,F12
N4,L1,,40.759679,-73.974064,3,F12
F12S,,5 Av/53 St,40.760167,-73.975224,0,F12`
	dataReader := strings.NewReader(data)
	dataReadCloser := io.NopCloser(dataReader)

	reader, err := gogtfs.NewReader[model.Stop](dataReadCloser)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	for {
		stop, err := reader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Printf("error: %v\n", err)
			continue
		}

		fmt.Printf("%+v\n", stop)
	}
}
