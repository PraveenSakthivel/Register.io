package dbRequests

import (
	"golang.org/x/net/context"
	"main/models"
	"database/sql"
	"log"
)

type Server struct {
	Db *sql.DB
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

func ConvertClasses(classesList []models.Soc) ([]*Class, error) {
	
	var convertedList []*Class

	for _, soc_item := range classesList {
		var available = false
		if soc_item.Spots > 0 {
			available = true
		}
		var temp = Class{Location: soc_item.Location, Level: soc_item.Level, School: int32(soc_item.School),
		Department: int32(soc_item.Department), ClassNum: int32(soc_item.ClassNumber), Index: soc_item.Index, 
		Name: soc_item.Name, Section: soc_item.Section, MeetingLocation: soc_item.MeetingLocation, MeetingTimes: soc_item.MeetingTimes,
		Exam: soc_item.Exam, Instructors: soc_item.Instructors, Codes: soc_item.Codes, Synopsis: soc_item.Synopsis, 
		Books: soc_item.Books, Available: available}

		convertedList = append(convertedList, &temp)
	}

	return convertedList, nil
}