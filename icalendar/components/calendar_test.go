package components

import (
	"testing"

	"github.com/b3rsrk/caldav-go/icalendar"
	. "github.com/b3rsrk/check"
)

type CalendarSuite struct{ calendar Calendar }

var _ = Suite(new(CalendarSuite))

func TestCalendar(t *testing.T) { TestingT(t) }

// tests the current server for CalDAV support
func (s *CalendarSuite) TestMarshal(c *C) {
	enc, err := icalendar.Marshal(s.calendar)
	c.Assert(err, IsNil)
	c.Assert(enc, Equals, "BEGIN:VCALENDAR\r\nVERSION:2.0\r\nPRODID:-//b3rsrk/caldav-go//NONSGML v1.0.0//EN\r\nEND:VCALENDAR")
}
