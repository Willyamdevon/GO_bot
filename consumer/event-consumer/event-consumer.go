package event_consumer

import (
	"log"
	"main/events"
	"time"
)

type Consumer struct {
	fetcher   events.Fetcher
	processor events.Processor
	batchSize int
}

func New(fetcher events.Fetcher, processor events.Processor, batchSize int)Consumer{
	return Consumer{
		fetcher: fetcher,
		processor: processor,
		batchSize: batchSize,
	}
}

func (c Consumer) Start() error{
	for {
		gotEvents, err := c.fetcher.Fetch(c.batchSize)
		if err!=nil{
			log.Printf("[ERR] consumer: %s", err.Error())

			continue
		}

		if len(gotEvents) == 0{
			time.Sleep(1 * time.Second)

			continue
		}
		
		if err:=c.handlEvents(gotEvents); err!=nil{
			log.Print(err)

			continue
		}

	}
}

func (c *Consumer) handlEvents(events []events.Event) error{
	for _, event := range events{
		log.Printf("got new event: %s", event.Text)

		if err:=c.processor.Process(event); err!=nil{
			log.Printf("can`t handle event: %s", err.Error())

			continue
		}
	}

	return nil
}