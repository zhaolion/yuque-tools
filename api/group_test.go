package api

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_GroupCreate(t *testing.T) {
	type args struct {
		req        *GroupCreateRequest
		method     string
		path       string
		statusCode int
		golden     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "200",
			args: args{
				req: &GroupCreateRequest{
					Name:        "test-lion",
					Login:       "test-lion",
					Description: "test-lion",
				},
				method:     "POST",
				path:       "/groups",
				statusCode: 200,
				golden:     "./testdata/groups.create.golden.json",
			},
			wantErr: false,
		},
		{
			name: "401",
			args: args{
				req: &GroupCreateRequest{
					Name:        "",
					Login:       "test-lion",
					Description: "test-lion",
				},
				method:     "POST",
				path:       "/groups",
				statusCode: 400,
				golden:     "./testdata/groups.create.400.json",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := GoldenEndpointServer(t, tt.args.method, tt.args.path, tt.args.statusCode, tt.args.golden)
			defer ts.Close()

			c := client(ts.URL)
			got, err := c.GroupCreate(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hello() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				require.Nil(t, got)
				if e, ok := err.(*YuqueError); ok {
					require.NotEmpty(t, e.Status)
					require.NotEmpty(t, e.Message)
				}
			} else {
				require.NotNil(t, got)
				require.NotEmpty(t, got.Data)
			}
		})
	}
}
