package classtiming

import (
	"errors"
	"log"
	"strings"
	"time"
)

type ClassSlot struct {
	startTime time.Time
	endTime   time.Time
	location  string
	day       time.Weekday
	next      *ClassSlot
}

func FormatTime(s string) (time.Time, error) {
	return time.Parse("03:04 PM", s)
}

func CreateClassSlot(startTime string, endTime string, location string, day time.Weekday, next *ClassSlot) (ClassSlot, error) {
	start, err := FormatTime(startTime)
	if err != nil {
		return ClassSlot{}, err
	}
	end, err := FormatTime(endTime)
	if err != nil {
		return ClassSlot{}, err
	}

	return ClassSlot{startTime: start, endTime: end, location: location, day: day, next: next}, nil

}

func BuildClassSlots(times string, location string) ([]ClassSlot, error) {
	var timesArr = strings.Split(times, "|")
	var slotsArr []ClassSlot
	for _, timestr := range timesArr {
		// Strings should be formatted as DHH:MM-HH:MM
		// D = weekday character, H = hour, M = minute
		var dayOfWeek = timestr[0]
		var startTime = timestr[1:9]
		var endTime = timestr[10:]

		var wd time.Weekday

		switch dayOfWeek {
		case 'U':
			wd = 0
		case 'M':
			wd = 1
		case 'T':
			wd = 2
		case 'W':
			wd = 3
		case 'R':
			wd = 4
		case 'F':
			wd = 5
		case 'S':
			wd = 6
		default:
			return nil, errors.New("Invalid day of Week: " + string([]byte{dayOfWeek}))
		}

		var classSlot, err = CreateClassSlot(startTime, endTime, location, wd, nil)
		if err != nil {
			return nil, err
		}
		slotsArr = append(slotsArr, classSlot)
	}
	return slotsArr, nil
}

func CheckTime(course_times map[time.Weekday]*ClassSlot, classToAdd *ClassSlot) (bool, error) {
	var startTime = classToAdd.startTime
	var endTime = classToAdd.endTime
	var location = classToAdd.location
	var wd = classToAdd.day
	var front = course_times[wd]
	if front == nil {
		return true, nil
	}
	var ptr = front
	var prev *ClassSlot
	for ptr.next != nil {
		if !startTime.Before(ptr.startTime) {
			break
		}
		prev = ptr
		ptr = ptr.next
	}

	if ptr.next == nil && !startTime.Before(ptr.startTime) {
		return false, nil
	}

	// checking the end to make sure it follows rules
	log.Println("(just to avoid errorrs) Checking location: " + location)
	// include something that subtracts time based on location, need to read rules for that, for now, just setting default to 20 mins

	var minsApart = 20
	if prev == nil {
		prev = ptr
	}
	var differenceInTime = prev.startTime.Sub(endTime).Minutes()

	if int(differenceInTime) < minsApart {
		return false, nil
	}

	return true, nil
}

func InsertTime(course_times map[time.Weekday]*ClassSlot, classToAdd *ClassSlot) (map[time.Weekday]*ClassSlot, error) {
	// Assumes that it is already able to fit, does not check for that
	var wd = classToAdd.day
	var ptr = course_times[wd]
	var startTime = classToAdd.startTime
	var prev *ClassSlot
	if ptr == nil {
		course_times[wd] = classToAdd
		return course_times, nil
	}
	for ptr.next != nil {
		if !startTime.Before(ptr.next.startTime) {
			break
		}
		prev = ptr
		ptr = ptr.next
	}

	if prev == nil {
		classToAdd.next = ptr
		course_times[wd] = classToAdd
		return course_times, nil
	}

	prev.next = classToAdd
	classToAdd.next = ptr
	return course_times, nil
}

func CheckTimesAndInsert(slotsArr []ClassSlot, course_times map[time.Weekday]*ClassSlot) (map[time.Weekday]*ClassSlot, error) {
	for _, slot := range slotsArr {
		ans, err := CheckTime(course_times, &slot)
		if err != nil {
			return nil, err
		}
		if ans == false {
			return nil, nil
		}
		temp, err := InsertTime(course_times, &slot)
		if err != nil {
			return nil, nil
		}
		course_times = temp
	}
	// insert if all are good

	return course_times, nil
}

func printMap(course_times map[time.Weekday]*ClassSlot) {
	for day, x := range course_times {
		log.Println(day)
		ptr := x
		for ptr != nil {
			log.Println(ptr.startTime.Format("03:04 PM") + " - " + ptr.endTime.Format("03:04 PM"))
			ptr = ptr.next
		}
	}

}
  
func BuildSchedule(indices []string, lookupPtr *map[string][]ClassSlot) (map[time.Weekday]*ClassSlot, error)  {
	lookup := *lookupPtr
	courseTimes := make(map[time.Weekday]*ClassSlot)
	for _, index := range indices {
		if _, ok := lookup[index]; !ok {
			log.Println("ERROR: Cannot build student schedule|Lookup")
			return nil, errors.New("Cannot build student schedule")
		}
		class := lookup[index]
		courseTimes, err := CheckTimesAndInsert(class, courseTimes)
		if err != nil && courseTimes != nil {
			log.Println("ERROR: Cannot build student schedule|Build: ", err)
			return nil, err
		}
	}
	return courseTimes, nil
}
