package main

import (
	"context"
	"flag"
	"log"
	"main/models"
	Analytics "main/protobuf"
	"net"
	"strings"

	_ "github.com/lib/pq"

	"google.golang.org/grpc"
)

var debug = false

// Server ...
type Server struct {
	Analytics.UnimplementedAnalyticsEndpointServer
	debug bool
}

func dprint(msg ...interface{}) {
	if debug {
		log.Println(msg...)
	}
}

// NewServer -- Pulls Info from DB and creates new Server Struct
func NewServer() *Server {
	s := &Server{}
	models.ConnectDB()
	return s
}

// GetHeatmap ...
// func (s *Server) GetHeatmap(ctx context.Context, in *Analytics.Empty) (*Analytics.Heatmap, error) {
// 	test := make(map[string]*Analytics.DayofWeek)
// 	test["Monday"] = &Analytics.DayofWeek{Times: make(map[string]*Analytics.Location)}
// 	test["Monday"].Times["8:30AM"] = &Analytics.Location{Data: make(map[string]int64)}
// 	test["Monday"].Times["8:30AM"].Data["1|1"] = 1
// 	test["Monday"].Times["8:30AM"].Data["2|2"] = 10
// 	test["Monday"].Times["11:30AM"] = &Analytics.Location{Data: make(map[string]int64)}
// 	test["Monday"].Times["11:30AM"].Data["5|1"] = 50
// 	test["Monday"].Times["11:30AM"].Data["2|10"] = 100

// 	test["Tuesday"] = &Analytics.DayofWeek{Times: make(map[string]*Analytics.Location)}
// 	test["Tuesday"].Times["8:30AM"] = &Analytics.Location{Data: make(map[string]int64)}
// 	test["Tuesday"].Times["8:30AM"].Data["5|1"] = 12
// 	test["Tuesday"].Times["8:30AM"].Data["4|2"] = 14

// 	return &Analytics.Heatmap{Days: test}, nil
// }

// GetHeatmap ...
func (s *Server) GetHeatmap(ctx context.Context, in *Analytics.Empty) (*Analytics.Heatmap, error) {
	dayLookup := map[byte]string{
		'M': "Monday",
		'T': "Tuesday",
		'W': "Wednesday",
		'R': "Thursday",
		'F': "Friday",
		'S': "Saturday",
		'U': "Sunday",
	}

	result := make(map[string]*Analytics.DayofWeek)
	result["Monday"] = &Analytics.DayofWeek{Times: make(map[string]*Analytics.Location)}
	result["Tuesday"] = &Analytics.DayofWeek{Times: make(map[string]*Analytics.Location)}
	result["Wednesday"] = &Analytics.DayofWeek{Times: make(map[string]*Analytics.Location)}
	result["Thursday"] = &Analytics.DayofWeek{Times: make(map[string]*Analytics.Location)}
	result["Friday"] = &Analytics.DayofWeek{Times: make(map[string]*Analytics.Location)}
	result["Saturday"] = &Analytics.DayofWeek{Times: make(map[string]*Analytics.Location)}
	result["Sunday"] = &Analytics.DayofWeek{Times: make(map[string]*Analytics.Location)}
	result["By Arrangement"] = &Analytics.DayofWeek{Times: make(map[string]*Analytics.Location)}
	result["By Arrangement"].Times["By Arrangement"] = &Analytics.Location{Data: make(map[string]int64)}
	registrations := []models.CourseRegistration{}
	models.DB.Find(&registrations)
	// fmt.Println(registrations)
	for _, reg := range registrations {
		socEntry := models.Soc{}
		models.DB.Where("index = ?", reg.ClassIndex).First(&socEntry)
		// fmt.Println(socEntry)
		if socEntry.MeetingTimes == "By Arrangement" || socEntry.MeetingTimes == "" {
			result["By Arrangement"].Times["By Arrangement"].Data["By Arrangement"]++
		} else {
			times := strings.Split(socEntry.MeetingTimes, "|")
			locations := strings.Split(socEntry.MeetingLocation, "|")
			// fmt.Println(times)
			// fmt.Println(locations)
			if len(locations) != len(times) {
				for index := 0; index < len(times); {
					time := times[index]
					timestring := "By Arrangement"
					date := "By Arrangement"
					if time != "By Arrangement" {
						date = dayLookup[time[0]]
						if strings.Index(time, "-") == -1 {
							timestring = time[1:]
						} else {
							timestring = time[1:strings.Index(time, "-")]
						}
					}
					location := locations[0]
					latlong := "N/A"
					if location != "N/A" {
						entry := models.BuildingLookup{}
						models.DB.Where("Code = ?", location[0:strings.Index(location, " ")]).First(&entry)
						if entry.Lat != "" {
							latlong = entry.Lat + "|" + entry.Long
						}
					}

					if _, found := result[date].Times[timestring]; found == true {
						result[date].Times[timestring].Data[latlong]++
					} else {
						result[date].Times[timestring] = &Analytics.Location{Data: make(map[string]int64)}
						result[date].Times[timestring].Data[latlong] = 1
					}
					index++
				}
			} else {
				for index := 0; index < len(times); {
					time := times[index]
					timestring := "By Arrangement"
					date := "By Arrangement"
					if time != "By Arrangement" {
						date = dayLookup[time[0]]
						if strings.Index(time, "-") == -1 {
							timestring = time[1:]
						} else {
							timestring = time[1:strings.Index(time, "-")]
						}
					}
					location := locations[index]
					latlong := "N/A"
					if location != "N/A" {
						entry := models.BuildingLookup{}
						models.DB.Where("Code = ?", location[0:strings.Index(location, " ")]).First(&entry)
						if entry.Lat != "" {
							latlong = entry.Lat + "|" + entry.Long
						}
					}

					if _, found := result[date].Times[timestring]; found == true {
						result[date].Times[timestring].Data[latlong]++
					} else {
						result[date].Times[timestring] = &Analytics.Location{Data: make(map[string]int64)}
						result[date].Times[timestring].Data[latlong] = 1
					}
					index++
				}

			}
		}
	}
	return &Analytics.Heatmap{Days: result}, nil
}

func main() {
	debugPrnt := flag.Bool("debug", false, "Debug Print all Requests")
	flag.Parse()
	debug = *debugPrnt

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Failed to listen on port 8080: ", err)
	}

	s := NewServer()
	grpcServer := grpc.NewServer()
	Analytics.RegisterAnalyticsEndpointServer(grpcServer, s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to listen on port 8080: ", err)
	}
}
