package main

import (
	"context"
	"fmt"

	eventgrid "github.com/Azure/azure-sdk-for-go/profiles/latest/eventgrid/eventgrid"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	uuid "github.com/google/uuid"
)

func main() {
	eventGridKey := "<YOUR KEY HERE>"
	eventGridTopicHostName := "<TOPIC HOSTNAME HERE IN THE FORMAT OF name.region.eventgrid.azure.net>"

	keyAuthorizer := autorest.NewEventGridKeyAuthorizer(eventGridKey)
	eventGridClient := eventgrid.BaseClient{}
	eventGridClient.Authorizer = keyAuthorizer

	id := uuid.New().String()
	subject := "Sample-event"
	eventType := "Demo"
	dataVersion := "v2"

	//All of the fields are mandatory - please see https://docs.microsoft.com/en-us/azure/event-grid/event-schema#event-properties
	events := []eventgrid.Event{
		{
			ID:          &id,
			Subject:     &subject,
			EventType:   &eventType,
			EventTime:   &date.Time{},
			DataVersion: &dataVersion,
			Data:        "Hello World!",
		},
	}

	req, publishPrepareErr := eventGridClient.PublishEventsPreparer(context.Background(), eventGridTopicHostName, events)
	if publishPrepareErr != nil {
		fmt.Println(publishPrepareErr)
	}

	req, prepareErr := autorest.Prepare(req, eventGridClient.WithAuthorization())
	if prepareErr != nil {
		fmt.Println(prepareErr)
	}

	send, sendErr := eventGridClient.PublishEventsSender(req)
	if sendErr != nil {
		fmt.Println(sendErr)
	}

	resp, respondErr := eventGridClient.PublishEventsResponder(send)
	if respondErr != nil {
		fmt.Println(respondErr)
	}

	fmt.Println(resp.Status)
}
