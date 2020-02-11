package restify

import (
	"os"
	"testing"
	"time"

	restify "github.com/SpaceStock/go-restify"
	"github.com/SpaceStock/go-restify/enum/onfailure"
	"github.com/SpaceStock/go-restify/scenario"
	"github.com/stretchr/testify/assert"
)

func Test_Scenario(t *testing.T) {
	scn := scenario.New()
	results := scn.
		Set().ID("").Name("Complex Testing").
		AddCase(restify.TestCase{
			Order:       1,
			Name:        "Firebase Auth",
			Description: "",
			Request: restify.Request{
				URL:    "https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyPassword?key=AIzaSyD-HHHsWb82AFmdXtm0t86Nb9uoMJutrU0",
				Method: "POST",
				Payload: map[string]interface{}{
					"email":             "superadmin@spacestock.com",
					"password":          "admin@123",
					"returnSecureToken": true,
				},
			},
			Expect: restify.Expect{
				StatusCode: 200,
				Evaluate: []restify.Expression{
					"idToken != ''",
				},
			},
			Pipeline: restify.Pipeline{
				Cache:     true,
				CacheAs:   "auth",
				OnFailure: onfailure.Exit,
			},
		}).
		AddCase(restify.TestCase{
			Order:       2,
			Name:        "Get Complex Apartment",
			Description: "",
			Request: restify.Request{
				URL:    "https://stg-satpam.spacestock.com/1.0/complex/apartment?page=1&size=1",
				Method: "GET",
				Headers: map[string]string{
					"Authorization": "Bearer {auth.idToken}",
				},
				Payload: nil,
			},
			Expect: restify.Expect{
				StatusCode: 200,
				Evaluate:   []restify.Expression{},
			},
			Pipeline: restify.Pipeline{
				Cache:     true,
				CacheAs:   "list",
				OnFailure: onfailure.Exit,
			},
		}).
		AddCase(restify.TestCase{
			Order:       3,
			Name:        "Get One Apartment",
			Description: "",
			Request: restify.Request{
				URL:    "https://stg-satpam.spacestock.com/1.0/complex/apartment/{list.data.[0].id}",
				Method: "GET",
				Headers: map[string]string{
					"Authorization": "Bearer {auth.idToken}",
				},
				Payload: nil,
			},
			Expect: restify.Expect{
				StatusCode: 200,
				Evaluate: []restify.Expression{
					"id === '{list.data.[0.id]}'",
				},
			},
			Pipeline: restify.Pipeline{
				Cache:     true,
				CacheAs:   "oneApt",
				OnFailure: onfailure.Exit,
			},
		}).
		AddCase(restify.TestCase{
			Order:       4,
			Name:        "Create Apartment",
			Description: "",
			Request: restify.Request{
				URL:    "https://stg-satpam.spacestock.com/1.0/complex/apartment",
				Method: "POST",
				Headers: map[string]string{
					"Authorization": "Bearer {auth.idToken}",
				},
				Payload: nil,
			},
			Expect: restify.Expect{
				StatusCode: 201,
				Evaluate:   []restify.Expression{},
			},
			Pipeline: restify.Pipeline{
				Cache:     false,
				CacheAs:   "aptOne",
				OnFailure: onfailure.Exit,
			},
		}).End().
		Run(os.Stdout)

	if len(results) <= 0 {
		t.Error("No result returned")
	}

	if results[0].Success {
		t.Error("This case should have failed")
	}
}

