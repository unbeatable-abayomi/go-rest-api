// *build e2e
package test



import(
	_"fmt"
	"testing"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
) 


func TestGetComments(t *testing.T){
	client := resty.New()
	resp, err := client.R().Get(BASE_URL+"api/comment")
	if err != nil {
		t.Fail()
	}
	assert.Equal(t,200,resp.StatusCode())
    //fmt.Println(resp.StatusCode())
}

func TestPostComments(t *testing.T){
	client := resty.New()
	//resp, err := client.R().Post(BASE_URL+"api/comment")
	resp, err := client.R().SetBody(
		`{"slug":"/", "author":"234849", "body":"hello world"}`).Post(BASE_URL+"api/comment")

	assert.NoError(t,err)
	if err != nil {
		t.Fail()
	}
	assert.Equal(t,200,resp.StatusCode())
    //fmt.Println(resp.StatusCode())
}