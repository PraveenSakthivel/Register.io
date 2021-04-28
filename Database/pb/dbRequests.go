package dbRequests

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"golang.org/x/net/context"
	"registerio/db/models"
	"log"
	data "registerio/db/database"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"errors"
	"fmt"
	"strconv"
)

type Server struct {
	Db *data.DB
	UnimplementedDatabaseWrapperServer 
	Debug bool
	TokenSecret string
}

type userClaims struct {
	NetID        string           `json:"name"`
	ClassHistory map[string]int32 `json:"classHistory"`
	SpecialCases map[string]bool  `json:"specialCases"`
	jwt.StandardClaims
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

func (s *Server) ReturnDepartments(ctx context.Context, input *ReceiveDepartmentsParams) (*DepartmentsResponse, error) {
	dprint(s, "REQUEST: Retrieve list of departments")
	departments, err:= models.RetrieveDepartments(s.Db)
	if err != nil {
		return nil, err
	}
	dprint(s, "OK: Successfully retrieved all departments")
	var resp DepartmentsResponse
	resp.Departments = departments
	dprint(s, "OK: Serialized department map output")
	return &resp, nil
}

func (s *Server) ClassAddStatus(ctx context.Context, input *ClassAddStatusParams) (*AddStatusResponse, error) {
	dprint(s, "REQUEST: Checking statuses of classes")
	log.Println("Secret: "+ s.TokenSecret)
	netid, err := s.parseJWT(input.Token)
	if err != nil {
		return nil, err
	}
	dprint(s, "OK: Successfully decoded netid: ", netid)
	var resp AddStatusResponse
	resp.Statuses = make(map[string]AddStatus)	
	for _, index := range input.Index {
		status, err := models.CheckStatus(s.Db, netid, index)
		if err != nil {
			return nil, err
		}
		switch status {
		case 0:
			resp.Statuses[index] = AddStatus_PENDING
		case 1:
			resp.Statuses[index] = AddStatus_ADDED
		case 2:
			resp.Statuses[index] = AddStatus_FAILED
		default:
			return nil, errors.New("Could not properly get status for index: "+index+" and netid: "+netid)
		}
	}
	
	dprint(s, "OK: Successfully serialized all statuses")
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
		ClassNum: int32(individual_class.ClassNumber), Credits: int32(individual_class.Credits), Name: individual_class.Name, Codes: individual_class.Codes, Synopsis: individual_class.Synopsis, 
		Books: individual_class.Books, Sections: sectionList}

		convertedList = append(convertedList, &temp)
	}

	return convertedList, nil
}

func (s *Server) Decrypt(encryptedString string) (string, error) {
	// key, _ := hex.DecodeString(tokenObj.Token[0:32])
	enc, _ := hex.DecodeString(encryptedString)

	block, err := aes.NewCipher([]byte(s.TokenSecret[0:32]))
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		log.Print(err.Error())
		return "", err
	}

	nonceSize := aesGCM.NonceSize()

	if len(enc) < nonceSize {
		return "", err
	}

	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Print(err.Error())
		return "", err
	}

	return fmt.Sprintf("%s", plaintext), nil
}

func (s *Server)parseJWT(encodedToken string) (string, error) {
	decodedToken, err := s.Decrypt(encodedToken)
	if err != nil {
		return "", err
	}
	token, err := jwt.ParseWithClaims(decodedToken, &userClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return "", errors.New("Invalid token")
		}
		return []byte(s.TokenSecret), nil
	})

	if err != nil {
		log.Println("Error Parsing JWT: ", err)
		return "", err
	}

	if claims, ok := token.Claims.(*userClaims); ok {
		specCases := make(map[int32]bool)
		for key, val := range claims.SpecialCases {
			intkey, err := strconv.ParseInt(key, 10, 32)
			if err != nil {
				log.Println("Error Parsing Cases: ", err)
				continue
			}
			specCases[int32(intkey)] = val
		}
		return claims.NetID, nil
	}
	return "", errors.New("Unable to Parse JWT")
}