//	Data True
func Test_Scenario2(t *testing.T) {
	scn := scenario.New()
	results := scn.
		Set().ID("").Name("Scenario 2").
		AddCase(restify.TestCase{
			Order:       1,
			Name:        "Test Case 1",
			Description: "",
			Request: restify.Request{
				URL:     "http://jsonplaceholder.typicode.com/posts/1",
				Method:  "GET",
				Payload: nil,
			},
			Expect: restify.Expect{
				StatusCode: 200,
				Evaluate: []restify.Expression{
					"userId && userId === 1",
					"id && id === 1",
				},
			},
			Pipeline: restify.Pipeline{
				Cache:     true,
				CacheAs:   "tc1",
				OnFailure: onfailure.Exit,
			},
		}).
		AddCase(restify.TestCase{
			Order:       2,
			Name:        "Test Case 2",
			Description: "",
			Request: restify.Request{
				URL:     "http://jsonplaceholder.typicode.com/posts/2",
				Method:  "GET",
				Payload: nil,
			},
			Expect: restify.Expect{
				StatusCode: 200,
				Evaluate: []restify.Expression{
					"userId && userId === 1",
					"id && id === 2",
				},
			},
			Pipeline: restify.Pipeline{
				Cache:     true,
				CacheAs:   "tc2",
				OnFailure: onfailure.Exit,
			},
		}).
		AddCase(restify.TestCase{
			Order:       3,
			Name:        "Test Case 3",
			Description: "",
			Request: restify.Request{
				URL:     "http://jsonplaceholder.typicode.com/posts/3",
				Method:  "GET",
				Payload: nil,
			},
			Expect: restify.Expect{
				StatusCode: 200,
				Evaluate: []restify.Expression{
					"userId && userId === 1",
					"id && id === 3",
				},
			},
			Pipeline: restify.Pipeline{
				Cache:     true,
				CacheAs:   "tc3",
				OnFailure: onfailure.Exit,
			},
		}).End().
		Run(os.Stdout)

	assert.NotEqual(t, 0, len(results), "Seharusnya bukan 0")
	assert.True(t, results[0].Success)
	assert.True(t, results[1].Success)
}

//	Pipeline.OnFailure=Exit
func Test_Scenario3(t *testing.T) {
	scn := scenario.New()
	results := scn.
		Set().ID("").Name("Scenario 2").
		AddCase(restify.TestCase{
			Order:       1,
			Name:        "Test Case 1",
			Description: "",
			Request: restify.Request{
				URL:     "http://jsonplaceholder.typicode.com/posts/1",
				Method:  "GET",
				Payload: nil,
			},
			Expect: restify.Expect{
				StatusCode: 200,
				Evaluate: []restify.Expression{
					"userId && userId === 1",
					"id && id === 1",
				},
			},
			Pipeline: restify.Pipeline{
				Cache:     true,
				CacheAs:   "tc1",
				OnFailure: onfailure.Exit,
			},
		}).
		AddCase(restify.TestCase{
			Order:       2,
			Name:        "Test Case 2",
			Description: "",
			Request: restify.Request{
				URL:     "http://jsonplaceholder.typicode.com/posts/2",
				Method:  "GET",
				Payload: nil,
			},
			Expect: restify.Expect{
				StatusCode: 200,
				Evaluate: []restify.Expression{
					"userId && userId === 1",
					"id && id === 3", //	False
				},
			},
			Pipeline: restify.Pipeline{
				Cache:     true,
				CacheAs:   "tc2",
				OnFailure: onfailure.Exit,
			},
		}).
		AddCase(restify.TestCase{
			Order:       3,
			Name:        "Test Case 3",
			Description: "",
			Request: restify.Request{
				URL:     "http://jsonplaceholder.typicode.com/posts/3",
				Method:  "GET",
				Payload: nil,
			},
			Expect: restify.Expect{
				StatusCode: 200,
				Evaluate: []restify.Expression{
					"userId && userId === 1",
					"id && id === 3",
				},
			},
			Pipeline: restify.Pipeline{
				Cache:     true,
				CacheAs:   "tc3",
				OnFailure: onfailure.Exit,
			},
		}).End().
		Run(os.Stdout)

	assert.Equal(t, 2, len(results), "Test Case = 2")
	assert.True(t, results[0].Success)
	assert.False(t, results[1].Success)
}

