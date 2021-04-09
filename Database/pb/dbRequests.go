package dbRequests

import (
	"golang.org/x/net/context"
	"registerio/db/models"
	"log"
	data "registerio/db/database"
	"strings"
)

type Server struct {
	Db *data.DB
	UnimplementedDatabseWrapperServer 
	Debug bool
}

func dprint(s *Server, msg ...interface{}) {
	if s.Debug {
		log.Println(msg...)
	}
}

func (s *Server) RetrieveClasses(ctx context.Context, input *ReceiveClassesParams) (*ClassesResponse, error) {
	dprint(s, "REQUEST: Retrieve all Classes")
	classes, err:= models.RetrieveAllClasses(s.Db)
	if err != nil {
		return nil, err
	}
	dprint(s, "OK: Successfully retrieved all classes")
	var resp ClassesResponse
	resp.Classes, err = ConvertClasses(classes)
	if err != nil {
		return nil, err
	}
	dprint(s, "OK: Successfully serialized all classes")
	return &resp, nil
}

func ConvertClasses(classesMap map[string][]models.Soc) ([]*Class, error) {
	
	var convertedList []*Class

	for _, soc_list := range classesMap {
		var sectionList []*Section
		var individual_class models.Soc
		for _, soc_item := range soc_list {
			var available = false
			if soc_item.Spots > 0 {
				available = true
			}
			var meetings []*Meeting
			split_times := strings.Split(soc_item.MeetingTimes, "|")
			split_locations := strings.Split(soc_item.MeetingLocation, "|")
			split_campus := strings.Split(soc_item.Location, "|")
			for i := range split_times {
				var time string
				var location string
				var campus string
				if(len(split_times) == 1) {
					time = split_times[0]
				}else {
					time = split_times[i]
				}

				if(len(split_locations) == 1) {
					location = split_locations[0]
				}else {
					location = split_locations[i]
				}

				if(len(split_campus) == 1) {
					campus = split_campus[0]
				}else {
					campus = split_campus[i]
				}
				temp := Meeting{MeetingTime: time, MeetingLocation: location, Campus: campus}
				meetings = append(meetings, &temp)
			}
			temp := Section{Index: soc_item.Index, Section: soc_item.Section, Meetings: meetings, 
			Instructors: soc_item.Instructors, Available: available, Exam: soc_item.Exam}
			sectionList = append(sectionList, &temp)
			individual_class = soc_item
		}
		temp := Class{Level: individual_class.Level, School: int32(individual_class.School), Department: int32(individual_class.Department),
		ClassNum: int32(individual_class.ClassNumber), Name: individual_class.Name, Codes: individual_class.Codes, Synopsis: individual_class.Synopsis, 
		Books: individual_class.Books, Sections: sectionList}

		convertedList = append(convertedList, &temp)
	}

	return convertedList, nil
}