package services

import (
	"fmt"
	"strconv"
	"testing"
)

var addID int64
var addID2 int64
var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJhY2Nlc3MiLCJncmFudCI6ImNsaWVudF9jcmVkZW50aWFscyIsImNsaWVudElkIjo0MDMsInJvbGVVcmlzIjpbeyJjbGllbnRSb2xlSWQiOjEsInJvbGUiOiJhZG1pbiIsInVyaUlkIjoyNzMsInVyaSI6Ii9ycy9pbWFnZS9hZGQiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6MSwicm9sZSI6ImFkbWluIiwidXJpSWQiOjI2MSwidXJpIjoiL3JzL29yZGVyL3VwZGF0ZSIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoxLCJyb2xlIjoiYWRtaW4iLCJ1cmlJZCI6MjQyLCJ1cmkiOiIvcnMvb3B0aW9ucy9kZWxldGUiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6MSwicm9sZSI6ImFkbWluIiwidXJpSWQiOjIzMiwidXJpIjoiL3JzL2N1c3RvbWVyL2RlbGV0ZSIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoxLCJyb2xlIjoiYWRtaW4iLCJ1cmlJZCI6MjAwLCJ1cmkiOiIvcnMvcm9sZS9saXN0IiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjEsInJvbGUiOiJhZG1pbiIsInVyaUlkIjoxNjgsInVyaSI6Ii9ycy91c2VyL2FkZCIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoxLCJyb2xlIjoiYWRtaW4iLCJ1cmlJZCI6MjY1LCJ1cmkiOiIvcnMvb3JkZXIvZGVsZXRlIiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjEsInJvbGUiOiJhZG1pbiIsInVyaUlkIjoyNDMsInVyaSI6Ii9ycy9kZXRhaWxzL2FkZCIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoxLCJyb2xlIjoiYWRtaW4iLCJ1cmlJZCI6MjMzLCJ1cmkiOiIvcnMvcHJvZHVjdC9hZGQiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6MSwicm9sZSI6ImFkbWluIiwidXJpSWQiOjIwMSwidXJpIjoiL3JzL21haWwvc2VuZCIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoxLCJyb2xlIjoiYWRtaW4iLCJ1cmlJZCI6MTkyLCJ1cmkiOiIvcnMvdXNlci91cGRhdGUiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6MSwicm9sZSI6ImFkbWluIiwidXJpSWQiOjI2NiwidXJpIjoiL3JzL29yZGVyL2l0ZW0vYWRkIiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjEsInJvbGUiOiJhZG1pbiIsInVyaUlkIjoyNDQsInVyaSI6Ii9ycy9kZXRhaWxzL3VwZGF0ZSIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoxLCJyb2xlIjoiYWRtaW4iLCJ1cmlJZCI6Mjc0LCJ1cmkiOiIvcnMvaW1hZ2UvdXBkYXRlIiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjEsInJvbGUiOiJhZG1pbiIsInVyaUlkIjoyMjMsInVyaSI6Ii9ycy9hZGRyZXNzL2FkZCIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoxLCJyb2xlIjoiYWRtaW4iLCJ1cmlJZCI6MTkzLCJ1cmkiOiIvcnMvdXNlci9nZXQiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6MSwicm9sZSI6ImFkbWluIiwidXJpSWQiOjI0OSwidXJpIjoiL3JzL2RldGFpbHMvZGVsZXRlIiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjEsInJvbGUiOiJhZG1pbiIsInVyaUlkIjoyNjcsInVyaSI6Ii9ycy9vcmRlci9pdGVtL3VwZGF0ZSIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoxLCJyb2xlIjoiYWRtaW4iLCJ1cmlJZCI6Mjc4LCJ1cmkiOiIvcnMvaW1hZ2UvZGVsZXRlIiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjEsInJvbGUiOiJhZG1pbiIsInVyaUlkIjoyMjQsInVyaSI6Ii9ycy9hZGRyZXNzL3VwZGF0ZSIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoxLCJyb2xlIjoiYWRtaW4iLCJ1cmlJZCI6MTk0LCJ1cmkiOiIvcnMvdXNlci9kZWxldGUiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6MSwicm9sZSI6ImFkbWluIiwidXJpSWQiOjI1NSwidXJpIjoiL3JzL2JhckNvZGUvYWRkIiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjEsInJvbGUiOiJhZG1pbiIsInVyaUlkIjoyMzQsInVyaSI6Ii9ycy9wcm9kdWN0L3VwZGF0ZSIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoxLCJyb2xlIjoiYWRtaW4iLCJ1cmlJZCI6MjY4LCJ1cmkiOiIvcnMvb3JkZXIvaXRlbS9kZWxldGUiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6MSwicm9sZSI6ImFkbWluIiwidXJpSWQiOjE5NSwidXJpIjoiL3JzL3VzZXIvbGlzdCIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoxLCJyb2xlIjoiYWRtaW4iLCJ1cmlJZCI6Mjc5LCJ1cmkiOiIvcnMvY29udGVudC9hZGQiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6MSwicm9sZSI6ImFkbWluIiwidXJpSWQiOjIyNywidXJpIjoiL3JzL2FkZHJlc3MvZGVsZXRlIiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjEsInJvbGUiOiJhZG1pbiIsInVyaUlkIjoyNjksInVyaSI6Ii9ycy9vcmRlci9wYWNrYWdlL2FkZCIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoxLCJyb2xlIjoiYWRtaW4iLCJ1cmlJZCI6MjU2LCJ1cmkiOiIvcnMvYmFyQ29kZS91cGRhdGUiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6MSwicm9sZSI6ImFkbWluIiwidXJpSWQiOjIzNiwidXJpIjoiL3JzL3Byb2R1Y3QvZGVsZXRlIiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjEsInJvbGUiOiJhZG1pbiIsInVyaUlkIjoxOTYsInVyaSI6Ii9ycy9yb2xlL2FkZCIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoxLCJyb2xlIjoiYWRtaW4iLCJ1cmlJZCI6NjIsInVyaSI6Imh0dHA6Ly9sb2NhbGhvc3QvcnMvYWRkQ2xpZW50IiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjEsInJvbGUiOiJhZG1pbiIsInVyaUlkIjoyODAsInVyaSI6Ii9ycy9jb250ZW50L3VwZGF0ZSIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoxLCJyb2xlIjoiYWRtaW4iLCJ1cmlJZCI6MjI4LCJ1cmkiOiIvcnMvY3VzdG9tZXIvYWRkIiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjEsInJvbGUiOiJhZG1pbiIsInVyaUlkIjoyNzAsInVyaSI6Ii9ycy9vcmRlci9wYWNrYWdlL3VwZGF0ZSIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoxLCJyb2xlIjoiYWRtaW4iLCJ1cmlJZCI6MjU5LCJ1cmkiOiIvcnMvYmFyQ29kZS9kZWxldGUiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6MSwicm9sZSI6ImFkbWluIiwidXJpSWQiOjIzNywidXJpIjoiL3JzL29wdGlvbnMvYWRkIiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjEsInJvbGUiOiJhZG1pbiIsInVyaUlkIjoyMjksInVyaSI6Ii9ycy9jdXN0b21lci91cGRhdGUiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6MSwicm9sZSI6ImFkbWluIiwidXJpSWQiOjE5OCwidXJpIjoiL3JzL3JvbGUvZ2V0IiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjEsInJvbGUiOiJhZG1pbiIsInVyaUlkIjo2MywidXJpIjoiaHR0cDovL2xvY2FsaG9zdC9ycy91cGRhdGVDbGllbnQiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6MSwicm9sZSI6ImFkbWluIiwidXJpSWQiOjI4MywidXJpIjoiL3JzL2NvbnRlbnQvZGVsZXRlIiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjEsInJvbGUiOiJhZG1pbiIsInVyaUlkIjoyNzIsInVyaSI6Ii9ycy9vcmRlci9wYWNrYWdlL2RlbGV0ZSIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoxLCJyb2xlIjoiYWRtaW4iLCJ1cmlJZCI6MjYwLCJ1cmkiOiIvcnMvb3JkZXIvYWRkIiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjEsInJvbGUiOiJhZG1pbiIsInVyaUlkIjoyMzgsInVyaSI6Ii9ycy9vcHRpb25zL3VwZGF0ZSIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoxLCJyb2xlIjoiYWRtaW4iLCJ1cmlJZCI6MjMxLCJ1cmkiOiIvcnMvY3VzdG9tZXIvbGlzdCIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoxLCJyb2xlIjoiYWRtaW4iLCJ1cmlJZCI6MTk5LCJ1cmkiOiIvcnMvcm9sZS9kZWxldGUiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6MSwicm9sZSI6ImFkbWluIiwidXJpSWQiOjc3LCJ1cmkiOiJodHRwOi8vbG9jYWxob3N0L3JzL2FkZENsaWVudFNjb3BlIiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjEsInJvbGUiOiJhZG1pbiIsInVyaUlkIjozMzAsInVyaSI6Ii9ycy9jb250ZW50L2hpdHMiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6Miwicm9sZSI6InVzZXIiLCJ1cmlJZCI6Mjc3LCJ1cmkiOiIvcnMvaW1hZ2UvbGlzdCIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoyLCJyb2xlIjoidXNlciIsInVyaUlkIjoyNDAsInVyaSI6Ii9ycy9vcHRpb25zL2dldEJ5RGV0YWlscyIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoyLCJyb2xlIjoidXNlciIsInVyaUlkIjo2OCwidXJpIjoiaHR0cDovL2xvY2FsaG9zdC9ycy9kZWxldGVDbGllbnRBbGxvd2VkVXJpIiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjIsInJvbGUiOiJ1c2VyIiwidXJpSWQiOjI4MSwidXJpIjoiL3JzL2NvbnRlbnQvZ2V0IiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjIsInJvbGUiOiJ1c2VyIiwidXJpSWQiOjI1OCwidXJpIjoiL3JzL2JhckNvZGUvZ2V0QnlEZXRhaWxzIiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjIsInJvbGUiOiJ1c2VyIiwidXJpSWQiOjgwLCJ1cmkiOiJodHRwOi8vbG9jYWxob3N0L3JzL2FkZENsaWVudFJvbGVVcmkiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6Miwicm9sZSI6InVzZXIiLCJ1cmlJZCI6MjQxLCJ1cmkiOiIvcnMvb3B0aW9ucy9zZWFyY2hCeU9wdGlvbiIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoyLCJyb2xlIjoidXNlciIsInVyaUlkIjoyODIsInVyaSI6Ii9ycy9jb250ZW50L2xpc3QiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6Miwicm9sZSI6InVzZXIiLCJ1cmlJZCI6MjYyLCJ1cmkiOiIvcnMvb3JkZXIvZ2V0IiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjIsInJvbGUiOiJ1c2VyIiwidXJpSWQiOjIyNSwidXJpIjoiL3JzL2FkZHJlc3MvZ2V0IiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjIsInJvbGUiOiJ1c2VyIiwidXJpSWQiOjI0NSwidXJpIjoiL3JzL2RldGFpbHMvZ2V0IiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjIsInJvbGUiOiJ1c2VyIiwidXJpSWQiOjI2MywidXJpIjoiL3JzL29yZGVyL2xpc3QiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6Miwicm9sZSI6InVzZXIiLCJ1cmlJZCI6MjI2LCJ1cmkiOiIvcnMvYWRkcmVzcy9saXN0IiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjIsInJvbGUiOiJ1c2VyIiwidXJpSWQiOjI0NiwidXJpIjoiL3JzL2RldGFpbHMvZ2V0QnlQcm9kdWN0IiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjIsInJvbGUiOiJ1c2VyIiwidXJpSWQiOjI2NCwidXJpIjoiL3JzL29yZGVyL2N1c3RvbWVyL2xpc3QiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6Miwicm9sZSI6InVzZXIiLCJ1cmlJZCI6MjQ3LCJ1cmkiOiIvcnMvZGV0YWlscy9nZXRCeVNrdSIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoyLCJyb2xlIjoidXNlciIsInVyaUlkIjoyMzAsInVyaSI6Ii9ycy9jdXN0b21lci9nZXQiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6Miwicm9sZSI6InVzZXIiLCJ1cmlJZCI6MjcxLCJ1cmkiOiIvcnMvb3JkZXIvcGFja2FnZS9nZXQiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6Miwicm9sZSI6InVzZXIiLCJ1cmlJZCI6MjQ4LCJ1cmkiOiIvcnMvZGV0YWlscy9nZXRCeUJhckNvZGUiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6Miwicm9sZSI6InVzZXIiLCJ1cmlJZCI6MjMxLCJ1cmkiOiIvcnMvY3VzdG9tZXIvbGlzdCIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoyLCJyb2xlIjoidXNlciIsInVyaUlkIjoyNzUsInVyaSI6Ii9ycy9pbWFnZS9kZXRhaWxzIiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjIsInJvbGUiOiJ1c2VyIiwidXJpSWQiOjI1NywidXJpIjoiL3JzL2JhckNvZGUvZ2V0IiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjIsInJvbGUiOiJ1c2VyIiwidXJpSWQiOjIzNSwidXJpIjoiL3JzL3Byb2R1Y3QvZ2V0IiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjIsInJvbGUiOiJ1c2VyIiwidXJpSWQiOjI3NiwidXJpIjoiL3JzL2ltYWdlL3BhZ2UvY291bnQiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6Miwicm9sZSI6InVzZXIiLCJ1cmlJZCI6MjM5LCJ1cmkiOiIvcnMvb3B0aW9ucy9nZXQiLCJjbGllbnRJZCI6NDAzfV0sImV4cGlyZXNJbiI6MzYwMCwiaWF0IjoxNTAzMTg5NTg3LCJ0b2tlblR5cGUiOiJhY2Nlc3MiLCJleHAiOjE1MDMxOTMxODcsImlzcyI6IlVsYm9yYSBPYXV0aDIgU2VydmVyIn0.sGQWRNfU9Bdq2TYHhqIwFh1n2s1tFRa-Ud2kikbpm4E"

