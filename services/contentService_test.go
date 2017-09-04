package services

import (
	"fmt"
	"strconv"
	"testing"
)

var addID int64
var addID2 int64
var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJhY2Nlc3MiLCJncmFudCI6ImNsaWVudF9jcmVkZW50aWFscyIsImNsaWVudElkIjo0MDMsInJvbGVVcmlzIjpbeyJjbGllbnRSb2xlSWQiOjEsInJvbGUiOiJhZG1pbiIsInVyaUlkIjozMzEsInVyaSI6Ii9ycy9tYWlsU2VydmVyL2FkZCIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoxLCJyb2xlIjoiYWRtaW4iLCJ1cmlJZCI6MjMyLCJ1cmkiOiIvcnMvY3VzdG9tZXIvZGVsZXRlIiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjEsInJvbGUiOiJhZG1pbiIsInVyaUlkIjoyMDAsInVyaSI6Ii9ycy9yb2xlL2xpc3QiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6MSwicm9sZSI6ImFkbWluIiwidXJpSWQiOjE2OCwidXJpIjoiL3JzL3VzZXIvYWRkIiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjEsInJvbGUiOiJhZG1pbiIsInVyaUlkIjoyNDMsInVyaSI6Ii9ycy9kZXRhaWxzL2FkZCIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoxLCJyb2xlIjoiYWRtaW4iLCJ1cmlJZCI6MjY1LCJ1cmkiOiIvcnMvb3JkZXIvZGVsZXRlIiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjEsInJvbGUiOiJhZG1pbiIsInVyaUlkIjoxOTIsInVyaSI6Ii9ycy91c2VyL3VwZGF0ZSIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoxLCJyb2xlIjoiYWRtaW4iLCJ1cmlJZCI6MzMyLCJ1cmkiOiIvcnMvbWFpbFNlcnZlci91cGRhdGUiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6MSwicm9sZSI6ImFkbWluIiwidXJpSWQiOjIzMywidXJpIjoiL3JzL3Byb2R1Y3QvYWRkIiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjEsInJvbGUiOiJhZG1pbiIsInVyaUlkIjoyMDEsInVyaSI6Ii9ycy9tYWlsL3NlbmQiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6MSwicm9sZSI6ImFkbWluIiwidXJpSWQiOjI0NCwidXJpIjoiL3JzL2RldGFpbHMvdXBkYXRlIiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjEsInJvbGUiOiJhZG1pbiIsInVyaUlkIjoyNjYsInVyaSI6Ii9ycy9vcmRlci9pdGVtL2FkZCIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoxLCJyb2xlIjoiYWRtaW4iLCJ1cmlJZCI6MTkzLCJ1cmkiOiIvcnMvdXNlci9nZXQiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6MSwicm9sZSI6ImFkbWluIiwidXJpSWQiOjMzMywidXJpIjoiL3JzL21haWxTZXJ2ZXIvZ2V0IiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjEsInJvbGUiOiJhZG1pbiIsInVyaUlkIjoyNzQsInVyaSI6Ii9ycy9pbWFnZS91cGRhdGUiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6MSwicm9sZSI6ImFkbWluIiwidXJpSWQiOjIyMywidXJpIjoiL3JzL2FkZHJlc3MvYWRkIiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjEsInJvbGUiOiJhZG1pbiIsInVyaUlkIjoyNjcsInVyaSI6Ii9ycy9vcmRlci9pdGVtL3VwZGF0ZSIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoxLCJyb2xlIjoiYWRtaW4iLCJ1cmlJZCI6MjQ5LCJ1cmkiOiIvcnMvZGV0YWlscy9kZWxldGUiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6MSwicm9sZSI6ImFkbWluIiwidXJpSWQiOjE5NCwidXJpIjoiL3JzL3VzZXIvZGVsZXRlIiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjEsInJvbGUiOiJhZG1pbiIsInVyaUlkIjoyNzgsInVyaSI6Ii9ycy9pbWFnZS9kZWxldGUiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6MSwicm9sZSI6ImFkbWluIiwidXJpSWQiOjIyNCwidXJpIjoiL3JzL2FkZHJlc3MvdXBkYXRlIiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjEsInJvbGUiOiJhZG1pbiIsInVyaUlkIjoyNjgsInVyaSI6Ii9ycy9vcmRlci9pdGVtL2RlbGV0ZSIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoxLCJyb2xlIjoiYWRtaW4iLCJ1cmlJZCI6MjU1LCJ1cmkiOiIvcnMvYmFyQ29kZS9hZGQiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6MSwicm9sZSI6ImFkbWluIiwidXJpSWQiOjIzNCwidXJpIjoiL3JzL3Byb2R1Y3QvdXBkYXRlIiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjEsInJvbGUiOiJhZG1pbiIsInVyaUlkIjoyMjcsInVyaSI6Ii9ycy9hZGRyZXNzL2RlbGV0ZSIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoxLCJyb2xlIjoiYWRtaW4iLCJ1cmlJZCI6MTk1LCJ1cmkiOiIvcnMvdXNlci9saXN0IiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjEsInJvbGUiOiJhZG1pbiIsInVyaUlkIjoyNzksInVyaSI6Ii9ycy9jb250ZW50L2FkZCIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoxLCJyb2xlIjoiYWRtaW4iLCJ1cmlJZCI6MjY5LCJ1cmkiOiIvcnMvb3JkZXIvcGFja2FnZS9hZGQiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6MSwicm9sZSI6ImFkbWluIiwidXJpSWQiOjI1NiwidXJpIjoiL3JzL2JhckNvZGUvdXBkYXRlIiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjEsInJvbGUiOiJhZG1pbiIsInVyaUlkIjoyMzYsInVyaSI6Ii9ycy9wcm9kdWN0L2RlbGV0ZSIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoxLCJyb2xlIjoiYWRtaW4iLCJ1cmlJZCI6MjI4LCJ1cmkiOiIvcnMvY3VzdG9tZXIvYWRkIiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjEsInJvbGUiOiJhZG1pbiIsInVyaUlkIjoxOTYsInVyaSI6Ii9ycy9yb2xlL2FkZCIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoxLCJyb2xlIjoiYWRtaW4iLCJ1cmlJZCI6NjIsInVyaSI6Imh0dHA6Ly9sb2NhbGhvc3QvcnMvYWRkQ2xpZW50IiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjEsInJvbGUiOiJhZG1pbiIsInVyaUlkIjoyODAsInVyaSI6Ii9ycy9jb250ZW50L3VwZGF0ZSIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoxLCJyb2xlIjoiYWRtaW4iLCJ1cmlJZCI6MjcwLCJ1cmkiOiIvcnMvb3JkZXIvcGFja2FnZS91cGRhdGUiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6MSwicm9sZSI6ImFkbWluIiwidXJpSWQiOjI1OSwidXJpIjoiL3JzL2JhckNvZGUvZGVsZXRlIiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjEsInJvbGUiOiJhZG1pbiIsInVyaUlkIjoyMzcsInVyaSI6Ii9ycy9vcHRpb25zL2FkZCIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoxLCJyb2xlIjoiYWRtaW4iLCJ1cmlJZCI6MjgzLCJ1cmkiOiIvcnMvY29udGVudC9kZWxldGUiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6MSwicm9sZSI6ImFkbWluIiwidXJpSWQiOjIyOSwidXJpIjoiL3JzL2N1c3RvbWVyL3VwZGF0ZSIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoxLCJyb2xlIjoiYWRtaW4iLCJ1cmlJZCI6MTk4LCJ1cmkiOiIvcnMvcm9sZS9nZXQiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6MSwicm9sZSI6ImFkbWluIiwidXJpSWQiOjYzLCJ1cmkiOiJodHRwOi8vbG9jYWxob3N0L3JzL3VwZGF0ZUNsaWVudCIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoxLCJyb2xlIjoiYWRtaW4iLCJ1cmlJZCI6MjcyLCJ1cmkiOiIvcnMvb3JkZXIvcGFja2FnZS9kZWxldGUiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6MSwicm9sZSI6ImFkbWluIiwidXJpSWQiOjI2MCwidXJpIjoiL3JzL29yZGVyL2FkZCIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoxLCJyb2xlIjoiYWRtaW4iLCJ1cmlJZCI6MjM4LCJ1cmkiOiIvcnMvb3B0aW9ucy91cGRhdGUiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6MSwicm9sZSI6ImFkbWluIiwidXJpSWQiOjMzMCwidXJpIjoiL3JzL2NvbnRlbnQvaGl0cyIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoxLCJyb2xlIjoiYWRtaW4iLCJ1cmlJZCI6MjMxLCJ1cmkiOiIvcnMvY3VzdG9tZXIvbGlzdCIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoxLCJyb2xlIjoiYWRtaW4iLCJ1cmlJZCI6MTk5LCJ1cmkiOiIvcnMvcm9sZS9kZWxldGUiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6MSwicm9sZSI6ImFkbWluIiwidXJpSWQiOjc3LCJ1cmkiOiJodHRwOi8vbG9jYWxob3N0L3JzL2FkZENsaWVudFNjb3BlIiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjEsInJvbGUiOiJhZG1pbiIsInVyaUlkIjoyNDIsInVyaSI6Ii9ycy9vcHRpb25zL2RlbGV0ZSIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoxLCJyb2xlIjoiYWRtaW4iLCJ1cmlJZCI6MjczLCJ1cmkiOiIvcnMvaW1hZ2UvYWRkIiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjEsInJvbGUiOiJhZG1pbiIsInVyaUlkIjoyNjEsInVyaSI6Ii9ycy9vcmRlci91cGRhdGUiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6Miwicm9sZSI6InVzZXIiLCJ1cmlJZCI6MjMxLCJ1cmkiOiIvcnMvY3VzdG9tZXIvbGlzdCIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoyLCJyb2xlIjoidXNlciIsInVyaUlkIjoyNzUsInVyaSI6Ii9ycy9pbWFnZS9kZXRhaWxzIiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjIsInJvbGUiOiJ1c2VyIiwidXJpSWQiOjI0NywidXJpIjoiL3JzL2RldGFpbHMvZ2V0QnlTa3UiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6Miwicm9sZSI6InVzZXIiLCJ1cmlJZCI6MjM1LCJ1cmkiOiIvcnMvcHJvZHVjdC9nZXQiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6Miwicm9sZSI6InVzZXIiLCJ1cmlJZCI6Mjc2LCJ1cmkiOiIvcnMvaW1hZ2UvcGFnZS9jb3VudCIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoyLCJyb2xlIjoidXNlciIsInVyaUlkIjoyNDgsInVyaSI6Ii9ycy9kZXRhaWxzL2dldEJ5QmFyQ29kZSIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoyLCJyb2xlIjoidXNlciIsInVyaUlkIjoyMzksInVyaSI6Ii9ycy9vcHRpb25zL2dldCIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoyLCJyb2xlIjoidXNlciIsInVyaUlkIjoyNzcsInVyaSI6Ii9ycy9pbWFnZS9saXN0IiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjIsInJvbGUiOiJ1c2VyIiwidXJpSWQiOjI1NywidXJpIjoiL3JzL2JhckNvZGUvZ2V0IiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjIsInJvbGUiOiJ1c2VyIiwidXJpSWQiOjI0MCwidXJpIjoiL3JzL29wdGlvbnMvZ2V0QnlEZXRhaWxzIiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjIsInJvbGUiOiJ1c2VyIiwidXJpSWQiOjY4LCJ1cmkiOiJodHRwOi8vbG9jYWxob3N0L3JzL2RlbGV0ZUNsaWVudEFsbG93ZWRVcmkiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6Miwicm9sZSI6InVzZXIiLCJ1cmlJZCI6MjgxLCJ1cmkiOiIvcnMvY29udGVudC9nZXQiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6Miwicm9sZSI6InVzZXIiLCJ1cmlJZCI6MjU4LCJ1cmkiOiIvcnMvYmFyQ29kZS9nZXRCeURldGFpbHMiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6Miwicm9sZSI6InVzZXIiLCJ1cmlJZCI6MjQxLCJ1cmkiOiIvcnMvb3B0aW9ucy9zZWFyY2hCeU9wdGlvbiIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoyLCJyb2xlIjoidXNlciIsInVyaUlkIjo4MCwidXJpIjoiaHR0cDovL2xvY2FsaG9zdC9ycy9hZGRDbGllbnRSb2xlVXJpIiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjIsInJvbGUiOiJ1c2VyIiwidXJpSWQiOjI2MiwidXJpIjoiL3JzL29yZGVyL2dldCIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoyLCJyb2xlIjoidXNlciIsInVyaUlkIjoyODIsInVyaSI6Ii9ycy9jb250ZW50L2xpc3QiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6Miwicm9sZSI6InVzZXIiLCJ1cmlJZCI6MjQ1LCJ1cmkiOiIvcnMvZGV0YWlscy9nZXQiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6Miwicm9sZSI6InVzZXIiLCJ1cmlJZCI6MjI1LCJ1cmkiOiIvcnMvYWRkcmVzcy9nZXQiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6Miwicm9sZSI6InVzZXIiLCJ1cmlJZCI6MjYzLCJ1cmkiOiIvcnMvb3JkZXIvbGlzdCIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoyLCJyb2xlIjoidXNlciIsInVyaUlkIjoyNDYsInVyaSI6Ii9ycy9kZXRhaWxzL2dldEJ5UHJvZHVjdCIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoyLCJyb2xlIjoidXNlciIsInVyaUlkIjoyMjYsInVyaSI6Ii9ycy9hZGRyZXNzL2xpc3QiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6Miwicm9sZSI6InVzZXIiLCJ1cmlJZCI6MjY0LCJ1cmkiOiIvcnMvb3JkZXIvY3VzdG9tZXIvbGlzdCIsImNsaWVudElkIjo0MDN9LHsiY2xpZW50Um9sZUlkIjoyLCJyb2xlIjoidXNlciIsInVyaUlkIjoyMzAsInVyaSI6Ii9ycy9jdXN0b21lci9nZXQiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6Miwicm9sZSI6InVzZXIiLCJ1cmlJZCI6MjcxLCJ1cmkiOiIvcnMvb3JkZXIvcGFja2FnZS9nZXQiLCJjbGllbnRJZCI6NDAzfV0sImV4cGlyZXNJbiI6MzYwMCwiaWF0IjoxNTAzNzc1ODAzLCJ0b2tlblR5cGUiOiJhY2Nlc3MiLCJleHAiOjE1MDM3Nzk0MDMsImlzcyI6IlVsYm9yYSBPYXV0aDIgU2VydmVyIn0.JToaPwx265riRGwaEal8x7KW374JNdUQumn735uSG9A"

func TestContentService_AddContent(t *testing.T) {
	var c ContentService
	c.ClientID = "403"
	c.Host = "http://localhost:3008"
	c.Token = token
	var ct Content
	ct.Category = "books1"
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
	ct.Category = "music2"
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
	ct.Category = "books1"
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

func TestContentService_GetContentListCategory(t *testing.T) {
	var c ContentService
	c.Host = "http://localhost:3008"
	_, res := c.GetContentListCategory("403", "books1")
	fmt.Print("res category list: ")
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
