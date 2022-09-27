package api

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_UserCurrent(t *testing.T) {
	type args struct {
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
				method:     "GET",
				path:       "/user",
				statusCode: 200,
				golden:     "./testdata/user.golden.json",
			},
			wantErr: false,
		},
		{
			name: "401",
			args: args{
				method:     "GET",
				path:       "/user",
				statusCode: 401,
				golden:     "./testdata/user.401.json",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := GoldenEndpointServer(t, tt.args.method, tt.args.path, tt.args.statusCode, tt.args.golden)
			defer ts.Close()

			c := client(ts.URL)
			got, err := c.UserCurrent()
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