func TestContentService_AddContent(t *testing.T) {
	var c ContentService
	c.ClientID = "403"
	c.Host = "http://localhost:3008"
	c.Token = token
	var ct Content
	ct.Category = "books"
	ct.MetaAuthorName = "ken"
	ct.MetaDesc = "a book"
	ct.Text = "some book text"
	ct.Title = "the best book ever"
	res := c.AddContent(&ct)
	fmt.Print("res: ")
	fmt.Println(res)
	addID = res.ID
	if res.Success != true {
		t.Fail()
	}
}

func TestContentService_UpdateContent(t *testing.T) {
	var c ContentService
	c.ClientID = "403"
	c.Host = "http://localhost:3008"
	c.Token = token
	var ct Content
	ct.ID = addID
	ct.Category = "music"
	ct.MetaAuthorName = "ken"
	ct.MetaDesc = "a song"
	ct.Text = "some music text"
	ct.Title = "the best song ever"
	res := c.UpdateContent(&ct)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestContentService_UpdateContentHits(t *testing.T) {
	var c ContentService
	c.ClientID = "403"
	c.Host = "http://localhost:3008"
	c.Token = token
	var ct Content
	ct.ID = addID
	ct.Hits = 500
	res := c.UpdateContentHits(&ct)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestContentService_AddAnotherContent(t *testing.T) {
	var c ContentService
	c.ClientID = "403"
	c.Host = "http://localhost:3008"
	c.Token = token
	var ct Content
	ct.Category = "books"
	ct.MetaAuthorName = "ken"
	ct.MetaDesc = "a book"
	ct.Text = "some book text"
	ct.Title = "the best book ever"
	res := c.AddContent(&ct)
	fmt.Print("res: ")
	fmt.Println(res)
	addID2 = res.ID
	if res.Success != true {
		t.Fail()
	}
}

func TestContentService_GetContent(t *testing.T) {
	var c ContentService
	c.Host = "http://localhost:3008"
	res := c.GetContent(strconv.FormatInt(addID, 10), "403")
	fmt.Print("res: ")
	fmt.Println(res)
	if res.ID != addID && res.Hits == 500 {
		t.Fail()
	}
}

func TestContentService_GetContentList(t *testing.T) {
	var c ContentService
	c.Host = "http://localhost:3008"
	res := c.GetContentList("403")
	fmt.Print("res list: ")
	fmt.Println(res)
	fmt.Print("len: ")
	fmt.Println(len(*res))
	if res == nil || len(*res) == 0 {
		t.Fail()
	}
}

func TestContentService_DeleteContent(t *testing.T) {
	var c ContentService
	c.ClientID = "403"
	c.Host = "http://localhost:3008"
	c.Token = token

	res := c.DeleteContent(strconv.FormatInt(addID, 10))
	fmt.Print("res deleted: ")
	fmt.Println(res)
	addID = res.ID
	if res.Success != true {
		t.Fail()
	}
}

func TestContentService_DeleteContent2(t *testing.T) {
	var c ContentService
	c.ClientID = "403"
	c.Host = "http://localhost:3008"
	c.Token = token

	res := c.DeleteContent(strconv.FormatInt(addID2, 10))
	fmt.Print("res deleted: ")
	fmt.Println(res)
	addID = res.ID
	if res.Success != true {
		t.Fail()
	}
}