// Pipeline.OnFailure=Fallthrough
func Test_Scenario4(t *testing.T) {
	scn := scenario.New()
	results := scn.
		Set().ID("").Name("Scenario 3").
		AddCase(restify.TestCase{
			Order:       1,
			Name:        "Test Case 1",
			Description: "",
			Request: restify.Request{
				URL:     "http://jsonplaceholder.typicode.com/posts/1",
				Method:  "GET",
				Payload: nil,
			},
			Expect: restify.Expect{
				StatusCode: 200,
				Evaluate: []restify.Expression{
					"userId && userId === 1",
					"id && id === 1",
				},
			},
			Pipeline: restify.Pipeline{
				Cache:     true,
				CacheAs:   "tc1",
				OnFailure: onfailure.Exit,
			},
		}).
		AddCase(restify.TestCase{
			Order:       2,
			Name:        "Test Case 2",
			Description: "",
			Request: restify.Request{
				URL:     "http://jsonplaceholder.typicode.com/posts/2",
				Method:  "GET",
				Payload: nil,
			},
			Expect: restify.Expect{
				StatusCode: 200,
				Evaluate: []restify.Expression{
					"userId && userId === 1",
					"id && id === 1", //	False
				},
			},
			Pipeline: restify.Pipeline{
				Cache:     true,
				CacheAs:   "tc2",
				OnFailure: onfailure.Fallthrough,
			},
		}).
		AddCase(restify.TestCase{
			Order:       3,
			Name:        "Test Case 3",
			Description: "",
			Request: restify.Request{
				URL:     "http://jsonplaceholder.typicode.com/posts/3",
				Method:  "GET",
				Payload: nil,
			},
			Expect: restify.Expect{
				StatusCode: 200,
				Evaluate: []restify.Expression{
					"userId && userId === 1",
					"id && id === 3",
				},
			},
			Pipeline: restify.Pipeline{
				Cache:     true,
				CacheAs:   "tc3",
				OnFailure: onfailure.Exit,
			},
		}).
		End().
		Run(os.Stdout)

	assert.Equal(t, 3, len(results), "Test Case = 3")

	assert.True(t, results[0].Success)
	assert.False(t, results[1].Success)
	assert.True(t, results[2].Success)
}

//	Time Delay
func Test_Scenario5(t *testing.T) {
	before := time.Now()
	results := scenario.New().
		Set().ID("").Name("Scenario 5").
		AddCase(restify.TestCase{
			Order:       1,
			Name:        "Test Case 1",
			Description: "",
			Request: restify.Request{
				URL:     "http://jsonplaceholder.typicode.com/posts/1",
				Method:  "GET",
				Payload: nil,
			},
			Expect: restify.Expect{
				StatusCode: 200,
				Evaluate: []restify.Expression{
					"userId && userId === 1",
					"id && id === 1",
				},
			},
			Pipeline: restify.Pipeline{
				Cache:     true,
				CacheAs:   "tc1",
				OnFailure: onfailure.Exit,
				Delay:     restify.Delay("3s"),
			},
		}).End().
		Run(os.Stdout)
	after := time.Since(before).Seconds() // .Second -> float64

	assert.GreaterOrEqual(t, after, 3.0)
	assert.NotEqual(t, 0, len(results), "Seharusnya bukan 0")
	assert.True(t, results[0].Success)
	assert.False(t, results.IsFailed())
}

//	assess cache - memastikan cache 1 bernilai benar di chache 2
func Test_Scenario6(t *testing.T) {
	results := scenario.New().
		Set().ID("").Name("Scenario 6").
		AddCase(restify.TestCase{
			Order:       1,
			Name:        "Test Case 1",
			Description: "",
			Request: restify.Request{
				URL:     "http://jsonplaceholder.typicode.com/posts/1",
				Method:  "GET",
				Payload: nil,
			},
			Expect: restify.Expect{
				StatusCode: 200,
				Evaluate: []restify.Expression{
					"userId && userId === 1",
				},
			},
			Pipeline: restify.Pipeline{
				Cache:     true,
				CacheAs:   "result_one",
				OnFailure: onfailure.Exit,
			},
		}).
		AddCase(restify.TestCase{
			Order:       2,
			Name:        "Test Case 2",
			Description: "",
			Request: restify.Request{
				URL:     "http://jsonplaceholder.typicode.com/posts/1",
				Method:  "GET",
				Payload: nil,
			},
			Expect: restify.Expect{
				StatusCode: 200,
				Evaluate: []restify.Expression{
					"id && id === result_one.id",
				},
			},
			Pipeline: restify.Pipeline{
				Cache:     true,
				CacheAs:   "result_two",
				OnFailure: onfailure.Exit,
			},
		}).
		AddCase(restify.TestCase{
			Order:       3,
			Name:        "Test Case 3",
			Description: "",
			Request: restify.Request{
				URL:     "http://jsonplaceholder.typicode.com/posts/2",
				Method:  "GET",
				Payload: nil,
			},
			Expect: restify.Expect{
				StatusCode: 200,
				Evaluate: []restify.Expression{
					"id && id === result_one.id",
				},
			},
			Pipeline: restify.Pipeline{
				Cache:     true,
				CacheAs:   "result_three",
				OnFailure: onfailure.Exit,
			},
		}).End().
		Run(os.Stdout)

	assert.True(t, results[0].Success)
	assert.True(t, results[1].Success)
	assert.False(t, results[2].Success)
}
