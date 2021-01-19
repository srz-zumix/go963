package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"google.golang.org/api/calendar/v3"
)

const (
	dateFormat string = "2006-01-02"
)

func getGo963Summary(user string, summary string) string {
	return fmt.Sprintf("%s %s: %s", options.Prefix, user, summary)
}
func getKintaiSummary(summary string) string {
	return strings.TrimSpace(strings.Replace(summary, options.Prefix, "", 1))
}

func makeEvent(summary string, date string, timezone string, description string) *calendar.Event {
	event := &calendar.Event{
		Summary: summary,
		Start: &calendar.EventDateTime{
			Date:     date,
			TimeZone: timezone,
		},
		End: &calendar.EventDateTime{
			Date:     date,
			TimeZone: timezone,
		},
		Description: description,
		ExtendedProperties: &calendar.EventExtendedProperties{
			Shared: map[string]string{
				"go963": "true",
			},
		},
	}
	return event
}

func listEventCall(svc *calendar.Service, minTime time.Time, maxTime time.Time) *calendar.EventsListCall {
	zone, _ := minTime.Zone()
	caller := svc.Events.List(options.CalendarId).TimeMin(minTime.Format(time.RFC3339)).TimeMax(maxTime.Format(time.RFC3339)).TimeZone(zone)
	if options.Strict {
		return caller.SharedExtendedProperty("go963=true")
	}
	return caller
}

func listDayofEventCall(svc *calendar.Service, date time.Time) *calendar.EventsListCall {
	minTime := date
	maxTime := date.AddDate(0, 0, 1)
	return listEventCall(svc, minTime, maxTime)
}

func listEvent(client *http.Client, user string, date time.Time) (*calendar.Events, error) {
	svc, err := calendar.New(client)
	if err != nil {
		log.Fatalf("Unable to create Calendar service: %v", err)
		return nil, err
	}
	events, err := listDayofEventCall(svc, date).Do()
	if err != nil {
		log.Fatalf("Unable to list calendar event: %v", err)
		return nil, err
	}

	var matched_events calendar.Events

	for _, event := range events.Items {
		if strings.HasPrefix(event.Summary, fmt.Sprintf("%s %s", options.Prefix, user)) {
			matched_events.Items = append(matched_events.Items, event)
		}
	}
	events.Items = matched_events.Items
	return events, err
}

func findEvent(svc *calendar.Service, user string, date time.Time) (*calendar.Event, error) {
	events, err := listDayofEventCall(svc, date).Do()
	if err != nil {
		log.Fatalf("Unable to list calendar event: %v", err)
		return nil, err
	}
	for _, event := range events.Items {
		if strings.HasPrefix(event.Summary, fmt.Sprintf("%s %s", options.Prefix, user)) {
			return event, nil
		}
	}
	return nil, nil
}

func setEvent(client *http.Client, user string, date time.Time, ev *calendar.Event) (*calendar.Event, error) {
	svc, err := calendar.New(client)
	if err != nil {
		log.Fatalf("Unable to create Calendar service: %v", err)
		return nil, err
	}
	find_event, err := svc.Events.Get(options.CalendarId, ev.Id).Do()
	if len(find_event.Id) == 0 {
		find_event, err = findEvent(svc, user, date)
	}

	if find_event != nil {
		fmt.Println("update event")
		ev, err = svc.Events.Update(options.CalendarId, find_event.Id, ev).Do()
	} else {
		fmt.Println("create new event")
		ev, err = svc.Events.Insert(options.CalendarId, ev).Do()
	}
	if err != nil {
		log.Fatalf("Failed set calendar: %v", err)
	}
	return ev, err
}

func createEventFromJson(client *http.Client, event_json_file string) ([]*calendar.Event, error) {
	eventJson, err := ioutil.ReadFile(event_json_file)
	if err != nil {
		log.Fatalf("Unable to read event json file: ", err)
		return nil, err
	}

	svc, err := calendar.New(client)
	if err != nil {
		log.Fatalf("Unable to create Calendar service: %v", err)
		return nil, err
	}

	// event := &calendar.Event{
	// 	Summary: "test",
	// 	Start: &calendar.EventDateTime{
	// 		DateTime: time.Now().String(),
	// 		TimeZone: "Asia/Tokyo",
	// 	},
	// 	End: &calendar.EventDateTime{
	// 		DateTime: time.Now().String(),
	// 		TimeZone: "Asia/Tokyo",
	// 	},
	// }

	var events []*calendar.Event
	json.Unmarshal(eventJson, &events)

	for _, event := range events {
		event, err = svc.Events.Insert(options.CalendarId, event).Do()
	}
	return events, err
}

func deleteEvent(client *http.Client, user string, date time.Time) error {
	svc, err := createService(client)
	if err != nil {
		return err
	}
	ev, err := findEvent(svc, user, date)
	if err != nil {
		return err
	}
	if ev == nil {
		return errors.New("event not found")
	}
	return svc.Events.Delete(options.CalendarId, ev.Id).Do()
}

func createService(client *http.Client) (*calendar.Service, error) {
	svc, err := calendar.New(client)
	if err != nil {
		log.Fatalf("Unable to create Calendar service: %v", err)
		return nil, err
	}
	return svc, nil
}
